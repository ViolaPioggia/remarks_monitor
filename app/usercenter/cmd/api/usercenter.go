package main

import (
	"flag"
	"fmt"
	"remarks_monitor/app/casbin"
	"remarks_monitor/common/tool"

	"remarks_monitor/app/usercenter/cmd/api/internal/config"
	"remarks_monitor/app/usercenter/cmd/api/internal/handler"
	"remarks_monitor/app/usercenter/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var configFile = flag.String("f", tool.GetWD()+"/app/usercenter/cmd/api/etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	casbin.InitCasbin()
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
