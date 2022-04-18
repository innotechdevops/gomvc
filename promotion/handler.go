package promotion

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Get(c *fiber.Ctx) error
	Post(c *fiber.Ctx) error
	Put(c *fiber.Ctx) error
	Patch(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type handler struct {
	Repo Repository
}

func (h *handler) Get(c *fiber.Ctx) error {
	return c.SendString("Hello, Get!")
}

func (h *handler) Post(c *fiber.Ctx) error {
	return c.SendString("Hello, Post!")
}

func (h *handler) Put(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("Hello, Put!, id = %s", c.Params("id")))
}

func (h *handler) Patch(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("Hello, Patch!, id = %s", c.Params("id")))
}

func (h *handler) Delete(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("Hello, Patch!, id = %s", c.Params("id")))
}

func NewHandler(repo Repository) Handler {
	return &handler{repo}
}
