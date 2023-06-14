package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	UserCenterRpcConf zrpc.RpcClientConf
	rest.RestConf
	MasterRpcConf zrpc.RpcClientConf
	JwtAuth       struct {
		AccessSecret string
	}
}
