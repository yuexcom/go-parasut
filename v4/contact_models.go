package parasut

// contactData ...
type contactData struct {
	ID            string         `json:"id,omitempty"`
	Type          string         `json:"type,omitempty"` // contacts
	Attributes    *contact       `json:"attributes,omitempty"`
	Relationships *relationships `json:"relationships,omitempty"`
}

// contact ...
type contact struct {
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

// contactRequest ... private kullanılacak complex create modeli
type contactRequest struct {
	Data *contactData `json:"data,omitempty"`
}

// contactListParams ... private kullanılacak complex listing modeli
type contactListParams struct {
	Filter  *ContactListFilter `url:"filter,omitempty"`
	Sort    string             `url:"sort,omitempty"` // "id", "balance", "name", "email"
	Page    *PageOptions       `url:"page,omitempty"`
	Include string             `url:"include,omitempty"` // "category", "contact_portal", "contact_people"
}

// contactPersonData ...
type contactPersonData struct {
	ID         string         `json:"id,omitempty"`
	Type       string         `json:"type,omitempty"`
	Attributes *contactPerson `json:"attributes,omitempty"`
}

// contactPerson ...
type contactPerson struct {
	Name  string `json:"name,osmitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
	Notes string `json:"notes,omitempty"`
}

// ContactListFilter ... export edilen, dışarıdan istek yaparken kullanılacak listing modeli
type ContactListFilter struct {
	Name      string `url:"name,omitempty"`
	Email     string `url:"email,omitempty"`
	TaxNumber string `url:"tax_number,omitempty"`
	TaxOffice string `url:"tax_office,omitempty"`
	City      string `url:"city,omitempty"`
}

// ContactListOptions ... export edilen, dışarıdan istek yaparken kullanılacak listing modeli
type ContactListOptions struct {
	Sort       string `json:"sort,omitempty"`
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	Include    string `json:"include,omitempty"`
}

// Contact ...
type Contact struct {
	Name          string           `json:"name,omitempty"`         // required
	AccountType   string           `json:"account_type,omitempty"` // required, "customer", "supplier"
	ContactType   string           `json:"contact_type,omitempty"`
	Email         string           `json:"email,omitempty"`
	ShortName     string           `json:"short_name,omitempty"`
	TaxNumber     string           `json:"tax_number,omitempty"`
	TaxOffice     string           `json:"tax_office,omitempty"`
	City          string           `json:"city,omitempty"`
	District      string           `json:"district,omitempty"`
	Address       string           `json:"address,omitempty"`
	Phone         string           `json:"phone,omitempty"`
	CategoryID    string           `json:"category_id,omitempty"`
	ContactPeople []*ContactPerson `json:"contact_people,omitempty"`
}

// ContactPerson ...
type ContactPerson struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
	Notes string `json:"notes,omitempty"`
}

// ContactResponse ...
type ContactResponse struct {
	Data     *contactData `json:"data,omitempty"`
	Included []*included  `json:"included,omitempty"`
}

// ContactsResponse ...
type ContactsResponse struct {
	Data     []*contactData `json:"data,omitempty"`
	Included []*included    `json:"included,omitempty"`
	Meta     Meta           `json:"meta,omitempty"`
}
