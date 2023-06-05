package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"remarks_monitor/app/map/cmd/rpc/mapwork"
	"remarks_monitor/app/master/cmd/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	MapRpc mapwork.MapWork
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		MapRpc: mapwork.NewMapWork(zrpc.MustNewClient(c.MapRpcConf)),
	}
}
