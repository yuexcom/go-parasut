package parasut

// --------------------------------
// ----- SALES INVOICE MODELS -----
// --------------------------------

// SalesInvoice ...
type SalesInvoice struct {
	Archived               bool    `json:"archived"`
	InvoiceNo              string  `json:"invoice_no"`
	NetTotal               float64 `json:"net_total"`
	GrossTotal             float64 `json:"gross_total"`
	Withholding            float64 `json:"withholding"`
	TotalExciseDuty        float64 `json:"total_excise_duty"`
	TotalCommunicationsTax float64 `json:"total_communications_tax"`
	TotalVat               float64 `json:"total_vat"`
	VatWithholding         float64 `json:"vat_withholding"`
	TotalDiscount          float64 `json:"total_discount"`
	TotalInvoiceDiscount   float64 `json:"total_invoice_discount"`
	BeforeTaxesTotal       float64 `json:"before_taxes_total"`
	Remaining              float64 `json:"remaining"`
	RemainingInTRL         float64 `json:"remaining_in_trl"`
	PaymentStatus          string  `json:"payment_status"`
	CreatedAt              string  `json:"created_at"`
	UpdatedAt              string  `json:"updated_at"`
	ItemType               string  `json:"item_type"`
	Description            string  `json:"description"`
	IssueDate              string  `json:"issue_date"`
	DueDate                string  `json:"due_date"`
	InvoiceSeries          string  `json:"invoice_series"`
	InvoiceID              int     `json:"invoice_id"`
	Currency               string  `json:"currency"`
	ExchangeRate           float64 `json:"exchange_rate"`
	WithholdingRate        float64 `json:"withholding_rate"`
	VatWithholdingRate     float64 `json:"vat_withholding_rate"`
	InvoiceDiscountType    string  `json:"invoice_discount_type"`
	InvoiceDiscount        float64 `json:"invoice_discount"`
	BillingAddress         string  `json:"billing_address"`
	BillingPhone           string  `json:"billing_phone"`
	BillingFax             string  `json:"billing_fax"`
	TaxOffice              string  `json:"tax_office"`
	TaxNumber              string  `json:"tax_number"`
	City                   string  `json:"city"`
	District               string  `json:"district"`
	IsAbroad               string  `json:"is_abroad"`
	OrderNo                string  `json:"order_no"`
	OrderDate              string  `json:"order_date"`
	ShipmentAddress        string  `json:"shipment_address"`
}

// SalesInvoiceIndexRequest ...
type SalesInvoiceIndexRequest struct {
	Filter  map[string]interface{} `json:"filter"`
	Sort    string                 `json:"sort"`
	Page    map[string]int         `json:"page"`
	Include string                 `json:"include"`
}

// SalesInvoiceCreateRequest ...
type SalesInvoiceCreateRequest struct {
	Data struct {
		ID            string       `json:"id"`
		Type          string       `json:"type"`
		Attributes    SalesInvoice `json:"attributes"`
		Relationships struct {
			Category Relationship  `json:"category"`
			Contact  Relationship  `json:"contact"`
			Details  Relationships `json:"details"`
			Tags     Relationships `json:"tags"`
		} `json:"relationships"`
	} `json:"data"`
}

// SalesInvoiceEditRequest ...
type SalesInvoiceEditRequest struct {
	Data struct {
		ID            string       `json:"id"`
		Type          string       `json:"type"`
		Attributes    SalesInvoice `json:"attributes"`
		Relationships struct {
			Category Relationship  `json:"category"`
			Contact  Relationship  `json:"contact"`
			Details  Relationships `json:"details"`
			Tags     Relationships `json:"tags"`
		} `json:"relationships"`
	} `json:"data"`
}

// SalesInvoiceShowRequest ...
type SalesInvoiceShowRequest struct {
	Include string `json:"include"`
}

// SalesInvoiceResponse ...
type SalesInvoiceResponse struct {
	Data struct {
		ID            string       `json:"id"`
		Type          string       `json:"type"`
		Attributes    SalesInvoice `json:"attributes"`
		Relationships struct {
			Category        Relationship  `json:"scategory"`
			Contact         Relationship  `json:"contact"`
			Details         Relationships `json:"details"`
			Payments        Relationships `json:"payments"`
			Tags            Relationships `json:"tags"`
			Sharings        Relationships `json:"sharings"`
			RecurrencePlan  Relationship  `json:"recurrence_plan"`
			ActiveEDocument Relationship  `json:"active_e_document"`
		} `json:"relationships"`
	} `json:"data"`
	Included []Included `json:"included"`
	Meta     Meta       `json:"meta"`
}

// --------------------------
// ----- CONTACT MODELS -----
// --------------------------

// Contact ...
type Contact struct {
	CreatedAt                 string   `json:"created_at"`
	UpdatedAt                 string   `json:"updated_at"`
	ContactType               string   `json:"contact_type"`
	Name                      string   `json:"name"`
	Email                     string   `json:"email"`
	ShortName                 string   `json:"short_name"`
	Balance                   string   `json:"balance"`
	TRLBalance                string   `json:"trl_balance"`
	USDBalance                string   `json:"usd_balance"`
	EURBalance                string   `json:"eur_balance"`
	GBPBalance                string   `json:"gbp_balance"`
	TaxNumber                 string   `json:"tax_number"`
	TaxOffice                 string   `json:"tax_office"`
	Archived                  bool     `json:"archived"`
	AccountType               string   `json:"account_type"`
	City                      string   `json:"city"`
	District                  string   `json:"district"`
	Address                   string   `json:"address"`
	Phone                     string   `json:"phone"`
	Fax                       string   `json:"fax"`
	IsAbroad                  bool     `json:"is_abroad"`
	TermDays                  string   `json:"term_days"`
	InvoicingPreferences      struct{} `json:"invoicing_preferences"`
	SharingsCount             int      `json:"sharings_count"`
	IBANS                     []string `json:"ibans"`
	ExchangeRateType          string   `json:"exchange_rate_type"`
	IBAN                      string   `json:"iban"`
	SharingPreviewURL         string   `json:"sharing_preview_url"`
	PaymentReminderPreviewURL string   `json:"payment_reminder_preview_url"`
}

// ContactIndexRequest ...
type ContactIndexRequest struct {
	Filter  map[string]string `json:"filter"`
	Sort    string            `json:"sort"`
	Page    map[string]int    `json:"page"`
	Include string            `json:"include"`
}

// ContactCreateRequest ...
type ContactCreateRequest struct {
	Data struct {
		ID            string  `json:"id"`
		Type          string  `json:"type"`
		Attributes    Contact `json:"attributes"`
		Relationships struct {
			Category Relationship  `json:"category"`
			Contact  Relationship  `json:"contact"`
			Details  Relationships `json:"details"`
			Tags     Relationships `json:"tags"`
		} `json:"relationships"`
	} `json:"data"`
}

// ContactEditRequest ...
type ContactEditRequest struct {
	Data struct {
		ID            string  `json:"id"`
		Type          string  `json:"type"`
		Attributes    Contact `json:"attributes"`
		Relationships struct {
			Category      Relationship  `json:"category"`
			ContactPeople Relationships `json:"contact_people"`
		} `json:"relationships"`
	} `json:"data"`
}

// ContactShowRequest ...
type ContactShowRequest struct {
	Data struct {
		ID            string  `json:"id"`
		Type          string  `json:"type"`
		Attributes    Contact `json:"attributes"`
		Relationships struct {
			Category      Relationship  `json:"category"`
			ContactPortal Relationship  `json:"contact_portal"`
			ContactPeople Relationships `json:"contact_people"`
		} `json:"relationships"`
	} `json:"data"`
	Included []Included `json:"included"`
}

// ContactResponse ...
type ContactResponse struct {
	Data []struct {
		ID            string  `json:"id"`
		Type          string  `json:"type"`
		Attributes    Contact `json:"attributes"`
		Relationships struct {
			Category        Relationship  `json:"category"`
			ContactPortal   Relationship  `json:"contact_portal"`
			ContactPeople   Relationships `json:"contact_people"`
			PriceList       Relationship  `json:"price_list"`
			Activities      Relationships `json:"activities"`
			EInvoiceInboxes Relationships `json:"e_invoice_inboxes"`
			Sharings        Relationships `json:"sharings"`
		} `json:"relationships"`
	} `json:"data"`
	Meta Meta `json:"meta"`
}

// -------------------------
// ----- COMMON MODELS -----
// -------------------------

// Included ...
type Included struct {
	ID            string   `json:"id"`
	Type          string   `json:"type"`
	Attributes    struct{} `json:"attributes"`
	Relationships struct{} `json:"relationships"`
}

// Relationship ...
type Relationship struct {
	Data struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// Relationships ...
type Relationships struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// Meta ...
type Meta struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	TotalCount  int `json:"total_count"`
}
