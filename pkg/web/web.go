package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

type ApiServer struct {
	app *fiber.App
}

func NewApiServer() *ApiServer {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(helmet.New())

	srv := &ApiServer{
		app: app,
	}
	srv.Register("/healthcheck", healthHandler())
	return srv
}

func (a *ApiServer) Listen(addr string) error {
	return a.app.Listen(addr)
}

func (a *ApiServer) Register(path string, handlers ...fiber.Handler) {
	a.app.Get(path, handlers...)
}

func healthHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Send([]byte("healthy"))
	}
}
