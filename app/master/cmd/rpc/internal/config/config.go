package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	MapRpcConf1    zrpc.RpcClientConf
	MapRpcConf2    zrpc.RpcClientConf
	MapRpcConf3    zrpc.RpcClientConf
	ReduceRpcConf1 zrpc.RpcClientConf
	ReduceRpcConf2 zrpc.RpcClientConf
	JwtAuth        struct {
		AccessSecret string
		AccessExpire int64
	}
	Consul consul.Conf
}
