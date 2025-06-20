package controllers

import (
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/apps/models"
	"github.com/gofiber/fiber/v2"
)

func GetListSearch(c *fiber.Ctx) error {
	q := c.Query("q")
	respData := models.ListSearch(c, q)

	return RespSucess(c, "", ParseListAyat(respData))
}
