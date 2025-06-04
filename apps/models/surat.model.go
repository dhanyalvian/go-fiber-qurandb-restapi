package models

import (
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/apps/entities"
	"github.com/gofiber/fiber/v2"
)

func ListSurat(c *fiber.Ctx) entities.ResponseData {
	collection := getSuratCollection()
	fields := getListSuratFields()
	queryBy := "name_ltn"
	orders := []string{"no:asc"}

	return GetListDocuments(c, collection, "", fields, queryBy, nil, orders)
}

func DetailSurat(c *fiber.Ctx, uid string) entities.ResponseData {
	collection := getSuratCollection()
	id := GetIdBySlug(uid)

	return GetDocument(collection, id)
}

func getSuratCollection() string {
	return entities.Surat.CollectionName(entities.Surat{})
}

func getListSuratFields() []string {
	return []string{
		"id",
		"no",
		"name",
		"name_ltn",
		"name_id",
		"type",
		"total",
	}
}
