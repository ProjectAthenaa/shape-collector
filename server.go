package main

import (
	"context"
	"github.com/ProjectAthenaa/sonic-core/sonic/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("/index.html")
	})

	app.Get("/base.js", func(c *fiber.Ctx) error {
		return c.SendFile("/base.js")
	})

	app.Get("/dynamic.js", func(c *fiber.Ctx) error {
		return c.SendFile("/dynamic.js")
	})

	app.Post("/data", func(ctx *fiber.Ctx) error {
		core.Base.GetRedis("cache").SAdd(context.Background(), "shape:newbalance:collector", string(ctx.Body()))
		return ctx.JSON(fiber.Map{"success": true})
	})

	//var a=new XMLHttpRequest;a.open("POST","/submit",!0),a.send();setTimeout(()=>{console.log(JSON.stringify(tagobj))},500);
	app.Post("/submit", func(c *fiber.Ctx) error {
		return c.SendString("nice job")
	})

	app.Listen(":8080")
}