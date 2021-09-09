package main

import (
	"context"
	"github.com/ProjectAthenaa/sonic-core/sonic/core"
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

	//var a=new XMLHttpRequest;a.open("POST","https://www.newbalance.com/on/demandware.store/Sites-NBUS-Site/en_US/Cart-AddProduct",!0),a.send(),console.log(JSON.stringify(tagobj));
	app.Post("/submit", func(c *fiber.Ctx) error {
		return c.SendString("nice job")
	})

	app.Static("/", "./scripts")

	app.Listen(":8080")
}