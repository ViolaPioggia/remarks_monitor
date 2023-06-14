package logic

import (
	"context"
	"fmt"
	"remarks_monitor/app/map1/cmd/rpc/mapwork"
	mapwork2 "remarks_monitor/app/map2/cmd/rpc/mapwork"
	mapwork3 "remarks_monitor/app/map3/cmd/rpc/mapwork"
	"remarks_monitor/app/master/cmd/rpc/internal/svc"
	"remarks_monitor/app/master/cmd/rpc/pb"
	"remarks_monitor/common/tool"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMapLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMapLogic {
	return &GetMapLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMapLogic) GetMap(in *pb.GetMapReq) (*pb.GetMapResp, error) {
	kind := in.Type
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	//MNum := tool.GetFileNum(tool.GetWD() + "/data/remarks_monitor/input/")
	Paths := tool.GetFilePaths(tool.GetWD() + "/data/remarks_monitor/input/")

	flag := make([]int, len(Paths))
	//fallBehindChan := make(chan int64, len(Paths))
	for i := 0; i < len(Paths); i++ {
		fmt.Println("start")
		select {
		case ch1 <- i:
			fmt.Println(1)
			go l.sendReqToMap1(int64(i), Paths[i], flag, ch1, kind)
			<-ch1
			fmt.Println(2)
		case ch2 <- i:
			fmt.Println(3)
			go l.sendReqToMap2(int64(i), Paths[i], flag, ch2, kind)
			fmt.Println(4)
		case ch3 <- i:
			fmt.Println(5)
			go l.sendReqToMap3(int64(i), Paths[i], flag, ch3, kind)
			fmt.Println(6)
		}
	}
	time.Sleep(1 * time.Second)
	for k, v := range flag {
		if v == -1 {
			select {
			case ch1 <- k:
				go l.sendReqToMap1(int64(k), Paths[k], flag, ch1, kind)
			case ch2 <- k:
				go l.sendReqToMap2(int64(k), Paths[k], flag, ch2, kind)
			case ch3 <- k:
				go l.sendReqToMap3(int64(k), Paths[k], flag, ch3, kind)
			}
		}
	}
	return &pb.GetMapResp{}, nil
}

func (l *GetMapLogic) sendReqToMap1(MNum int64, Paths string, flag []int, chanlock chan int, kind int64) {
	flag[MNum] = -1
	_, err := l.svcCtx.MapRpc1.MapWork(l.ctx, &mapwork.GetMapWorkReq{MNum: strconv.FormatInt(MNum, 10), Paths: Paths, Type: kind})
	if err != nil {
		logx.Errorf("send map request to MapNode 1 failed")
		return
	}
	fmt.Println("send request to map 1 success")
	// 将结果发送到通道
	flag[MNum] = 0
	<-chanlock
}
func (l *GetMapLogic) sendReqToMap2(MNum int64, Paths string, flag []int, chanlock chan int, kind int64) {
	flag[MNum] = -1
	_, err := l.svcCtx.MapRpc2.MapWork(l.ctx, &mapwork2.GetMapWorkReq{MNum: strconv.FormatInt(MNum, 10), Paths: Paths, Type: kind})
	if err != nil {
		logx.Errorf("send map request to MapNode 2 failed")
		return
	}

	// 将结果发送到通道
	flag[MNum] = 0
	<-chanlock
}
func (l *GetMapLogic) sendReqToMap3(MNum int64, Paths string, flag []int, chanlock chan int, kind int64) {
	flag[MNum] = -1
	_, err := l.svcCtx.MapRpc3.MapWork(l.ctx, &mapwork3.GetMapWorkReq{MNum: strconv.FormatInt(MNum, 10), Paths: Paths, Type: kind})
	if err != nil {
		logx.Errorf("send map request to MapNode 3 failed")
		return
	}

	// 将结果发送到通道
	flag[MNum] = 0
	<-chanlock
}
