package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"remarks_monitor/common/tool"
	"sort"

	"remarks_monitor/app/reduce1/cmd/rpc/internal/svc"
	"remarks_monitor/app/reduce1/cmd/rpc/reduce1"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReduceWork1Logic struct {
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
func NewReduceWork1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *ReduceWork1Logic {
	return &ReduceWork1Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReduceWork1Logic) ReduceWork1(in *reduce1.ReduceWorkReq) (*reduce1.ReduceWorkResp, error) {
	RNum := in.RNum //map的数量
	files := in.Paths
	DoReduceTask(RNum, files)
	return &reduce1.ReduceWorkResp{}, nil
}
func DoReduceTask(reduceFileNum int64, fileslice []string) {
	intermediate := shuffle(fileslice)
	dir := tool.GetWD() + "/data/remarks_monitor/reduce/"
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
}

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
