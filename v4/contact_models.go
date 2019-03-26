package parasut

// ContactData ...
type ContactData struct {
	ID            string         `json:"id"`
	Type          string         `json:"type"`
	Attributes    *Contact       `json:"attributes"`
	Relationships *Relationships `json:"relationships"`
}

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
	TermDays                  int      `json:"term_days"`
	SharingsCount             int      `json:"sharings_count"`
	IBANS                     []string `json:"ibans"`
	ExchangeRateType          string   `json:"exchange_rate_type"`
	IBAN                      string   `json:"iban"`
	SharingPreviewURL         string   `json:"sharing_preview_url"`
	PaymentReminderPreviewURL string   `json:"payment_reminder_preview_url"`
}

// Relationships ...
type Relationships map[string]*Relationship

// Relationship interface{} çünkü, struct veya slice olabiliyor
type Relationship struct {
	Data interface{} `json:"data"`
}

// RelationshipData 1:1 ilişki tipi için
type RelationshipData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// RelationshipDataList 1:m ilişki tipi için
type RelationshipDataList []*RelationshipData

// contactCreate ... private kullanılacak complex create modeli
type contactCreate struct {
	Data          *ContactData   `json:"data"`
	Relationships *Relationships `json:"relationships"`
}

// contactUpdate ... private kullanılacak complex update modeli
type contactUpdate struct {
	Data *ContactData `json:"data"`
}

// contactListParams ... private kullanılacak complex listing modeli
type contactListParams struct {
	Filter  *ContactListFilter `url:"filter,omitempty"`
	Sort    string             `url:"sort,omitempty"`
	Page    *PageOptions       `url:"page,omitempty"`
	Include string             `url:"include,omitempty"`
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

// ContactResponse ...
type ContactResponse struct {
	Data     *ContactData `json:"data"`
	Included []*Included  `json:"included"`
}

// ContactsResponse ...
type ContactsResponse struct {
	Data     []*ContactData `json:"data"`
	Included []*Included    `json:"included"`
	Meta     Meta           `json:"meta"`
}

// Contact Person

// ContactPersonData ...
type ContactPersonData struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes *ContactPerson `json:"attributes"`
}

// ContactPerson ...
type ContactPerson struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Notes string `json:"notes"`
}

// // Category

// // Category ...
// type Category struct {
// 	FullPath     string `json:"full_path"`
// 	CreatedAt    string `json:"created_at"`
// 	UpdatedAt    string `json:"updated_at"`
// 	Name         string `json:"name"`
// 	BgColor      string `json:"bg_color"`
// 	TextColor    string `json:"text_color"`
// 	CategoryType string `json:"category_type"`
// 	ParentID     int    `json:"parent_id"`
// }

// // CategoryData ...
// type CategoryData struct {
// 	ID            string        `json:"id"`
// 	Type          string        `json:"type"`
// 	Attributes    Category      `json:"attributes"`
// 	Relationships Relationships `json:"relationships"`
// }

// // ParentCategoryData ...
// type ParentCategoryData struct {
// 	ID   string `json:"id"`
// 	Type string `json:"type"`
// }

// // SubcategoryData ...
// type SubcategoryData struct {
// 	ID   string `json:"id"`
// 	Type string `json:"type"`
// }
