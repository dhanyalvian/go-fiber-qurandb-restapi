package routes

import (
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/apps/controllers"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	app.Get("/surat", controllers.GetListSurat)
	app.Get("/surat/:uid", controllers.GetDetailSurat)
}
