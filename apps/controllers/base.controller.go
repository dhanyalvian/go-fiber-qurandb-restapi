package controllers

import (
	"fmt"

	"github.com/dhanyalvian/go-fiber-qurandb-restapi/apps/entities"
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/packages/response"
	"github.com/gofiber/fiber/v2"
)

func GetUid(c *fiber.Ctx) string {
	return c.Params("uid")
}

func RespSucess(c *fiber.Ctx, message string, data entities.ResponseData) error {
	var resp entities.Response

	resp.Meta.RequestId = response.GetResponseReqId(c)
	resp.Meta.Code = response.GetResponseStatusCode(c)

	resp.Message = message
	resp.Data = data

	return c.Status(fiber.StatusOK).JSON(resp)
}

func RespFailed() {}

func GetDocId(uid interface{}) int {
	id := 0

	if val, ok := uid.(int); ok {
		id = val
	} else if val, ok := uid.(string); ok {
		fmt.Sscanf(val, "%d", &id)
	}

	return id
}
