package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"remarks_monitor/common/tool"

	"remarks_monitor/app/reduce1/cmd/rpc/internal/config"
	"remarks_monitor/app/reduce1/cmd/rpc/internal/server"
	"remarks_monitor/app/reduce1/cmd/rpc/internal/svc"
	"remarks_monitor/app/reduce1/cmd/rpc/reduce1"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", tool.GetWD()+"/app/reduce1/cmd/rpc/etc/reduce1.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		reduce1.RegisterReduceWork1Server(grpcServer, server.NewReduceWork1Server(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	_ = consul.RegisterService(c.ListenOn, c.Consul)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
