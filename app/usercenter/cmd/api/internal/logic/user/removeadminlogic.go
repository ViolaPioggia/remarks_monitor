package user

import (
	"context"
	"errors"
	"remarks_monitor/app/casbin"
	"remarks_monitor/app/usercenter/cmd/rpc/usercenter"
	"remarks_monitor/common/ctxdata"

	"remarks_monitor/app/usercenter/cmd/api/internal/svc"
	"remarks_monitor/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAdminLogic {
	return &RemoveAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveAdminLogic) RemoveAdmin(req *types.RemoveAdminReq) (resp *types.RemoveAdminResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	userInfoResp, err := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	username := userInfoResp.User.Username

	flag := casbin.CheckPermission(username, "write", "admin")
	if !flag {
		return nil, errors.New("you don't have the access")
	}
	casbin.RemovePermission(req.Username, "read", "admin")
	return &types.RemoveAdminResp{Message: "remove new admin success"}, nil
}
