package parasut

// SalesInvoice ...
type SalesInvoice struct {
	ID                  string                 `json:"sales_invoice_id,omitempty"`
	ItemType            string                 `json:"item_type,omitempty"`  // required, "invoice" "estimate" "cancelled" "recurring_invoice" "recurring_estimate" "refund"
	IssueDate           string                 `json:"issue_date,omitempty"` // required
	DueDate             string                 `json:"due_date,omitempty"`
	Description         string                 `json:"description,omitempty"`
	SalesInvoiceDetails []*SalesInvoiceDetails `json:"sales_invoice_details,omitempty"`
	ContactID           string                 `json:"contact_id,omitempty"`
	CategoryID          string                 `json:"category_id,omitempty"`
	TagsID              []string               `json:"tags_id,omitempty"`
}

// SalesInvoiceDetails ...
type SalesInvoiceDetails struct {
	ID        string  `json:"id,omitempty"`
	Quantity  float64 `json:"quantity,omitempty"`   // required
	UnitPrice float64 `json:"unit_price,omitempty"` // required
	VatRate   float64 `json:"vat_rate,omitempty"`   // required
	ProductID string  `json:"product_id,omitempty"`
}

// SalesInvoiceListFilter ... export edilen, dışarıdan istek yaparken kullanılacak listing modeli
type SalesInvoiceListFilter struct {
	IssueDate     string `url:"issue_date,omitempty"`
	DueDate       string `url:"due_date,omitempty"`
	ContactID     int    `url:"contact_id,omitempty"`
	InvoiceID     int    `url:"invoice_id,omitempty"`
	InvoiceSeries string `url:"invoice_series,omitempty"`
}

// SalesInvoiceListOptions ... export edilen, dışarıdan istek yaparken kullanılacak listing modeli
type SalesInvoiceListOptions struct {
	Sort       string `json:"sort,omitempty"`
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	Include    string `json:"include,omitempty"`
}

// SalesInvoiceResponse ...
type SalesInvoiceResponse struct {
	Data     *salesInvoiceData `json:"data"`
	Included []*included       `json:"included"`
}

// SalesInvoicesResponse ...
type SalesInvoicesResponse struct {
	Data     []*salesInvoiceData `json:"data"`
	Included []*included         `json:"included"`
	Meta     Meta                `json:"meta"`
}

// salesInvoiceData ...
type salesInvoiceData struct {
	ID            string         `json:"id"`
	Type          string         `json:"type"`
	Attributes    *salesInvoice  `json:"attributes"`
	Relationships *relationships `json:"relationships"`
}

// salesInvoice ...
type salesInvoice struct {
	Archived               bool   `json:"archived,omitempty"`
	InvoiceNo              string `json:"invoice_no,omitempty"`
	NetTotal               string `json:"net_total,omitempty"`
	GrossTotal             string `json:"gross_total,omitempty"`
	Withholding            string `json:"withholding,omitempty"`
	TotalExciseDuty        string `json:"total_excise_duty,omitempty"`
	TotalCommunicationsTax string `json:"total_communications_tax,omitempty"`
	TotalVat               string `json:"total_vat,omitempty"`
	VatWithholding         string `json:"vat_withholding,omitempty"`
	TotalDiscount          string `json:"total_discount,omitempty"`
	TotalInvoiceDiscount   string `json:"total_invoice_discount,omitempty"`
	BeforeTaxesTotal       string `json:"before_taxes_total,omitempty"`
	Remaining              string `json:"remaining,omitempty"`
	RemainingInTRL         string `json:"remaining_in_trl,omitempty"`
	PaymentStatus          string `json:"payment_status,omitempty"` // "paid" "overdue" "unpaid" "partially_paid"
	CreatedAt              string `json:"created_at,omitempty"`
	UpdatedAt              string `json:"updated_at,omitempty"`
	ItemType               string `json:"item_type,omitempty"` // required, "invoice" "estimate" "cancelled" "recurring_invoice" "recurring_estimate" "refund"
	Description            string `json:"description,omitempty"`
	IssueDate              string `json:"issue_date,omitempty"` // required
	DueDate                string `json:"due_date,omitempty"`
	InvoiceSeries          string `json:"invoice_series,omitempty"`
	InvoiceID              int    `json:"invoice_id,omitempty"`
	Currency               string `json:"currency,omitempty"` // "TRL" "USD" "EUR" "GBP"
	ExchangeRate           string `json:"exchange_rate,omitempty"`
	WithholdingRate        string `json:"withholding_rate,omitempty"`
	VatWithholdingRate     string `json:"vat_withholding_rate,omitempty"`
	InvoiceDiscountType    string `json:"invoice_discount_type,omitempty"` // "percentage" "amount"
	InvoiceDiscount        string `json:"invoice_discount,omitempty"`
	BillingAddress         string `json:"billing_address,omitempty"`
	BillingPhone           string `json:"billing_phone,omitempty"`
	BillingFax             string `json:"billing_fax,omitempty"`
	TaxOffice              string `json:"tax_office,omitempty"`
	TaxNumber              string `json:"tax_number,omitempty"`
	City                   string `json:"city,omitempty"`
	District               string `json:"district,omitempty"`
	IsAbroad               bool   `json:"is_abroad,omitempty"`
	OrderNo                string `json:"order_no,omitempty"`
	OrderDate              string `json:"order_date,omitempty"`
	ShipmentAddress        string `json:"shipment_address,omitempty"`
}

// salesInvoiceRequest ... private kullanılacak complex create modeli
type salesInvoiceRequest struct {
	Data *salesInvoiceData `json:"data"`
}

// salesInvoiceListParams ... private kullanılacak complex listing modeli
type salesInvoiceListParams struct {
	Filter  *SalesInvoiceListFilter `url:"filter,omitempty"`
	Sort    string                  `url:"sort,omitempty"` // id, issue_date, due_date, remaining, remaining_in_trl, description, net_total
	Page    *PageOptions            `url:"page,omitempty"`
	Include string                  `url:"include,omitempty"` // category, contact, details, details.product, payments, payments.transaction, tags, sharings, recurrence_plan, active_e_document
}

// salesInvoiceDetailsData ...
type salesInvoiceDetailsData struct {
	ID            string               `json:"id"`
	Type          string               `json:"type"`
	Attributes    *salesInvoiceDetails `json:"attributes"`
	Relationships *relationships       `json:"relationships"`
}

// salesInvoiceDetails ...
type salesInvoiceDetails struct {
	Quantity             float64 `json:"quantity,omitempty"`      // required
	UnitPrice            float64 `json:"unit_price,omitempty"`    // required
	VatRate              float64 `json:"vat_rate,omitempty"`      // required
	DiscountType         string  `json:"discount_type,omitempty"` // "percentage", "amount"
	DiscountValue        float64 `json:"discount_value,omitempty"`
	ExciseDutyType       string  `json:"excise_duty_type,omitempty"` // "percentage", "amount"
	ExciseDutyValue      float64 `json:"excise_duty_value,omitempty"`
	CommunicationTaxRate float64 `json:"communication_tax_rate,omitempty"`
	Description          string  `json:"description,omitempty"`
}
