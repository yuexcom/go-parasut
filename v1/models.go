package parasut

import (
	"fmt"
	"net/http"
)

// --------------------------------
// ----- SALES INVOICE MODELS -----
// --------------------------------

// SalesInvoice ...
type SalesInvoice struct {
	ID              int       `json:"id"`
	Description     string    `json:"description"`
	IssueDate       string    `json:"issue_date"`
	DueDate         string    `json:"due_date"`
	GrossTotal      string    `json:"gross_total"`
	TotalDiscount   string    `json:"total_discount"`
	TotalVat        string    `json:"total_vat"`
	CreatedAt       string    `json:"created_at"`
	ItemType        string    `json:"item_type"`
	Remaining       string    `json:"remaining"`
	TotalPaids      string    `json:"total_paids"`
	InvoiceSeries   string    `json:"invoice_series"`
	InvoiceID       int       `json:"invoice_id"`
	InvoiceNo       string    `json:"invoice_no"`
	PrintedAt       string    `json:"printed_at"`
	Archiced        bool      `json:"archiced"`
	BillingAddress  string    `json:"billing_address"`
	BillingPhone    string    `json:"billing_phone"`
	BillingFax      string    `json:"billing_fax"`
	TaxOffice       string    `json:"tax_office"`
	TaxNumber       string    `json:"tax_number"`
	ShipmentDate    string    `json:"shipment_date"`
	NetTotal        string    `json:"net_total"`
	PaymentStatus   string    `json:"payment_status"`
	DaysOverdue     int       `json:"days_overdue"`
	DaysOutstanding int       `json:"days_outstanding"`
	Category        Category  `json:"category"`
	Tags            []Tag     `json:"tags"`
	Contact         Contact   `json:"contact"`
	Details         []Detail  `json:"details"`
	Payments        []Payment `json:"payments"`
}

// SalesInvoiceListResponse ...
type SalesInvoiceListResponse struct {
	Items []SalesInvoice `json:"items"`
	Meta  Meta           `json:"meta"`
}

// SalesInvoiceCreateRequest ...
type SalesInvoiceCreateRequest struct {
	SalesInvoice SalesInvoiceRequest `json:"sales_invoice"`
}

// SalesInvoiceUpdateRequest ...
// TODO: "details_attributes" alanı update etmiyor ekliyor, destek al
type SalesInvoiceUpdateRequest struct {
	SalesInvoice SalesInvoiceRequest `json:"sales_invoice"`
}

// SalesInvoiceResponse ...
type SalesInvoiceResponse struct {
	SalesInvoice SalesInvoice `json:"sales_invoice"`
}

// SalesInvoiceRequest ...
type SalesInvoiceRequest struct {
	Description       string             `json:"description"`
	InvoiceID         int                `json:"invoice_id"`
	InvoiceSeries     string             `json:"invoice_series"`
	ItemType          string             `json:"item_type"`
	IssueDate         string             `json:"issue_date"`
	DueDate           string             `json:"due_date"`
	CategoryID        int                `json:"category_id"`
	ContactID         int                `json:"contact_id"`
	BillingAddress    string             `json:"billing_address"`
	BillingFax        string             `json:"billing_fax"`
	BillingPhone      string             `json:"billing_phone"`
	TaxNumber         string             `json:"tax_number"`
	TaxOffice         string             `json:"tax_office"`
	Archiced          bool               `json:"archiced"`
	DetailsAttributes []DetailsAttribute `json:"details_attributes"`
}

// DetailsAttribute ...
type DetailsAttribute struct {
	ProductID             int     `json:"product_id,omitempty"`
	Quantity              float64 `json:"quantity"`
	UnitPrice             float64 `json:"unit_price"`
	VatRate               float64 `json:"vat_rate"`
	DiscountType          string  `json:"discount_type"`
	DiscountValue         float64 `json:"discount_value"`
	ExciseDutyType        string  `json:"excise_duty_type"`
	ExciseDutyValue       float64 `json:"excise_duty_value"`
	CommunicationsTaxRate float64 `json:"communications_tax_rate"`
}

// Tag ...
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Detail ...
type Detail struct {
	ID           int     `json:"id"`
	DetailNo     int     `json:"detail_no"`
	Quantity     string  `json:"quantity"`
	UnitPrice    string  `json:"unit_price"`
	Discount     string  `json:"discount"`
	VatRate      string  `json:"vat_rate"`
	DiscountType string  `json:"discount_type"`
	DiscountRate string  `json:"discount_rate"`
	Product      Product `json:"product"`
}

// Payment ...
type Payment struct {
	ID          int    `json:"id"`
	Date        string `json:"date"`
	PayableID   int    `json:"payable_id"`
	PayableType string `json:"payable_type"`
	Amount      string `json:"amount"`
	Notes       string `json:"notes"`
	Flow        string `json:"flow"`
	IsOverdue   bool   `json:"is_overdue"`
	IsPaid      bool   `json:"is_paid"`
	PaidOn      string `json:"paid_on"`
}

// Product ...
type Product struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	VatRate   string `json:"vat_rate"`
	Currency  string `json:"currency"`
	ListPrice string `json:"list_price"`
}

// SalesInvoiceDeleteResponse ...
type SalesInvoiceDeleteResponse struct {
	Success string `json:"success"`
}

// --------------------------
// ----- CONTACT MODELS -----
// --------------------------

// Contact ...
type Contact struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	Email           string          `json:"email"`
	ContactType     string          `json:"contact_type"`
	TaxNumber       string          `json:"tax_number"`
	TaxOffice       string          `json:"tax_office"`
	Balance         string          `json:"balance"`
	EstimateBalance string          `json:"estimate_balance"`
	Archived        bool            `json:"archived"`
	District        string          `json:"district"`
	City            string          `json:"city"`
	Phone           string          `json:"phone"`
	ContactPeople   []ContactPerson `json:"contact_people"`
	Address         Address         `json:"address"`
	Category        Category        `json:"category"`
}

// ContactListResponse ...
type ContactListResponse struct {
	Items []Contact `json:"items"`
	Meta  Meta      `json:"meta"`
}

// ContactResponse ...
type ContactResponse struct {
	Contact Contact `json:"contact"`
}

// ContactCreateRequest ...
type ContactCreateRequest struct {
	Contact ContactRequest `json:"contact"`
}

// ContactUpdateRequest ...
// TODO: "contact_people_attributes" alanı update etmiyor ekliyor, destek al
type ContactUpdateRequest struct {
	Contact ContactRequest `json:"contact"`
}

// ContactRequest ...
type ContactRequest struct {
	Name                    string          `json:"name"`
	ContactType             string          `json:"contact_type"`
	Email                   string          `json:"emails"`
	TaxNumber               string          `json:"tax_number"`
	TaxOffice               string          `json:"tax_office"`
	CategoryID              int             `json:"category_id,omitempty"`
	City                    string          `json:"city"`
	District                string          `json:"district"`
	AddressAttributes       Address         `json:"address_attributes"`
	ContactPeopleAttributes []ContactPerson `json:"contact_people_attributes"`
}

// ContactPerson ...
type ContactPerson struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Notes string `json:"notes"`
}

// ContactDeleteResponse ...
type ContactDeleteResponse struct {
	Success string `json:"success"`
}

// -------------------------
// ----- COMMON MODELS -----
// -------------------------

// ListOptions ...
type ListOptions struct {
	LastSync  string `url:"last_synch,omitempty"`
	PerPage   int    `url:"per_page,omitempty"`
	Page      int    `url:"page,omitempty"`
	TaxNumber string `url:"tax_number,omitempty"`
}

// Category ...
type Category struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	BgColor   string `json:"bg_color"`
	TextColor string `json:"text_color"`
}

// Address ...
type Address struct {
	Address string `json:"address"`
	Fax     string `json:"fax"`
	Phone   string `json:"phone"`
}

// Meta ...
type Meta struct {
	TotalCount int `json:"total_count"`
	PageCount  int `json:"page_count"`
	PerPage    int `json:"per_page"`
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
	SalesInvoiceID string   `json:"sales_invoice_id"`
	Type           string   `json:"type"`
	Status         string   `json:"status"`
	Errors         []string `json:"errors"`
	PDF            string   `json:"pdf"`
}

func (r *EDocumentStatusResponseError) Error() string {
	return fmt.Sprintf("%+v [Fatura: #%s]", r.Errors, r.SalesInvoiceID)
}
