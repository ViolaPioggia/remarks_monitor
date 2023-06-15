package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"remarks_monitor/app/input/cmd/rpc/input"

	"remarks_monitor/app/input/cmd/rpc/internal/svc"
	"remarks_monitor/app/input/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLogic) Search(in *pb.SearchReq) (*pb.SearchResp, error) {
	username := in.Username
	remarks, err := l.svcCtx.RemarksModel.FindByUsername(l.ctx, username)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	var data []*input.InputReq
	_ = copier.Copy(&data, remarks)
	return &pb.SearchResp{Data: data}, nil
}
