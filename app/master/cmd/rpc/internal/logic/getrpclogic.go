package logic

import (
	"bufio"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"remarks_monitor/app/master/cmd/rpc/internal/svc"
	"remarks_monitor/app/master/cmd/rpc/pb"
	"remarks_monitor/app/master/model"
	"remarks_monitor/app/reduce1/cmd/rpc/reducework1"
	"remarks_monitor/app/reduce2/cmd/rpc/reducework2"
	"remarks_monitor/common/tool"
	"strconv"
	"strings"
	"time"
)

var t string

type GetRpcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRpcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRpcLogic {
	return &GetRpcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *GetRpcLogic) GetRpc(in *pb.GetRpcReq) (*pb.GetRpcResp, error) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	Paths := GetMapPaths(in.Type)
	fmt.Println(Paths)
	flag := make([]int, len(Paths))
	var j int64
	var i int64 = 0

	fmt.Println("11111")
	for x := 0; i < int64(len(Paths)); x++ {
		fmt.Println(2222222)
		if i+2 >= int64(len(Paths)) {
			j = int64(len(Paths))
		} else {
			j = i + 3
		}
		fmt.Println("reduce start")
		fmt.Println(i, j)
		select {
		case ch1 <- int(i):
			fmt.Println(1)
			fmt.Println(Paths[i:j])
			go l.sendReqToReduce1(int64(x), Paths[i:j], flag, ch1, in.Type)
			<-ch1
			fmt.Println(2)
		case ch2 <- int(i):
			fmt.Println(3)
			fmt.Println(Paths[i:j])
			go l.sendReqToReduce2(int64(x), Paths[i:j], flag, ch2, in.Type)
			fmt.Println(4)
		}
		i += 3
	}
	time.Sleep(1 * time.Second)
	for k, v := range flag {
		if v == -1 {
			if k+3 >= len(Paths) {
				j = int64(len(Paths))
			} else {
				j = i + 3
			}
			select {
			case ch1 <- k:
				fmt.Println(1)
				go l.sendReqToReduce1(int64(k/3), Paths[k:j], flag, ch1, in.Type)
				<-ch1
				fmt.Println(2)
			case ch2 <- k:
				fmt.Println(3)
				go l.sendReqToReduce2(int64(k/3), Paths[k:j], flag, ch2, in.Type)
				fmt.Println(4)
			}
		}
	}
	time.Sleep(time.Second)
	files := GetReducePaths(in.Type) // 文件列表

	sums := make(map[string]int) // 用于存储相同内容的值

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("Failed to open file: %s\n", file)
			continue
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, " ")
			if len(parts) != 2 {
				fmt.Printf("Invalid line format: %s\n", line)
				continue
			}

			str := parts[0]
			valueStr := parts[1]
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				fmt.Printf("Invalid value format: %s\n", valueStr)
				continue
			}

			sums[str] += value
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error scanning file: %s\n", file)
		}
	}

	// 输出累加结果
	var err error
	var data []model.Record
	if in.Type == 0 {
		t = "username"
	} else if in.Type == 1 {
		t = "domain"
	} else if in.Type == 2 {
		t = "content"
	}
	for k, v := range sums {
		r := &model.Record{
			Type:    t,
			Content: k,
			Nums:    int64(v),
		}
		data = append(data, *r)
	}
	l.svcCtx.RecordsModel.InsertMany(context.Background(), data)
	if err != nil {
		return nil, err
	}
	return &pb.GetRpcResp{}, nil
}

func (l *GetRpcLogic) sendReqToReduce1(RNum int64, Paths []string, flag []int, chanlock chan int, kind int64) {
	flag[RNum] = -1
	_, err := l.svcCtx.ReduceRpc1.ReduceWork1(l.ctx, &reducework1.ReduceWorkReq{RNum: RNum, Paths: Paths, Type: kind})
	if err != nil {
		logx.Errorf("send reduce request to ReduceNode 1 failed")
		return
	}
	fmt.Println("send request to reduce 1 success")
	// 将结果发送到通道
	flag[RNum] = 0
	<-chanlock
}
func (l *GetRpcLogic) sendReqToReduce2(RNum int64, Paths []string, flag []int, chanlock chan int, kind int64) {
	flag[RNum] = -1
	_, err := l.svcCtx.ReduceRpc2.ReduceWork2(l.ctx, &reducework2.ReduceWorkReq{RNum: RNum, Paths: Paths, Type: kind})
	if err != nil {
		logx.Errorf("send reduce request to ReduceNode 2 failed")
		return
	}
	fmt.Println("send request to reduce 2 success")
	// 将结果发送到通道
	flag[RNum] = 0
	<-chanlock
}
func GetMapPaths(kind int64) []string {
	switch kind {
	case 0:
		return tool.GetFilePaths(tool.GetWD() + "/data/remarks_monitor/map_username/")
	case 1:
		return tool.GetFilePaths(tool.GetWD() + "/data/remarks_monitor/map_domain/")
	case 2:
		return tool.GetFilePaths(tool.GetWD() + "/data/remarks_monitor/map_content/")
	}
	return nil
}
func GetReducePaths(kind int64) []string {
	switch kind {
	case 0:
		return tool.GetFilePaths(tool.GetWD() + "/data/remarks_monitor/reduce/username/")
	case 1:
		return tool.GetFilePaths(tool.GetWD() + "/data/remarks_monitor/reduce/domain/")
	case 2:
		return tool.GetFilePaths(tool.GetWD() + "/data/remarks_monitor/reduce/content/")
	}
	return nil
}
func (l *GetRpcLogic) WriteIntoDB(kind int64, content string, nums int64) error {

	r := &model.Record{
		Type:     t,
		Content:  content,
		Nums:     nums,
		UpdateAt: time.Time{},
		CreateAt: time.Time{},
	}
	err := l.svcCtx.RecordsModel.Insert(context.Background(), r)
	if err != nil {
		logx.Errorf("write record into mongo failed")
		fmt.Println(err)
		return err
	}

	return nil
}
