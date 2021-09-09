package main

import (
	"context"
	"flag"
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/ProjectAthenaa/sonic-core/sonic/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})



	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	app.Post("/data", func(ctx *fiber.Ctx) error {
		_, _ = core.Base.GetPg("cache").Device.Create().SetAdevice(sonic.Map{"data": string(ctx.Body())}).Save(ctx.UserContext())
		core.Base.GetRedis("cache").SAdd(context.Background(), "shape:newbalance:collector", string(ctx.Body()))
		return ctx.JSON(fiber.Map{"success": true})
	})

	//var a=new XMLHttpRequest;a.open("POST","/submit",!0),a.send();setTimeout(()=>{console.log(JSON.stringify(tagobj))},500);
	app.Post("/submit", func(c *fiber.Ctx) error {
		return c.SendString("nice job")
	})

	// Setup static files
	app.Static("/", "./static/public")

	// Listen on port 3000
	app.Listen(":8080") // go run app.go -port=:3000
}
