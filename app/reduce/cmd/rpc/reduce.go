package main

import (
	"flag"
	"fmt"
	"remarks_monitor/common/tool"

	"remarks_monitor/app/reduce/cmd/rpc/internal/config"
	"remarks_monitor/app/reduce/cmd/rpc/internal/server"
	"remarks_monitor/app/reduce/cmd/rpc/internal/svc"
	"remarks_monitor/app/reduce/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", tool.GetWD()+"/app/reduce/cmd/rpc/etc/reduce.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterReduceWorkServer(grpcServer, server.NewReduceWorkServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
