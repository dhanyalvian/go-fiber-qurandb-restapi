package controllers

import (
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/apps/entities"
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/apps/models"
	"github.com/gofiber/fiber/v2"
)

func GetListAyat(c *fiber.Ctx) error {
	uid := c.Params("uid")
	respData := models.ListAyat(c, uid)

	return RespSucess(c, "", ParseListAyat(respData))
}

func ParseListAyat(respData entities.ResponseData) entities.ResponseData {
	var results []entities.Ayat
	documents := respData.Documents

	if documents.Hits != nil {
		for _, rec := range *documents.Hits {
			doc := *rec.Document
			ayat := ParseListAyatDoc(doc)
			results = append(results, ayat)
		}
	}

	respData.Results = results
	respData.Documents = nil

	return respData
}

func ParseListAyatDoc(doc map[string]interface{}) entities.Ayat {
	return entities.Ayat{
		ID:      GetDocId(doc["id"]),
		IdSurat: GetDocId(doc["id_surat"]),
		No:      GetDocId(doc["no"]),
		Name:    doc["name"].(string),
		NameLtn: doc["name_ltn"].(string),
		NameId:  doc["name_id"].(string),
		Type:    doc["type"].(string),
		Text:    doc["text"].(string),
		TextLtn: doc["text_ltn"].(string),
		TextId:  doc["text_id"].(string),
	}
}
