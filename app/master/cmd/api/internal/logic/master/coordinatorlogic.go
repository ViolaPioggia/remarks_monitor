package master

import (
	"context"
	"fmt"
	"remarks_monitor/app/master/cmd/rpc/master"
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
	fmt.Println(time.Now())
	_, err = l.svcCtx.MasterRpc.GetMap(l.ctx, &master.GetMapReq{})
	fmt.Println(err)
	_, err = l.svcCtx.MasterRpc.GetRpc(l.ctx, &master.GetRpcReq{})
	fmt.Println(err)
	fmt.Println(time.Now())
	return
}
