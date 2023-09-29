package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

type ApiServer struct {
	app *fiber.App
}

func NewApiServer(mode string) *ApiServer {
	apiServer := new(ApiServer)
	app := fiber.New()
	app.Use(cors.New())
	app.Use(helmet.New())

	apiServer.app = app
	apiServer.Register(mode)
	return apiServer
}

func (a *ApiServer) Listen(addr string) error {
	return a.app.Listen(addr)
}

func (a *ApiServer) Register(mode string) {
	a.app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Send([]byte("healthy"))
	})
}
