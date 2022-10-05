package server

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	Group errgroup.Group
	HttpS *HTTPServer
}

func NewServer(app *HTTPServer, lc fx.Lifecycle) *Server {
	fmt.Println("Excuting NewServer")
	srv := &Server{
		HttpS: app,
		Group: errgroup.Group{},
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				srv.Group.Go(srv.HttpS.Run)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.HttpS.App.Shutdown()
			err := srv.Group.Wait()
			if err != nil {
				return err
			}
			return nil
		},
	})
	return srv
}
