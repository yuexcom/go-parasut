package parasut

// ContactData ...
type ContactData struct {
	ID            string         `json:"id,omitempty"`
	Type          string         `json:"type,omitempty"` // contacts
	Attributes    *Contact       `json:"attributes,omitempty"`
	Relationships *Relationships `json:"relationships,omitempty"`
}

// Contact ...
type Contact struct {
	CreatedAt                 string   `json:"created_at,omitempty"`
	UpdatedAt                 string   `json:"updated_at,omitempty"`
	ContactType               string   `json:"contact_type,omitempty"` // "person", "company"
	Name                      string   `json:"name,omitempty"`         // required
	Email                     string   `json:"email,omitempty"`
	ShortName                 string   `json:"short_name,omitempty"`
	Balance                   string   `json:"balance,omitempty"`
	TRLBalance                string   `json:"trl_balance,omitempty"`
	USDBalance                string   `json:"usd_balance,omitempty"`
	EURBalance                string   `json:"eur_balance,omitempty"`
	GBPBalance                string   `json:"gbp_balance,omitempty"`
	TaxNumber                 string   `json:"tax_number,omitempty"`
	TaxOffice                 string   `json:"tax_office,omitempty"`
	Archived                  bool     `json:"archived,omitempty"`
	AccountType               string   `json:"account_type,omitempty"` // required, "customer", "supplier"
	City                      string   `json:"city,omitempty"`
	District                  string   `json:"district,omitempty"`
	Address                   string   `json:"address,omitempty"`
	Phone                     string   `json:"phone,omitempty"`
	Fax                       string   `json:"fax,omitempty"`
	IsAbroad                  bool     `json:"is_abroad,omitempty"`
	TermDays                  int      `json:"term_days,omitempty"`
	SharingsCount             int      `json:"sharings_count,omitempty"`
	IBANS                     []string `json:"ibans,omitempty"`
	ExchangeRateType          string   `json:"exchange_rate_type,omitempty"`
	IBAN                      string   `json:"iban,omitempty"`
	SharingPreviewURL         string   `json:"sharing_preview_url,omitempty"`
	PaymentReminderPreviewURL string   `json:"payment_reminder_preview_url,omitempty"`
}

// ContactRequest ... private kullanılacak complex create modeli
type ContactRequest struct {
	Data *ContactData `json:"data,omitempty"`
}

// ContactPersonData ...
type ContactPersonData struct {
	ID         string         `json:"id,omitempty"`
	Type       string         `json:"type,omitempty"`
	Attributes *ContactPerson `json:"attributes,omitempty"`
}

// ContactPerson ...
type ContactPerson struct {
	Name  string `json:"name,osmitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
	Notes string `json:"notes,omitempty"`
}

// ContactListOptions ...
type ContactListOptions struct {
	Filter  *ContactListFilter `url:"filter,omitempty"`
	Sort    string             `url:"sort,omitempty"`
	Page    *PageOptions       `url:"page,omitempty"`
	Include string             `url:"include,omitempty"`
}

// ContactGetOptions ...
type ContactGetOptions struct {
	Include string `url:"include,omitempty"`
}

// ContactListFilter ...
type ContactListFilter struct {
	Name      string `url:"name,omitempty"`
	Email     string `url:"email,omitempty"`
	TaxNumber string `url:"tax_number,omitempty"`
	TaxOffice string `url:"tax_office,omitempty"`
	City      string `url:"city,omitempty"`
}

// ContactResponse ...
type ContactResponse struct {
	Data     *ContactData `json:"data,omitempty"`
	Included []*Included  `json:"included,omitempty"`
}

// ContactsResponse ...
type ContactsResponse struct {
	Data     []*ContactData `json:"data,omitempty"`
	Included []*Included    `json:"included,omitempty"`
	Meta     Meta           `json:"meta,omitempty"`
}
