package models

import (
	"fmt"

	"github.com/dhanyalvian/go-fiber-qurandb-restapi/apps/entities"
	"github.com/gofiber/fiber/v2"
)

func ListAyat(c *fiber.Ctx, uid string) entities.ResponseData {
	collection := GetAyatCollection()
	fields := GetListAyatFields()
	filterBy := []string{
		fmt.Sprintf("id_surat:=%s", uid),
	}
	orders := []string{
		"id_surat:asc",
		"no:asc",
	}

	return GetListDocuments(c, collection, "", fields, nil, filterBy, orders)
}

func GetAyatCollection() string {
	return entities.Ayat.CollectionName(entities.Ayat{})
}

func GetListAyatFields() []string {
	return []string{
		"id",
		"id_surat",
		"no",
		"name",
		"name_ltn",
		"name_id",
		"type",
		"text",
		"text_ltn",
		"text_id",
	}
}
