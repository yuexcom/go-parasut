package parasut

// SalesInvoiceData ...
type SalesInvoiceData struct {
	ID            string         `json:"id"`
	Type          string         `json:"type"`
	Attributes    *SalesInvoice  `json:"attributes"`
	Relationships *Relationships `json:"relationships"`
}

// SalesInvoice ...
type SalesInvoice struct {
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

// SalesInvoiceRequest ...
type SalesInvoiceRequest struct {
	Data *SalesInvoiceData `json:"data"`
}

// SalesInvoiceListOptions ...
type SalesInvoiceListOptions struct {
	Filter  *SalesInvoiceListFilter `url:"filter,omitempty"`
	Sort    string                  `url:"sort,omitempty"` // id, issue_date, due_date, remaining, remaining_in_trl, description, net_total
	Page    *PageOptions            `url:"page,omitempty"`
	Include string                  `url:"include,omitempty"` // category, contact, details, details.product, payments, payments.transaction, tags, sharings, recurrence_plan, active_e_document
}

// SalesInvoiceGetOptions ...
type SalesInvoiceGetOptions struct {
	Include string `url:"include,omitempty"` // category, contact, details, details.product, payments, payments.transaction, tags, sharings, recurrence_plan, active_e_document
}

// SalesInvoiceListFilter ...
type SalesInvoiceListFilter struct {
	IssueDate     string `url:"issue_date,omitempty"`
	DueDate       string `url:"due_date,omitempty"`
	ContactID     int    `url:"contact_id,omitempty"`
	InvoiceID     int    `url:"invoice_id,omitempty"`
	InvoiceSeries string `url:"invoice_series,omitempty"`
}

// SalesInvoiceDetailsData ...
type SalesInvoiceDetailsData struct {
	ID            string               `json:"id"`
	Type          string               `json:"type"`
	Attributes    *SalesInvoiceDetails `json:"attributes"`
	Relationships *Relationships       `json:"relationships"`
}

// SalesInvoiceDetails ...
type SalesInvoiceDetails struct {
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

// SalesInvoiceResponse ...
type SalesInvoiceResponse struct {
	Data     *SalesInvoiceData `json:"data"`
	Included []*Included       `json:"included"`
}

// SalesInvoicesResponse ...
type SalesInvoicesResponse struct {
	Data     []*SalesInvoiceData `json:"data"`
	Included []*Included         `json:"included"`
	Meta     Meta                `json:"meta"`
}
