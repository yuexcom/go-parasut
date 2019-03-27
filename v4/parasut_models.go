package parasut

import (
	"fmt"
	"net/http"
)

// relationships ...
type relationships map[string]*relationship

// relationship interface{} çünkü, struct veya slice olabiliyor
type relationship struct {
	Data interface{} `json:"data"`
}

// relationshipData 1:1 ilişki tipi için
type relationshipData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// relationshipDataList 1:m ilişki tipi için
type relationshipDataList []*relationshipData

// -- //

// PageOptions ...
type PageOptions struct {
	Number int `url:"number,omitempty"`
	Size   int `url:"size,omitempty"`
}

// included ...
type included struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Attributes    interface{}   `json:"attributes"`
	Relationships relationships `json:"relationships"`
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

// // RelationshipList ...
// type RelationshipList struct {
// 	Category        Relationship  `json:"category"`
// 	ContactPortal   Relationship  `json:"contact_portal"`
// 	ContactPeople   Relationships `json:"contact_people"`
// 	PriceList       Relationship  `json:"price_list"`
// 	Activities      Relationships `json:"activities"`
// 	EInvoiceInboxes Relationships `json:"e_invoice_inboxes"`
// 	Sharings        Relationships `json:"sharings"`
// }

// // Relationship ...
// type Relationship struct {
// 	Data RelationshipData `json:"data"`
// }

// // Relationships ...
// type Relationships struct {
// 	Data []RelationshipData `json:"data"`
// }

// // RelationshipData ...
// type RelationshipData struct {
// 	ID   string `json:"id"`
// 	Type string `json:"type"`
// }

// // DataMeta ...
// type DataMeta struct {
// 	CreatedAt string `json:"created_at"`
// 	UpdatedAt string `json:"updated_at"`
// }
