package main

import (
	"flag"
	"fmt"

	"github.com/YuanJun-93/CodeGenesis/internal/config"
	"github.com/YuanJun-93/CodeGenesis/internal/handler"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"

	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "configs/code-genesis-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	// Viper Configuration
	v := viper.New()
	v.SetConfigFile(*configFile)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := v.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
