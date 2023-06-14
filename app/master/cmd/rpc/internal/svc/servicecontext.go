package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	mapwork1 "remarks_monitor/app/map1/cmd/rpc/mapwork"
	mapwork2 "remarks_monitor/app/map2/cmd/rpc/mapwork"
	mapwork3 "remarks_monitor/app/map3/cmd/rpc/mapwork"
	"remarks_monitor/app/master/cmd/rpc/internal/config"
	"remarks_monitor/app/master/model"
	"remarks_monitor/app/reduce1/cmd/rpc/reducework1"
	"remarks_monitor/app/reduce2/cmd/rpc/reducework2"
)

type ServiceContext struct {
	Config       config.Config
	MapRpc1      mapwork1.MapWork
	MapRpc2      mapwork2.MapWork
	MapRpc3      mapwork3.MapWork
	ReduceRpc1   reducework1.ReduceWork1
	ReduceRpc2   reducework2.ReduceWork2
	RecordsModel model.RecordModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		MapRpc1:      mapwork1.NewMapWork(zrpc.MustNewClient(c.MapRpcConf1)),
		MapRpc2:      mapwork2.NewMapWork(zrpc.MustNewClient(c.MapRpcConf2)),
		MapRpc3:      mapwork3.NewMapWork(zrpc.MustNewClient(c.MapRpcConf3)),
		ReduceRpc1:   reducework1.NewReduceWork1(zrpc.MustNewClient(c.ReduceRpcConf1)),
		ReduceRpc2:   reducework2.NewReduceWork2(zrpc.MustNewClient(c.ReduceRpcConf2)),
		RecordsModel: model.NewRecordModel(c.DB.MongoDB.Url, c.DB.MongoDB.Db, c.DB.MongoDB.Collection, c.Cache),
	}
}
