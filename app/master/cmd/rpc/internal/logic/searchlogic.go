package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"remarks_monitor/app/master/cmd/rpc/internal/svc"
	"remarks_monitor/app/master/cmd/rpc/pb"

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
	records, err := l.svcCtx.RecordsModel.FindRecordsFromBtoS(l.ctx, in.Type, in.Offset, in.Num)
	if err != nil {
		logx.Errorf(err.Error())
	}
	fmt.Println(records)
	var data []*pb.Record
	_ = copier.Copy(&data, records)
	return &pb.SearchResp{Record: data}, nil
}
