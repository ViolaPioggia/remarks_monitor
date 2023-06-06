package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	MapRpcConf    zrpc.RpcClientConf
	ReduceRpcConf zrpc.RpcClientConf
	JwtAuth       struct {
		AccessSecret string
		AccessExpire int64
	}
	Consul consul.Conf
}
