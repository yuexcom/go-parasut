package parasut

// Contact ...
type Contact struct {
	CreatedAt                 string   `json:"created_at,omitempty"`
	UpdatedAt                 string   `json:"updated_at,omitempty"`
	ContactType               string   `json:"contact_type,omitempty"`
	Name                      string   `json:"name,omitempty"`
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
	AccountType               string   `json:"account_type,omitempty"`
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
	// InvoicingPreferences      *struct{} `json:"invoicing_preferences"`
}

// ContactPeople ...
type ContactPeople struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
	Notes string `json:"notes,omitempty"`
}

// Category ...
type Category struct {
	FullPath     string `json:"full_path,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	Name         string `json:"name,omitempty"`
	BgColor      string `json:"bg_color,omitempty"`
	TextColor    string `json:"text_color,omitempty"`
	CategoryType string `json:"category_type,omitempty"`
	ParentID     int    `json:"parent_id,omitempty"`
}

// CategoryData ...
type CategoryData struct {
	ID            string   `json:"id,omitempty"`
	Type          string   `json:"type,omitempty"`
	Attributes    Category `json:"attributes,omitempty"`
	Relationships *struct {
		ParentCategory ParentCategoryRelationship `json:"parent_category,omitempty"`
		Subcategories  SubcategoryRelationship    `json:"subcategories,omitempty"`
	} `json:"relationships,omitempty"`
}

// ParentCategoryData ...
type ParentCategoryData struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

// SubcategoryData ...
type SubcategoryData struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

// ContactData ...
type ContactData struct {
	ID            string  `json:"id,omitempty"`
	Type          string  `json:"type,omitempty"`
	Attributes    Contact `json:"attributes,omitempty"`
	Relationships *struct {
		Category      *CategoryRelationship      `json:"category,omitempty"`
		ContactPeople *ContactPeopleRelationship `json:"contact_people,omitempty"`
		ContactPortal *ContactPortalRelationship `json:"contact_portal,omitempty"`
	} `json:"relationships,omitempty"`
}

// ContactPeopleData ...
type ContactPeopleData struct {
	ID         string        `json:"id,omitempty"`
	Type       string        `json:"type,omitempty"`
	Attributes ContactPeople `json:"attributes,omitempty"`
}

// ContactPortalData ...
type ContactPortalData struct{}

// CategoryRelationship ...
type CategoryRelationship struct {
	Data CategoryData `json:"data,omitempty"`
}

// ContactPortalRelationship ...
type ContactPortalRelationship struct {
	Data ContactPortalData `json:"data,omitempty"`
}

// ContactPeopleRelationship ...
type ContactPeopleRelationship struct {
	Data []ContactPeopleData `json:"data,omitempty"`
}

// ParentCategoryRelationship ...
type ParentCategoryRelationship struct {
	Data ParentCategoryData `json:"data,omitempty"`
}

// SubcategoryRelationship ...
type SubcategoryRelationship struct {
	Data SubcategoryData `json:"data,omitempty"`
}

// ContactListOptions ...
type ContactListOptions struct {
	Filter  ContactFilterOptions `url:"filter,omitempty"`
	Sort    string               `url:"sort,omitempty"`
	Page    PageOptions          `url:"page,omitempty"`
	Include string               `url:"include,omitempty"`
}

// ContactFilterOptions ...
type ContactFilterOptions struct {
	Name      string `url:"name,omitempty"`
	Email     string `url:"email,omitempty"`
	TaxNumber string `url:"tax_number,omitempty"`
	TaxOffice string `url:"tax_office,omitempty"`
	City      string `url:"city,omitempty"`
}

// ContactCreateRequest ...
type ContactCreateRequest struct {
	Data ContactData `json:"data,omitempty"`
}

// ContactUpdateRequest ...
type ContactUpdateRequest struct {
	Data ContactData `json:"data,omitempty"`
}

// ContactShowRequest ...
type ContactShowRequest struct {
	Data     ContactData `json:"data,omitempty"`
	Included []Included  `json:"included,omitempty"`
}

// ContactResponse ...
type ContactResponse struct {
	Data     ContactData `json:"data,omitempty"`
	Included []Included  `json:"included,omitempty"`
	Meta     Meta        `json:"meta,omitempty"`
}

// ContactsResponse ...
type ContactsResponse struct {
	Data     []ContactData `json:"data,omitempty"`
	Included []Included    `json:"included,omitempty"`
	Links    struct{}      `json:"links,omitempty"`
	Meta     Meta          `json:"meta,omitempty"`
}
