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
	src := server.NewServer()

	err := fx.New(
		fx.NopLogger, // fx 依赖注入日志，与服务运行日志不同，该选项关闭以来注入日志
		fx.Provide(config.NewConfig, server.NewHTTPServer, server.NewFiber),
		fx.Populate(&src.HttpS),
		fx.Invoke(server.AddHandler),
	).Err()

	src.Group.Go(src.HttpS.Run)
	err = src.Group.Wait()

	return err
}
