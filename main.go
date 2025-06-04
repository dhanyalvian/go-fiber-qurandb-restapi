package main

import (
	"encoding/json"
	"log"

	"github.com/dhanyalvian/go-fiber-qurandb-restapi/configs"
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/initializations"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := configs.Get()
	app := fiber.New(fiber.Config{
		AppName:     cfg.Server.AppName,
		Prefork:     cfg.Server.Prefork,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	initializations.Init(app)

	log.Fatal(app.Listen(":" + cfg.Server.Port))
}
