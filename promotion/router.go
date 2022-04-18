package promotion

import (
	"github.com/gofiber/fiber/v2"
)

type Router interface {
	Initialize(app *fiber.App)
}

type router struct {
	Handle Handler
}

func (r *router) Initialize(app *fiber.App) {
	app.Get("/get", r.Handle.Get)
	app.Post("/post", r.Handle.Post)
	app.Put("/put/:id", r.Handle.Put)
	app.Patch("/patch/:id", r.Handle.Patch)
	app.Patch("/patch/:id", r.Handle.Patch)
	app.Delete("/delete/:id", r.Handle.Delete)
}

func NewRouter(handle Handler) Router {
	return &router{handle}
}
