package server

import (
	"encoding/json"
	"fmt"
	"fx_diy/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type HTTPServer struct {
	App *fiber.App
	Cfg *config.Config
}

func NewFiber() *fiber.App {
	fmt.Println("Excuting NewFiber")
	app := fiber.New(
		fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		})
	return app
}

func NewHTTPServer(cfg *config.Config, app *fiber.App) *HTTPServer {
	fmt.Println("Excuting NewHTTPServer")
	ret := &HTTPServer{
		App: app,
		Cfg: cfg,
	}
	return ret
}

func Start(cfg *config.Config, app *fiber.App) error {
	addr := cfg.ListenAddr()
	return app.Listen(addr)
}

func (h *HTTPServer) Run() error {
	return Start(h.Cfg, h.App)
}
