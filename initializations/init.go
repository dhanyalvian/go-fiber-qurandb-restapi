package initializations

import (
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/configs/databases"
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/configs/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Init(app *fiber.App) {
	initCompress(app)
	initCors(app)
	// initCsrf(app)
	initHelmet(app)
	initRequestId(app)
	initLogger(app)

	databases.Init()
	routes.Init(app)
}

func initCompress(app *fiber.App) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
}

func initCors(app *fiber.App) {
	app.Use(cors.New())
}

func initCsrf(app *fiber.App) {
	app.Use(csrf.New())
}

func initHelmet(app *fiber.App) {
	app.Use(helmet.New())
}

func initRequestId(app *fiber.App) {
	app.Use(requestid.New())
}

func initLogger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${host} | ${method} | ${path} | ${queryParams}\n",
		TimeFormat: "2006-01-26 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))
}
