package main

import (
	"flag"
	"fmt"

	"github.com/YuanJun-93/CodeGenesis/internal/config"
	"github.com/YuanJun-93/CodeGenesis/internal/handler"
	"github.com/YuanJun-93/CodeGenesis/internal/pkg/log"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "configs/code-genesis-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// Init Zap Logger
	zapLogger := log.Init(c.ZapLog)
	// Replace go-zero logx writer with Zap
	logx.SetWriter(log.NewZapWriter(zapLogger))

	server := rest.MustNewServer(c.RestConf, rest.WithCors("http://localhost:5173"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
