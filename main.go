package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/innotechdevops/gomvc/database"
	"github.com/innotechdevops/gomvc/promotion"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	// Solution 1.

	app.Get("/get", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Get!")
	})

	app.Post("/post", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Post!")
	})

	app.Put("/put/:id", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, Put!, id = %s", c.Params("id")))
	})

	app.Patch("/patch/:id", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, Patch!, id = %s", c.Params("id")))
	})

	app.Patch("/patch/:id", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, Patch!, id = %s", c.Params("id")))
	})

	app.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, Patch!, id = %s", c.Params("id")))
	})

	// GORM CRUD
	gormCrud()

	// Solution 2.
	db := database.Connection("database/test.db")
	promotionRepo := promotion.NewRepository(db)
	promotionHandler := promotion.NewHandler(promotionRepo)
	promotionRouter := promotion.NewRouter(promotionHandler)

	// Initial routers
	promotionRouter.Initialize(app)

	log.Fatal(app.Listen(":5001"))
}

func gormCrud() {
	db, err := gorm.Open(sqlite.Open("database/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	_ = db.AutoMigrate(&promotion.Promotion{})

	// Create
	db.Create(&promotion.Promotion{Id: 1, Name: "Get 1 Free 1 by Innotech Development"})

	// Read
	var promo promotion.Promotion
	db.First(&promo, 1)           // find promo with integer primary key
	db.First(&promo, "id = ?", 1) // find promo with code D42

	// Update - update promo's name to Get 1 Free 1
	db.Model(&promo).Update("name", "Get 1 Free 1")
	// Update - update multiple fields
	db.Model(&promo).Updates(promotion.Promotion{Id: 1, Name: "Get 1 Free 1"}) // non-zero fields
	db.Model(&promo).Updates(map[string]interface{}{"id": 1, "name": "Get 1 Free 2"})

	// Delete - delete promo
	db.Delete(&promo, 1)
}
