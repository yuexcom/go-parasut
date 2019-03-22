package parasut

import (
	"fmt"
	"net/http"
)

// PageOptions ...
type PageOptions struct {
	Number int `url:"number,omitempty"`
	Size   int `url:"size,omitempty"`
}

// Included ...
type Included struct {
	ID            string           `json:"id,omitempty"`
	Type          string           `json:"type,omitempty"`
	Attributes    interface{}      `json:"attributes"`
	Relationships RelationshipList `json:"relationships"`
}

// RelationshipList ...
type RelationshipList struct {
	Category        Relationship  `json:"category"`
	ContactPortal   Relationship  `json:"contact_portal"`
	ContactPeople   Relationships `json:"contact_people"`
	PriceList       Relationship  `json:"price_list"`
	Activities      Relationships `json:"activities"`
	EInvoiceInboxes Relationships `json:"e_invoice_inboxes"`
	Sharings        Relationships `json:"sharings"`
}

// Relationship ...
type Relationship struct {
	Data RelationshipData `json:"data"`
}

// Relationships ...
type Relationships struct {
	Data []RelationshipData `json:"data"`
}

// RelationshipData ...
type RelationshipData struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
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

// DataMeta ...
type DataMeta struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ResponseError wrapper
type ResponseError struct {
	Response       *http.Response
	Errors         []string `json:"errors,omitempty"`
	Base           []string `json:"base,omitempty"`
	SalesInvoiceID string   `json:"sales_invoice_id,omitempty"`
	Type           string   `json:"type,omitempty"`
	Status         string   `json:"status,omitempty"`
	PDF            string   `json:"pdf,omitempty"`
}

func (r *ResponseError) Error() string {
	if len(r.Base) > 0 {
		return fmt.Sprintf("%+v", r.Base)
	}
	return fmt.Sprintf("%+v", r.Errors)
}

// CommonResponseError ...
type CommonResponseError struct {
	Errors []string `json:"errors"`
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
	SalesInvoiceID string   `json:"sales_invoice_id,omitempty"`
	Type           string   `json:"type,omitempty"`
	Status         string   `json:"status,omitempty"`
	Errors         []string `json:"errors,omitempty"`
	PDF            string   `json:"pdf,omitempty"`
}

func (r *EDocumentStatusResponseError) Error() string {
	return fmt.Sprintf("%+v [Fatura: #%s]", r.Errors, r.SalesInvoiceID)
}
