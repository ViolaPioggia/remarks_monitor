package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"
	"remarks_monitor/app/casbin"
	"remarks_monitor/app/master/cmd/api/internal/config"
	"remarks_monitor/app/master/cmd/api/internal/handler"
	"remarks_monitor/app/master/cmd/api/internal/svc"
	"remarks_monitor/common/tool"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var configFile = flag.String("f", tool.GetWD()+"/app/master/cmd/api/etc/master.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	casbin.InitCasbin()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
