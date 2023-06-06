package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"remarks_monitor/app/map/cmd/rpc/mapwork"
	"remarks_monitor/app/master/cmd/rpc/internal/config"
	"remarks_monitor/app/reduce/cmd/rpc/reducework"
)

type ServiceContext struct {
	Config    config.Config
	MapRpc    mapwork.MapWork
	ReduceRpc reducework.ReduceWork
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		MapRpc:    mapwork.NewMapWork(zrpc.MustNewClient(c.MapRpcConf)),
		ReduceRpc: reducework.NewReduceWork(zrpc.MustNewClient(c.ReduceRpcConf)),
	}
}
