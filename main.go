package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/prologic/bitcask"
)

var db *bitcask.Bitcask

type shorten struct {
	URL string
}

func shortenHandler(ctx *fiber.Ctx) error {
	var id string

	set := new(shorten)
	err := ctx.BodyParser(&set)

	if err != nil {
		return err
	}

	for {
		id = uuid.NewString()[:5]
		if _, err := db.Get([]byte(id)); err != nil {
			// exit the loop if no such an ID in db
			break
		}
	}

	err = db.Put([]byte(id), []byte(set.URL))

	if err != nil {
		return err
	}

	return ctx.SendString(id)
}

func main() {
	db, _ = bitcask.Open("db")
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("public/index.html")
	})

	app.Post("/api/v1", shortenHandler)

	log.Fatal(app.Listen(":3000"))
}
