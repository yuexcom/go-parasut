package parasut

import (
	"fmt"
	"net/http"
)

// Relationships ...
type Relationships map[string]*Relationship

// Relationship interface{} çünkü, struct veya slice olabiliyor
type Relationship struct {
	Data interface{} `json:"data"`
}

// RelationshipData 1:1 ilişki tipi için
type RelationshipData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// RelationshipDataList 1:m ilişki tipi için
type RelationshipDataList []*RelationshipData

// -- //

// PageOptions ...
type PageOptions struct {
	Number int `url:"number,omitempty"`
	Size   int `url:"size,omitempty"`
}

// Included ...
type Included struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Attributes    interface{}   `json:"attributes"`
	Relationships Relationships `json:"Relationships"`
}

// Meta ...
type Meta struct {
	CurrentPage      int    `json:"current_page"`
	TotalPages       int    `json:"total_pages"`
	TotalCount       int    `json:"total_count"`
	PerPage          int    `json:"per_page"`
	PayableTotal     string `json:"payable_total"`
	CollectibleTotal string `json:"collectible_total"`
	ExportURL        string `json:"export_url"`
}

// ResponseError wrapper
type ResponseError struct {
	Response       *http.Response
	Errors         []Error  `json:"errors"`
	Base           []string `json:"base"`
	SalesInvoiceID string   `json:"sales_invoice_id"`
	Type           string   `json:"type"`
	Status         string   `json:"status"`
	PDF            string   `json:"pdf"`
}

// Error ...
type Error struct {
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func (r *ResponseError) Error() string {
	if len(r.Base) > 0 {
		return fmt.Sprintf("%+v", r.Base)
	}
	return fmt.Sprintf("%+v", r.Errors)
}

// CommonResponseError ...
type CommonResponseError struct {
	Errors []Error `json:"errors"`
}

func (r *CommonResponseError) Error() string {
	return fmt.Sprintf("%+v", r.Errors)
}

// BaseResponseError ...
type BaseResponseError struct {
	Base []string `json:"base"`
}

func (r *BaseResponseError) Error() string {
	return fmt.Sprintf("%+v", r.Base)
}

// EDocumentStatusResponseError ...
type EDocumentStatusResponseError struct {
	SalesInvoiceID string  `json:"sales_invoice_id"`
	Type           string  `json:"type"`
	Status         string  `json:"status"`
	Errors         []Error `json:"errors"`
	PDF            string  `json:"pdf"`
}

func (r *EDocumentStatusResponseError) Error() string {
	return fmt.Sprintf("%+v [Fatura: #%s]", r.Errors, r.SalesInvoiceID)
}
