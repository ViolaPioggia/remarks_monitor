package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"remarks_monitor/app/master/cmd/rpc/internal/svc"
	"remarks_monitor/app/master/cmd/rpc/pb"
	"remarks_monitor/app/reduce1/cmd/rpc/reducework1"
	"remarks_monitor/app/reduce2/cmd/rpc/reducework2"
	"remarks_monitor/common/tool"
	"time"
)

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
	Paths := tool.GetFilePaths(tool.GetWD() + "/data/remarks_monitor/MapReduce/")
	flag := make([]int, len(Paths))
	var j int64
	var i int64 = 0
	//fallBehindChan := make(chan int64, len(Paths))
	for x := 0; i+2 <= int64(len(Paths)); x++ {
		if i+2 >= int64(len(Paths)) {
			j = int64(len(Paths))
		} else {
			j = i + 3
		}
		fmt.Println("start")
		fmt.Println(i, j)
		select {
		case ch1 <- int(i):
			fmt.Println(1)
			fmt.Println(Paths[i:j])
			go l.sendReqToReduce1(int64(x), Paths[i:j], flag, ch1)
			<-ch1
			fmt.Println(2)
		case ch2 <- int(i):
			fmt.Println(3)
			fmt.Println(Paths[i:j])
			go l.sendReqToReduce2(int64(x), Paths[i:j], flag, ch2)
			fmt.Println(4)
		}
		i += 3
	}
	time.Sleep(1 * time.Second)
	for k, v := range flag {
		if v == -1 {
			fmt.Println("haihaihai")
			if k+3 >= len(Paths) {
				j = int64(len(Paths))
			} else {
				j = i + 3
			}
			select {
			case ch1 <- k:
				fmt.Println(1)
				go l.sendReqToReduce1(int64(k/3), Paths[k:j], flag, ch1)
				<-ch1
				fmt.Println(2)
			case ch2 <- k:
				fmt.Println(3)
				go l.sendReqToReduce2(int64(k/3), Paths[k:j], flag, ch2)
				fmt.Println(4)
			}
		}
	}
	return &pb.GetRpcResp{}, nil
}

func (l *GetRpcLogic) sendReqToReduce1(RNum int64, Paths []string, flag []int, chanlock chan int) {
	flag[RNum] = -1
	_, err := l.svcCtx.ReduceRpc1.ReduceWork1(l.ctx, &reducework1.ReduceWorkReq{RNum: RNum, Paths: Paths})
	if err != nil {
		logx.Errorf("send reduce request to ReduceNode 1 failed")
		return
	}
	fmt.Println("send request to reduce 1 success")
	// 将结果发送到通道
	flag[RNum] = 0
	<-chanlock
}
func (l *GetRpcLogic) sendReqToReduce2(RNum int64, Paths []string, flag []int, chanlock chan int) {
	flag[RNum] = -1
	_, err := l.svcCtx.ReduceRpc2.ReduceWork2(l.ctx, &reducework2.ReduceWorkReq{RNum: RNum, Paths: Paths})
	if err != nil {
		logx.Errorf("send reduce request to ReduceNode 2 failed")
		return
	}
	fmt.Println("send request to reduce 2 success")
	// 将结果发送到通道
	flag[RNum] = 0
	<-chanlock
}
