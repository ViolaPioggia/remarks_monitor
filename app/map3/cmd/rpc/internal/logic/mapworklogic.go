package logic

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"remarks_monitor/app/map3/cmd/rpc/internal/svc"
	"remarks_monitor/app/map3/cmd/rpc/map3"
	"remarks_monitor/common/tool"
	"strconv"
	"strings"
)

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

func (l *MapWorkLogic) MapWork(in *map3.GetMapWorkReq) (*map3.GetMapWorkResp, error) {
	MNum := in.MNum
	num, _ := strconv.ParseInt(MNum, 10, 64)
	File := in.Paths
	DoMapTask(File, int64(num))

	return &map3.GetMapWorkResp{Director: tool.GetWD() + "/remarks_monitor/data/remarks_monitor/MapReduce"}, nil
}

func DoMapTask(filepath string, id int64) {
	var intermediate []KeyValue
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
}

//func ihash(key string) int {
//	h := fnv.New32a()
//	h.Write([]byte(key))
//	return int(h.Sum32() & 0x7fffffff)
//}
