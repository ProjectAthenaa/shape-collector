package main

import (
	"github.com/ProjectAthenaa/sonic-core/sonic/core"
	"context"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./scripts/index.html")
	})

	app.Post("/data", func(ctx *fiber.Ctx) error {
		core.Base.GetRedis("cache").SAdd(context.Background(), "shape:newbalance:collector", string(ctx.Body()))
		return ctx.JSON(fiber.Map{"success": true})
	})

	app.Post("/submit", func(c *fiber.Ctx) error {
		return c.SendString("nice job")
	})

	app.Static("/", "./scripts")

	app.Listen(":3000")
}