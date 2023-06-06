package logic

import (
	"context"
	"remarks_monitor/app/master/cmd/rpc/internal/svc"
	"remarks_monitor/app/master/cmd/rpc/pb"
	"remarks_monitor/app/reduce/cmd/rpc/reducework"

	"github.com/zeromicro/go-zero/core/logx"
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
	//director := tool.GetWD() + "/remarks_monitor/data/remarks_monitor/MapReduce"
	_, err := l.svcCtx.ReduceRpc.ReduceWork(l.ctx, &reducework.ReduceWorkReq{})
	if err != nil {
		return nil, err
	}
	return &pb.GetRpcResp{}, nil
}
