package logic

import (
	"context"
	"remarks_monitor/app/input/cmd/rpc/input"
	"remarks_monitor/app/usercenter/cmd/rpc/usercenter"
	"remarks_monitor/common/ctxdata"

	"remarks_monitor/app/input/cmd/api/internal/svc"
	"remarks_monitor/app/input/cmd/api/internal/types"

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

func (l *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	data, err := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{Id: userId})
	username := data.User.Username
	searchResp, err := l.svcCtx.InputRpc.Search(l.ctx, &input.SearchReq{Username: username})
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
	return &types.SearchResponse{Info: Info}, nil
}
