package master

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"remarks_monitor/app/casbin"
	"remarks_monitor/app/master/cmd/rpc/master"
	"remarks_monitor/app/usercenter/cmd/rpc/usercenter"
	"time"

	"remarks_monitor/app/master/cmd/api/internal/svc"
	"remarks_monitor/app/master/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoordinatorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoordinatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoordinatorLogic {
	return &CoordinatorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoordinatorLogic) Coordinator(req *types.MasterReq) (resp *types.MasterResp, err error) {
	//userId := ctxdata.GetUidFromCtx(l.ctx)
	data, err := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{Id: 2})
	username := data.User.Username
	flag := casbin.CheckPermission(username, "read", "admin")
	if !flag {
		return nil, errors.New("you don't have the access")
	}
	fmt.Println(time.Now())
	_, err = l.svcCtx.MasterRpc.GetMap(l.ctx, &master.GetMapReq{Type: int64(req.Kind)})
	_, err = l.svcCtx.MasterRpc.GetRpc(l.ctx, &master.GetRpcReq{Type: int64(req.Kind)})
	fmt.Println(time.Now())
	return &types.MasterResp{Message: "do MapReduce work success"}, nil
}
