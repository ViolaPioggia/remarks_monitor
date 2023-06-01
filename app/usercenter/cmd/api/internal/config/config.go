package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	UserCenterRpcConf zrpc.RpcClientConf
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
	}
}
