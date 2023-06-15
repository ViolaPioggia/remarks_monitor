package user

import (
	"context"
	"github.com/pkg/errors"
	"remarks_monitor/app/casbin"
	"remarks_monitor/app/usercenter/cmd/rpc/usercenter"
	"remarks_monitor/common/ctxdata"

	"remarks_monitor/app/usercenter/cmd/api/internal/svc"
	"remarks_monitor/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAdminLogic {
	return &AddAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAdminLogic) AddAdmin(req *types.AddAdminReq) (resp *types.AddAdminResp, err error) {
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
	casbin.AddPermission(req.Username, "read", "admin")
	return &types.AddAdminResp{Message: "add new admin success"}, nil
}
