package logic

import (
	"context"
	"errors"
	"remarks_monitor/app/casbin"
	"remarks_monitor/app/input/cmd/rpc/input"
	"remarks_monitor/app/usercenter/cmd/rpc/usercenter"
	"remarks_monitor/common/ctxdata"

	"remarks_monitor/app/input/cmd/api/internal/svc"
	"remarks_monitor/app/input/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchAllLogic {
	return &SearchAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchAllLogic) SearchAll(req *types.SearchAllRequest) (resp *types.SearchAllResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	data, err := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{Id: userId})
	username := data.User.Username
	flag := casbin.CheckPermission(username, "read", "admin")
	if !flag {
		return nil, errors.New("you don't have the access")
	}
	searchResp, err := l.svcCtx.InputRpc.Search(l.ctx, &input.SearchReq{Username: req.Username})
	result := searchResp.Data
	var Info []types.Info
	for _, v := range result {
		info := types.Info{
			Username: v.Username,
			Domain:   v.Domain,
			Content:  v.Content,
			Time:     v.Time,
		}
		Info = append(Info, info)
	}
	return &types.SearchAllResponse{Info: Info}, nil
}
