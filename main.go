package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/google/uuid"
	_ "github.com/prologic/bitcask"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("public/index.html")
	})

	log.Fatal(app.Listen(":3000"))
}
