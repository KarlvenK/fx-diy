package server

import (
	"encoding/json"
	"fmt"
	"fx_diy/pkg/config"
	"github.com/gofiber/fiber/v2"
)

var cnt = 0

type HTTPServer struct {
	app *fiber.App
	cfg *config.Config
}

func NewFiber() *fiber.App {
	cnt++
	fmt.Println(cnt)

	app := fiber.New(
		fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		})
	return app
}

func NewHTTPServer(cfg *config.Config, app *fiber.App) *HTTPServer {

	ret := &HTTPServer{
		app: app,
		cfg: cfg,
	}
	return ret
}

func Start(cfg *config.Config, app *fiber.App) error {
	addr := cfg.ListenAddr()
	return app.Listen(addr)
}

func (h *HTTPServer) Run() error {
	return Start(h.cfg, h.app)
}
