package models

import (
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/apps/entities"
	"github.com/gofiber/fiber/v2"
)

func ListSearch(c *fiber.Ctx, query string) entities.ResponseData {
	collection := GetAyatCollection()
	fields := GetListAyatFields()
	queryBy := []string{
		"name",
		"name_id",
		"name_ltn",
		"text",
		"text_id",
		"text_ltn",
	}
	orders := []string{
		"id_surat:asc",
		"no:asc",
	}

	return GetListDocuments(c, collection, query, fields, queryBy, nil, orders)
}

func GetSearchCollection() string {
	return entities.Ayat.CollectionName(entities.Ayat{})
}
