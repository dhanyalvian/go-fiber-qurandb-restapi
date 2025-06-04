package models

import (
	"context"
	"log"
	"math"
	"strings"

	"github.com/dhanyalvian/go-fiber-qurandb-restapi/apps/entities"
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/configs/databases"
	"github.com/dhanyalvian/go-fiber-qurandb-restapi/packages/request"
	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v3/typesense"
	"github.com/typesense/typesense-go/v3/typesense/api"
	"github.com/typesense/typesense-go/v3/typesense/api/pointer"
)

func GetTs() *typesense.Client {
	return databases.Ts
}

func GetListDocuments(
	c *fiber.Ctx,
	collection string,
	query string,
	fields []string,
	queryBy string,
	filters []string,
	orders []string,
) entities.ResponseData {
	page := request.GetPage(c)
	next := page + 1
	limit := request.GetLimit(c)

	var totalPages int
	var totalRecords int64

	docs, err := SearchParams(
		collection,
		query,
		fields,
		queryBy,
		filters,
		orders,
		page,
		limit,
	)

	if err != nil {
		log.Printf("Search failed: %v\n", err)
	}

	if docs.Found != nil {
		totalRecords = int64(*docs.Found)
	} else {
		totalRecords = 0
	}

	totalPages = int(math.Ceil(float64(totalRecords) / float64(request.GetLimit(c))))
	if next > totalPages {
		next = 0
	}

	records := 0
	if docs.Hits != nil {
		records = len(*docs.Hits)
	}

	return entities.ResponseData{
		Pagination: &entities.ResponseDataPagination{
			Page:         page,
			Next:         next,
			Records:      records,
			TotalPages:   totalPages,
			TotalRecords: totalRecords,
		},
		Documents: docs,
	}
}

func SearchParams(
	collection string,
	query string,
	fields []string,
	queryBy string,
	filters []string,
	orders []string,
	page int,
	limit int,
) (*api.SearchResult, error) {
	ts := GetTs()
	if query == "" {
		query = "*"
	}

	searchParams := &api.SearchCollectionParams{
		Q:             pointer.String(query),
		QueryBy:       pointer.String(queryBy),
		FilterBy:      pointer.String(strings.Join(filters, ",")),
		SortBy:        pointer.String(strings.Join(orders, ",")),
		IncludeFields: pointer.String(strings.Join(fields, ",")),
		Page:          pointer.Int(page),
		PerPage:       pointer.Int(limit),
	}

	return ts.Collection(collection).Documents().Search(context.Background(), searchParams)
}

func GetDocument(collection string, uid string) entities.ResponseData {
	ts := GetTs()
	doc, err := ts.Collection(collection).Document(uid).Retrieve(context.Background())
	if err != nil {
		log.Printf("Search failed: %v\n", err)
	}

	return entities.ResponseData{
		Document: doc,
	}
}

func GetIdBySlug(uid string) string {
	ids := strings.Split(uid, "-")

	return ids[0]
}
