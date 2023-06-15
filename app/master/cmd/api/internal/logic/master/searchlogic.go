package master

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"remarks_monitor/app/casbin"
	"remarks_monitor/app/master/cmd/rpc/master"
	"remarks_monitor/app/usercenter/cmd/rpc/usercenter"
	"remarks_monitor/common/ctxdata"

	"remarks_monitor/app/master/cmd/api/internal/svc"
	"remarks_monitor/app/master/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	data, err := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{Id: userId})
	username := data.User.Username
	flag := casbin.CheckPermission(username, "read", "admin")
	if !flag {
		return nil, errors.New("you don't have the access")
	}
	info, err := l.svcCtx.MasterRpc.Search(l.ctx, &master.SearchReq{
		Type:   int64(req.Kind),
		Offset: int64(req.Offset),
		Num:    int64(req.Num),
	})
	if err != nil {
		logx.Errorf(err.Error())
		return &types.SearchResp{
			Message: "search failed",
			Info:    nil,
		}, err
	}
	var record []types.Info
	_ = copier.Copy(record, info.Record)
	return &types.SearchResp{
		Message: "search success",
		Info:    record,
	}, nil
}
