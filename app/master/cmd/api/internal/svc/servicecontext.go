package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"remarks_monitor/app/master/cmd/api/internal/config"
	"remarks_monitor/app/master/cmd/rpc/master"
	"remarks_monitor/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	UserCenterRpc usercenter.Usercenter
	MasterRpc     master.Master
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		MasterRpc:     master.NewMaster(zrpc.MustNewClient(c.MasterRpcConf)),
		UserCenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
	}
}
