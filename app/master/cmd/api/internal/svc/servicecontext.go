package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"remarks_monitor/app/master/cmd/api/internal/config"
	"remarks_monitor/app/master/cmd/rpc/master"
)

type ServiceContext struct {
	Config    config.Config
	MasterRpc master.Master
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		MasterRpc: master.NewMaster(zrpc.MustNewClient(c.MasterRpcConf)),
	}
}
