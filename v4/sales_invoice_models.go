package parasut

// SalesInvoice ...
type SalesInvoice struct {
	Archived               bool    `json:"archived,omitempty"`
	InvoiceNo              string  `json:"invoice_no,omitempty"`
	NetTotal               float64 `json:"net_total,omitempty"`
	GrossTotal             float64 `json:"gross_total,omitempty"`
	Withholding            float64 `json:"withholding,omitempty"`
	TotalExciseDuty        float64 `json:"total_excise_duty,omitempty"`
	TotalCommunicationsTax float64 `json:"total_communications_tax,omitempty"`
	TotalVat               float64 `json:"total_vat,omitempty"`
	VatWithholding         float64 `json:"vat_withholding,omitempty"`
	TotalDiscount          float64 `json:"total_discount,omitempty"`
	TotalInvoiceDiscount   float64 `json:"total_invoice_discount,omitempty"`
	BeforeTaxesTotal       float64 `json:"before_taxes_total,omitempty"`
	Remaining              float64 `json:"remaining,omitempty"`
	RemainingInTRL         float64 `json:"remaining_in_trl,omitempty"`
	PaymentStatus          string  `json:"payment_status,omitempty"`
	CreatedAt              string  `json:"created_at,omitempty"`
	UpdatedAt              string  `json:"updated_at,omitempty"`
	ItemType               string  `json:"item_type,omitempty"`
	Description            string  `json:"description,omitempty"`
	IssueDate              string  `json:"issue_date,omitempty"`
	DueDate                string  `json:"due_date,omitempty"`
	InvoiceSeries          string  `json:"invoice_series,omitempty"`
	InvoiceID              int     `json:"invoice_id,omitempty"`
	Currency               string  `json:"currency,omitempty"`
	ExchangeRate           float64 `json:"exchange_rate,omitempty"`
	WithholdingRate        float64 `json:"withholding_rate,omitempty"`
	VatWithholdingRate     float64 `json:"vat_withholding_rate,omitempty"`
	InvoiceDiscountType    string  `json:"invoice_discount_type,omitempty"`
	InvoiceDiscount        float64 `json:"invoice_discount,omitempty"`
	BillingAddress         string  `json:"billing_address,omitempty"`
	BillingPhone           string  `json:"billing_phone,omitempty"`
	BillingFax             string  `json:"billing_fax,omitempty"`
	TaxOffice              string  `json:"tax_office,omitempty"`
	TaxNumber              string  `json:"tax_number,omitempty"`
	City                   string  `json:"city,omitempty"`
	District               string  `json:"district,omitempty"`
	IsAbroad               string  `json:"is_abroad,omitempty"`
	OrderNo                string  `json:"order_no,omitempty"`
	OrderDate              string  `json:"order_date,omitempty"`
	ShipmentAddress        string  `json:"shipment_address,omitempty"`
}

// SalesInvoiceListOptions ...
type SalesInvoiceListOptions struct {
	Filter  SalesInvoiceFilterOptions `json:"filter,omitempty"`
	Sort    string                    `json:"sort,omitempty"`
	Page    PageOptions               `json:"page,omitempty"`
	Include string                    `json:"include,omitempty"`
}

// SalesInvoiceFilterOptions ...
type SalesInvoiceFilterOptions struct {
	IssueDate     string `json:"issue_date,omitempty"`
	DueDate       string `json:"due_date,omitempty"`
	ContactID     int    `json:"contact_id,omitempty"`
	InvoiceID     int    `json:"invoice_id,omitempty"`
	InvoiceSeries string `json:"invoice_series,omitempty"`
}

// SalesInvoiceCreateRequest ...
type SalesInvoiceCreateRequest struct {
	Data SalesInvoiceData `json:"data,omitempty"`
}

// SalesInvoiceData ...
type SalesInvoiceData struct {
	ID            string       `json:"id,omitempty"`
	Type          string       `json:"type,omitempty"`
	Attributes    SalesInvoice `json:"attributes,omitempty"`
	Relationships *struct {
		Category        CategoryRelationship        `json:"category,omitempty"`
		Contact         ContactRelationship         `json:"contact,omitempty"`
		Details         DetailsRelationship         `json:"details,omitempty"`
		Tags            TagsRelationship            `json:"tags,omitempty"`
		Payments        PaymentsRelationship        `json:"payments,omitempty"`
		Sharings        SharingsRelationship        `json:"sharings,omitempty"`
		RecurrencePlan  RecurrePlanRelationship     `json:"recurrence_plan,omitempty"`
		ActiveEDocument ActiveEDocumentRelationship `json:"active_e_document,omitempty"`
	} `json:"relationships,omitempty"`
}

// ContactRelationship ...
type ContactRelationship struct {
	Data ContactData `json:"data,omitempty"`
}

// DetailsRelationship ...
type DetailsRelationship struct {
	Data []DetailsData `json:"data,omitempty"`
}

// DetailsData ...
type DetailsData struct {
	ID            string  `json:"id,omitempty"`
	Type          string  `json:"type,omitempty"`
	Attributes    Details `json:"attributes,omitempty"`
	Relationships *struct {
		Product ProductRelationship `json:"product,omitempty"`
	} `json:"relationships,omitempty"`
}

// Details ...
type Details struct {
	Quantity             float64 `json:"quantity,omitempty"`
	UnitPrice            float64 `json:"unit_price,omitempty"`
	VatRate              float64 `json:"vat_rate,omitempty"`
	DiscountType         string  `json:"discount_type,omitempty"`
	DiscountValue        string  `json:"discount_value,omitempty"`
	ExciseDutyType       string  `json:"excise_duty_type,omitempty"`
	ExciseDutyValue      float64 `json:"excise_duty_value,omitempty"`
	CommunicationTaxRate float64 `json:"communication_tax_rate,omitempty"`
	Description          string  `json:"description,omitempty"`
}

// PaymentsRelationship ...
type PaymentsRelationship struct {
	Data []PaymentData `json:"data,omitempty"`
}

// PaymentData ...
type PaymentData struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

// TagsRelationship ...
type TagsRelationship struct {
	Data []TagData `json:"data,omitempty"`
}

// TagData ...
type TagData struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

// SharingsRelationship ...
type SharingsRelationship struct {
	Data []SharingData `json:"data,omitempty"`
}

// SharingData ...
type SharingData struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

// RecurrePlanRelationship ...
type RecurrePlanRelationship struct {
	Data RecurrePlanData `json:"data,omitempty"`
}

// RecurrePlanData ...
type RecurrePlanData struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

// ActiveEDocumentRelationship ...
type ActiveEDocumentRelationship struct {
	Data ActiveEDocumentData `json:"data,omitempty"`
}

// ActiveEDocumentData ...
type ActiveEDocumentData struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

// ProductRelationship ...
type ProductRelationship struct {
	Data ProductData `json:"data,omitempty"`
}

// ProductData ...
type ProductData struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

// SalesInvoiceUpdateRequest ...
type SalesInvoiceUpdateRequest struct {
	Data SalesInvoiceData `json:"data,omitempty"`
}

// SalesInvoiceShowRequest ...
type SalesInvoiceShowRequest struct {
	Include string `json:"include,omitempty"`
}

// SalesInvoicesResponse ...
type SalesInvoicesResponse struct {
	Data     []SalesInvoiceData `json:"data,omitempty"`
	Included []Included         `json:"included,omitempty"`
	Meta     Meta               `json:"meta,omitempty"`
}

// SalesInvoiceResponse ...
type SalesInvoiceResponse struct {
	Data     SalesInvoiceData `json:"data,omitempty"`
	Included []Included       `json:"included,omitempty"`
	Meta     Meta             `json:"meta,omitempty"`
}
