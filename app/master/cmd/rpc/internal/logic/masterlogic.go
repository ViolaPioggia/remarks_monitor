package logic

import (
	"context"

	"remarks_monitor/app/master/cmd/rpc/internal/svc"
	"remarks_monitor/app/master/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MasterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMasterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MasterLogic {
	return &MasterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MasterLogic) Master(in *pb.WorkReq) (*pb.WorkResp, error) {
	// todo: add your logic here and delete this line

	return &pb.WorkResp{}, nil
}
