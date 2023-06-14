package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	DB struct {
		MongoDB struct {
			Url        string
			Db         string
			Collection string
		}
	}
	Cache  cache.CacheConf
	Consul consul.Conf
}
