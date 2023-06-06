package logic

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"remarks_monitor/app/map/cmd/rpc/internal/svc"
	"remarks_monitor/app/map/cmd/rpc/pb"
	"remarks_monitor/common/tool"
	"strconv"
	"strings"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

var wg sync.WaitGroup

type KeyValue struct {
	Key   string
	Value string
}
type MapWorkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMapWorkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MapWorkLogic {
	return &MapWorkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MapWorkLogic) MapWork(in *pb.GetMapWorkReq) (*pb.GetMapWorkResp, error) {
	MNum := in.MNum
	num, _ := strconv.Atoi(MNum)
	Files := in.Paths
	//mapf, _ := loadPlugin(os.Args[1])
	wg.Add(num)
	for k, v := range Files {
		go DoMapTask(v, int64(k))
	}
	wg.Wait()
	return &pb.GetMapWorkResp{Director: tool.GetWD() + "/remarks_monitor/data/remarks_monitor/MapReduce"}, nil
}

//	func loadPlugin(filename string) (func(string, string) []KeyValue, func(string, []string) string) {
//		p, err := plugin.Open(filename)
//		if err != nil {
//			logx.Errorf("cannot load plugin %v", filename)
//		}
//		xmapf, err := p.Lookup("Map")
//		if err != nil {
//			logx.Errorf("cannot find Map in %v", filename)
//		}
//		mapf := xmapf.(func(string, string) []KeyValue)
//		xreducef, err := p.Lookup("Reduce")
//		if err != nil {
//			logx.Errorf("cannot find Reduce in %v", filename)
//		}
//		reducef := xreducef.(func(string, []string) string)
//
//		return mapf, reducef
//	}
func DoMapTask(filepath string, id int64) {
	var intermediate []KeyValue
	//file, err := os.Open(filepath)
	//if err != nil {
	//	logx.Errorf("cannot open %v", filepath)
	//}
	//// 通过io工具包获取content,作为mapf的参数
	//content, err := io.ReadAll(file)
	//if err != nil {
	//	logx.Errorf("cannot read %v", filepath)
	//}
	//file.Close()
	// 打开文件
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()
	// 创建文件读取器
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容
	for scanner.Scan() {
		line := scanner.Text()

		// 解析行内容为结构体字段
		fields := strings.Split(line, " ")

		// 创建结构体对象并添加到数组中
		v1 := KeyValue{
			Key:   fields[1],
			Value: strconv.Itoa(1),
		}
		intermediate = append(intermediate, v1)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
	tool.GetWD()
	oname := "mr-tmp-" + strconv.FormatInt(id, 10)
	ofile, err := os.Create(tool.GetWD() + "/data/remarks_monitor/MapReduce/" + oname)
	if err != nil {
		logx.Errorf("create mr-tmp file failed,err is", err)
	}
	enc := json.NewEncoder(ofile)
	for _, kv := range intermediate {
		enc.Encode(kv)
	}
	ofile.Close()
	wg.Done()

}

//func ihash(key string) int {
//	h := fnv.New32a()
//	h.Write([]byte(key))
//	return int(h.Sum32() & 0x7fffffff)
//}
