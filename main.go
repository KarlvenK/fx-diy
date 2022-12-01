package main

import (
	"fx_diy/pkg/config"
	"fx_diy/pkg/server"

	"go.uber.org/fx"
)

func main() {
	if err := start(); err != nil {
		panic(err)
	}
}

func start() error {

	app := fx.New(
		fx.NopLogger, // fx 依赖注入日志，与服务运行日志不同，该选项关闭以来注入日志
		fx.Provide(config.NewConfig, server.NewHTTPServer, server.NewFiber, server.NewServer),
		fx.Invoke(server.AddHandler),
	)

	app.Run()

	return nil
}
