package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nqmt/short-url/config"
	"github.com/nqmt/short-url/handler"
	"github.com/nqmt/short-url/service"
)

func setupMiddleware(app *fiber.App) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(logger.New())
	app.Use(recover.New())
}

func main() {
	env := config.GetEnv()
	sv := service.New()
	h := handler.NewFiberHandler(sv)

	app := fiber.New()
	setupMiddleware(app)
	h.SetupFiberRouter(app)

	if err := app.Listen(":" + env.Port); err != nil {
		panic(err)
	}
}
