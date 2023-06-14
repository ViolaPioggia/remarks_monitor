package logic

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"remarks_monitor/app/map2/cmd/rpc/map2"

	"github.com/zeromicro/go-zero/core/logx"
	"remarks_monitor/app/map2/cmd/rpc/internal/svc"
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

func (l *MapWorkLogic) MapWork(in *map2.GetMapWorkReq) (*map2.GetMapWorkResp, error) {
	MNum := in.MNum
	num, _ := strconv.ParseInt(MNum, 10, 64)
	File := in.Paths
	path := DoMapTask(File, num, in.Type)

	return &map2.GetMapWorkResp{Director: path}, nil
}

func DoMapTask(filepath string, id int64, kind int64) string {
	var intermediate []KeyValue
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return ""
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
			Key:   fields[kind],
			Value: strconv.Itoa(1),
		}
		intermediate = append(intermediate, v1)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read file:", err)
		return ""
	}
	tool.GetWD()
	oname := "mr-tmp-" + strconv.FormatInt(id, 10)
	if kind == 0 {
		ofile, err := os.Create(tool.GetWD() + "/data/remarks_monitor/map_username/" + oname)
		if err != nil {
			logx.Errorf(err.Error())
		}
		enc := json.NewEncoder(ofile)
		for _, kv := range intermediate {
			enc.Encode(kv)
		}
		ofile.Close()
	} else if kind == 1 {
		ofile, err := os.Create(tool.GetWD() + "/data/remarks_monitor/map_domain/" + oname)
		if err != nil {
			logx.Errorf(err.Error())
		}
		enc := json.NewEncoder(ofile)
		for _, kv := range intermediate {
			enc.Encode(kv)
		}
		ofile.Close()
	} else if kind == 2 {
		ofile, err := os.Create(tool.GetWD() + "/data/remarks_monitor/map_content/" + oname)
		if err != nil {
			logx.Errorf(err.Error())
		}
		enc := json.NewEncoder(ofile)
		for _, kv := range intermediate {
			enc.Encode(kv)
		}
		ofile.Close()
	}
	return tool.GetWD() + "/data/remarks_monitor/MapReduce/" + oname
}
