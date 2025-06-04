package entities

import "github.com/typesense/typesense-go/v3/typesense/api"

type Response struct {
	Meta    ResponseMeta `json:"meta"`
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type ResponseMeta struct {
	RequestId string `json:"req_id"`
	Code      int    `json:"code"`
}

type ResponseData struct {
	Pagination *ResponseDataPagination `json:"pagination,omitempty"`
	Results    interface{}             `json:"results,omitempty"`
	Result     interface{}             `json:"result,omitempty"`
	Documents  *api.SearchResult       `json:"documents,omitempty"`
	Document   interface{}             `json:"document,omitempty"`
}

type ResponseDataPagination struct {
	Page         int   `json:"page"`
	Next         int   `json:"next"`
	Records      int   `json:"records"`
	TotalPages   int   `json:"total_pages"`
	TotalRecords int64 `json:"total_records"`
}
