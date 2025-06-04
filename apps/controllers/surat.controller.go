package controllers

import (
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/apps/entities"
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/apps/models"
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/packages/conversion"
	"github.com/gofiber/fiber/v2"
)

func GetListSurat(c *fiber.Ctx) error {
	respData := models.ListSurat(c)

	return RespSucess(c, "", ParseListSurat(respData))
}

func GetDetailSurat(c *fiber.Ctx) error {
	uid := c.Params("uid")
	respData := models.DetailSurat(c, uid)

	return RespSucess(c, "", ParseDetailSurat(respData))
}

func ParseListSurat(respData entities.ResponseData) entities.ResponseData {
	var results []entities.Surat
	documents := respData.Documents

	if documents.Hits != nil {
		for _, rec := range *documents.Hits {
			doc := *rec.Document
			surat := ParseListSuratDoc(doc)
			results = append(results, surat)
		}
	}

	respData.Results = results
	respData.Documents = nil

	return respData
}

func ParseDetailSurat(respData entities.ResponseData) entities.ResponseData {
	var surat entities.Surat
	var suratDetail entities.SuratDetail

	document := respData.Document

	if document != nil {
		doc := document.(map[string]interface{})
		surat = ParseListSuratDoc(doc)
		suratDetail = ParseDetailSuratDoc(doc, surat)
	}

	respData.Result = suratDetail
	respData.Document = nil

	return respData
}

func ParseListSuratDoc(doc map[string]interface{}) entities.Surat {
	id := GetDocId(doc["id"])

	return entities.Surat{
		ID:      id,
		No:      id,
		Name:    doc["name"].(string),
		NameLtn: doc["name_ltn"].(string),
		NameId:  doc["name_id"].(string),
		Type:    doc["type"].(string),
		Total:   conversion.Float64ToInt(doc["total"]),
	}
}

func ParseDetailSuratDoc(doc map[string]interface{}, surat entities.Surat) entities.SuratDetail {
	return entities.SuratDetail{
		Surat:       surat,
		Description: doc["description"].(string),
		Audio:       doc["audio"].(string),
	}
}
