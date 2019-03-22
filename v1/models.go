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
	ID              int       `json:"id,omitempty"`
	Description     string    `json:"description,omitempty"`
	IssueDate       string    `json:"issue_date,omitempty"`
	DueDate         string    `json:"due_date,omitempty"`
	GrossTotal      string    `json:"gross_total,omitempty"`
	TotalDiscount   string    `json:"total_discount,omitempty"`
	TotalVat        string    `json:"total_vat,omitempty"`
	CreatedAt       string    `json:"created_at,omitempty"`
	ItemType        string    `json:"item_type,omitempty"`
	Remaining       string    `json:"remaining,omitempty"`
	TotalPaids      string    `json:"total_paids,omitempty"`
	InvoiceSeries   string    `json:"invoice_series,omitempty"`
	InvoiceID       int       `json:"invoice_id,omitempty"`
	InvoiceNo       string    `json:"invoice_no,omitempty"`
	PrintedAt       string    `json:"printed_at,omitempty"`
	Archiced        bool      `json:"archiced,omitempty"`
	BillingAddress  string    `json:"billing_address,omitempty"`
	BillingPhone    string    `json:"billing_phone,omitempty"`
	BillingFax      string    `json:"billing_fax,omitempty"`
	TaxOffice       string    `json:"tax_office,omitempty"`
	TaxNumber       string    `json:"tax_number,omitempty"`
	ShipmentDate    string    `json:"shipment_date,omitempty"`
	NetTotal        string    `json:"net_total,omitempty"`
	PaymentStatus   string    `json:"payment_status,omitempty"`
	DaysOverdue     int       `json:"days_overdue,omitempty"`
	DaysOutstanding int       `json:"days_outstanding,omitempty"`
	Category        Category  `json:"category,omitempty"`
	Tags            []Tag     `json:"tags,omitempty"`
	Contact         Contact   `json:"contact,omitempty"`
	Details         []Detail  `json:"details,omitempty"`
	Payments        []Payment `json:"payments,omitempty"`
}

// SalesInvoiceListResponse ...
type SalesInvoiceListResponse struct {
	Items []SalesInvoice `json:"items,omitempty"`
	Meta  Meta           `json:"meta,omitempty"`
}

// SalesInvoiceCreateRequest ...
type SalesInvoiceCreateRequest struct {
	SalesInvoice SalesInvoiceRequest `json:"sales_invoice,omitempty"`
}

// SalesInvoiceUpdateRequest ...
// TODO: "details_attributes" alanı update etmiyor ekliyor, destek al
type SalesInvoiceUpdateRequest struct {
	SalesInvoice SalesInvoiceRequest `json:"sales_invoice,omitempty"`
}

// SalesInvoiceResponse ...
type SalesInvoiceResponse struct {
	SalesInvoice SalesInvoice `json:"sales_invoice,omitempty"`
}

// SalesInvoiceRequest ...
type SalesInvoiceRequest struct {
	Description       string             `json:"description,omitempty"`
	InvoiceID         int                `json:"invoice_id,omitempty"`
	InvoiceSeries     string             `json:"invoice_series,omitempty"`
	ItemType          string             `json:"item_type,omitempty"`  // invoice (fatura), estimate (proforma)
	IssueDate         string             `json:"issue_date,omitempty"` // YYYY-MM-DD
	DueDate           string             `json:"due_date,omitempty"`   // YYYY-MM-DD
	CategoryID        int                `json:"category_id,omitempty"`
	ContactID         int                `json:"contact_id,omitempty"`
	BillingAddress    string             `json:"billing_address,omitempty"`
	BillingFax        string             `json:"billing_fax,omitempty"`
	BillingPhone      string             `json:"billing_phone,omitempty"`
	TaxNumber         string             `json:"tax_number,omitempty"`
	TaxOffice         string             `json:"tax_office,omitempty"`
	Archiced          bool               `json:"archiced,omitempty"`
	DetailsAttributes []DetailsAttribute `json:"details_attributes,omitempty"`
}

// DetailsAttribute ...
type DetailsAttribute struct {
	ProductID             int     `json:"product_id,omitempty"`
	Quantity              float64 `json:"quantity,omitempty"`
	UnitPrice             float64 `json:"unit_price,omitempty"`
	VatRate               float64 `json:"vat_rate,omitempty"`
	DiscountType          string  `json:"discount_type,omitempty"` // amount
	DiscountValue         float64 `json:"discount_value,omitempty"`
	ExciseDutyType        string  `json:"excise_duty_type,omitempty"` // percentage
	ExciseDutyValue       float64 `json:"excise_duty_value,omitempty"`
	CommunicationsTaxRate float64 `json:"communications_tax_rate,omitempty"`
}

// Tag ...
type Tag struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Detail ...
type Detail struct {
	ID           int     `json:"id,omitempty"`
	DetailNo     int     `json:"detail_no,omitempty"`
	Quantity     string  `json:"quantity,omitempty"`
	UnitPrice    string  `json:"unit_price,omitempty"`
	Discount     string  `json:"discount,omitempty"`
	VatRate      string  `json:"vat_rate,omitempty"`
	DiscountType string  `json:"discount_type,omitempty"`
	DiscountRate string  `json:"discount_rate,omitempty"`
	Product      Product `json:"product,omitempty"`
}

// Payment ...
type Payment struct {
	ID          int    `json:"id,omitempty"`
	Date        string `json:"date,omitempty"`
	PayableID   int    `json:"payable_id,omitempty"`
	PayableType string `json:"payable_type,omitempty"`
	Amount      string `json:"amount,omitempty"`
	Notes       string `json:"notes,omitempty"`
	Flow        string `json:"flow,omitempty"`
	IsOverdue   bool   `json:"is_overdue,omitempty"`
	IsPaid      bool   `json:"is_paid,omitempty"`
	PaidOn      string `json:"paid_on,omitempty"`
}

// Product ...
type Product struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	VatRate   string `json:"vat_rate,omitempty"`
	Currency  string `json:"currency,omitempty"`
	ListPrice string `json:"list_price,omitempty"`
}

// SalesInvoiceDeleteResponse ...
type SalesInvoiceDeleteResponse struct {
	Success string `json:"success,omitempty"`
}

// --------------------------
// ----- CONTACT MODELS -----
// --------------------------

// Contact ...
type Contact struct {
	ID              int             `json:"id,omitempty"`
	Name            string          `json:"name,omitempty"`
	Email           string          `json:"email,omitempty"`
	ContactType     string          `json:"contact_type,omitempty"`
	TaxNumber       string          `json:"tax_number,omitempty"`
	TaxOffice       string          `json:"tax_office,omitempty"`
	Balance         string          `json:"balance,omitempty"`
	EstimateBalance string          `json:"estimate_balance,omitempty"`
	Archived        bool            `json:"archived,omitempty"`
	District        string          `json:"district,omitempty"`
	City            string          `json:"city,omitempty"`
	Phone           string          `json:"phone,omitempty"`
	ContactPeople   []ContactPerson `json:"contact_people,omitempty"`
	Address         Address         `json:"address,omitempty"`
	Category        Category        `json:"category,omitempty"`
}

// ContactListResponse ...
type ContactListResponse struct {
	Items []Contact `json:"items,omitempty"`
	Meta  Meta      `json:"meta,omitempty"`
}

// ContactResponse ...
type ContactResponse struct {
	Contact Contact `json:"contact,omitempty"`
}

// ContactCreateRequest ...
type ContactCreateRequest struct {
	Contact ContactRequest `json:"contact,omitempty"`
}

// ContactUpdateRequest ...
// TODO: "contact_people_attributes" alanı update etmiyor ekliyor, destek al
type ContactUpdateRequest struct {
	Contact ContactRequest `json:"contact,omitempty"`
}

// ContactRequest ...
type ContactRequest struct {
	Name                    string          `json:"name,omitempty"`
	ContactType             string          `json:"contact_type,omitempty"` // company, person
	Email                   string          `json:"emails,omitempty"`
	TaxNumber               string          `json:"tax_number,omitempty"`
	TaxOffice               string          `json:"tax_office,omitempty"`
	CategoryID              int             `json:"category_id,omitempty"`
	City                    string          `json:"city,omitempty"`
	District                string          `json:"district,omitempty"`
	AddressAttributes       Address         `json:"address_attributes,omitempty"`
	ContactPeopleAttributes []ContactPerson `json:"contact_people_attributes,omitempty"`
}

// ContactPerson ...
type ContactPerson struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
	Notes string `json:"notes,omitempty"`
}

// ContactDeleteResponse ...
type ContactDeleteResponse struct {
	Success string `json:"success,omitempty"`
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
	Name      string `json:"name,omitempty"`
	BgColor   string `json:"bg_color,omitempty"`
	TextColor string `json:"text_color,omitempty"`
}

// Address ...
type Address struct {
	Address string `json:"address,omitempty"`
	Fax     string `json:"fax,omitempty"`
	Phone   string `json:"phone,omitempty"`
}

// Meta ...
type Meta struct {
	TotalCount int `json:"total_count,omitempty"`
	PageCount  int `json:"page_count,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
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
