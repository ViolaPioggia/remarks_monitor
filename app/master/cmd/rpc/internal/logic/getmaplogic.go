package logic

import (
	"context"
	"remarks_monitor/app/map/cmd/rpc/mapwork"
	"remarks_monitor/common/tool"
	"strconv"

	"remarks_monitor/app/master/cmd/rpc/internal/svc"
	"remarks_monitor/app/master/cmd/rpc/pb"

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
	MNum := tool.GetFileNum(tool.GetWD() + "/remarks_monitor/data/remarks_monitor/input")
	Paths := tool.GetFilePaths(tool.GetWD() + "/remarks_monitor/data/remarks_monitor/input")
	_, err := l.svcCtx.MapRpc.MapWork(l.ctx, &mapwork.GetMapWorkReq{MNum: strconv.FormatInt(MNum, 10), Paths: Paths})
	if err != nil {
		return nil, err
	}
	return &pb.GetMapResp{}, nil
}
