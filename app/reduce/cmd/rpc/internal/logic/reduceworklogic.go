package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"remarks_monitor/common/tool"
	"sort"
	"strconv"
	"sync"

	"remarks_monitor/app/reduce/cmd/rpc/internal/svc"
	"remarks_monitor/app/reduce/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var wg sync.WaitGroup

// Task worker向coordinator获取task的结构体
type Task struct {
	TaskType   TaskType // 任务类型判断到底是map还是reduce
	TaskId     int      // 任务的id
	ReducerNum int      // 传入的reducer的数量，用于hash
	FileSlice  []string // 输入文件的切片，map一个文件对应一个文件，reduce是对应多个temp中间值文件
}

// TaskArgs rpc应该传入的参数，可实际上应该什么都不用传,因为只是worker获取一个任务
type TaskArgs struct{}

// TaskType 对于下方枚举任务的父类型
type TaskType int

// Phase 对于分配任务阶段的父类型
type Phase int

// State 任务的状态的父类型
type State int

// 枚举任务的类型
const (
	MapTask TaskType = iota
	ReduceTask
	WaittingTask // Waittingen任务代表此时为任务都分发完了，但是任务还没完成，阶段未改变
	ExitTask     // exit
)

// 枚举阶段的类型
const (
	MapPhase    Phase = iota // 此阶段在分发MapTask
	ReducePhase              // 此阶段在分发ReduceTask
	AllDone                  // 此阶段已完成
)

// 任务状态类型
const (
	Working State = iota // 此阶段在工作
	Waiting              // 此阶段在等待执行
	Done                 // 此阶段已经做完
)

// Add your RPC definitions here.

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}

type ReduceWorkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}
type KeyValue struct {
	Key   string
	Value string
}
type SortedKey []KeyValue

// Len 重写len,swap,less才能排序
func (k SortedKey) Len() int           { return len(k) }
func (k SortedKey) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k SortedKey) Less(i, j int) bool { return k[i].Key < k[j].Key }

func NewReduceWorkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReduceWorkLogic {
	return &ReduceWorkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReduceWorkLogic) ReduceWork(in *pb.ReduceWorkReq) (*pb.ReduceWorkResp, error) {
	director := tool.GetWD() + "/data/remarks_monitor/MapReduce/"
	MNum := tool.GetFileNum(director) //map的数量
	var RNum int64                    //reduce的数量
	if MNum%10 != 0 {
		RNum = MNum/10 + 1
	} else {
		RNum = MNum / 10
	}
	files := tool.GetFilePaths(director)
	wg.Add(int(RNum))
	//_, reducef := loadPlugin(os.Args[1])
	var i int64
	DoReduceTask(i, files[:])

	wg.Wait()
	return &pb.ReduceWorkResp{}, nil
}
func DoReduceTask(reduceFileNum int64, fileslice []string) {
	intermediate := shuffle(fileslice)
	dir := tool.GetWD() + "/data/remarks_monitor/MapReduce/"
	//tempFile, err := ioutil.TempFile(dir, "mr-tmp-*")
	tempFile, err := os.CreateTemp(dir, "mr-tmp-*")
	fmt.Println(tempFile.Name())
	if err != nil {
		logx.Errorf("Failed to create temp file", err)
	}
	i := 0
	for i < len(intermediate) {
		j := i + 1
		for j < len(intermediate) && intermediate[j].Key == intermediate[i].Key {
			j++
		}
		var values int64
		for k := i; k < j; k++ {
			values++
		}
		//output := reducef(intermediate[i].Key, values)
		fmt.Fprintf(tempFile, "%v %v\n", intermediate[i].Key, values)
		i = j
	}
	tempFile.Close()
	// 在完全写入后进行重命名
	fn := fmt.Sprintf(dir+"mr-out-%d", reduceFileNum)
	err = os.Rename(tempFile.Name(), fn)
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
}

//func loadPlugin(filename string) (func(string, string) []KeyValue, func(string, []string) string) {
//	p, err := plugin.Open(filename)
//	if err != nil {
//		log.Fatalf("cannot load plugin %v", filename)
//	}
//	xmapf, err := p.Lookup("Map")
//	if err != nil {
//		log.Fatalf("cannot find Map in %v", filename)
//	}
//	mapf := xmapf.(func(string, string) []KeyValue)
//	xreducef, err := p.Lookup("Reduce")
//	if err != nil {
//		log.Fatalf("cannot find Reduce in %v", filename)
//	}
//	reducef := xreducef.(func(string, []string) string)
//
//	return mapf, reducef
//}

// 洗牌方法，得到一组排序好的kv数组
func shuffle(files []string) []KeyValue {
	var kva []KeyValue
	for _, filepath := range files {
		file, _ := os.Open(filepath)
		dec := json.NewDecoder(file)
		for {
			var kv KeyValue
			if err := dec.Decode(&kv); err != nil {
				break
			}
			kva = append(kva, kv)
		}
		file.Close()
	}
	sort.Sort(SortedKey(kva))
	return kva
}
