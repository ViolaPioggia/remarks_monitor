package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"remarks_monitor/app/input/cmd/rpc/internal/config"
	"remarks_monitor/app/input/model"
)

type ServiceContext struct {
	Config       config.Config
	RedisClient  *redis.Redis
	RemarksModel model.RemarksModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		RemarksModel: model.NewRemarksModel(c.DB.MongoDB.Url, c.DB.MongoDB.Db, c.DB.MongoDB.Collection, c.Cache),
	}
}
