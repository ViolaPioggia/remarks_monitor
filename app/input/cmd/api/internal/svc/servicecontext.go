package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"remarks_monitor/app/input/cmd/api/internal/config"
	"remarks_monitor/app/input/cmd/rpc/input"
	"remarks_monitor/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	UserCenterRpc usercenter.Usercenter
	InputRpc      input.Input
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserCenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
		InputRpc:      input.NewInput(zrpc.MustNewClient(c.InputRpcConf)),
	}
}
