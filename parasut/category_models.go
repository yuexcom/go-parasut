package parasut

// CategoryData ...
type CategoryData struct {
	ID            string         `json:"id,omitempty"`
	Type          string         `json:"type,omitempty"`
	Attributes    *Category      `json:"attributes,omitempty"`
	Relationships *Relationships `json:"relationships,omitempty"`
}

// Category ...
type Category struct {
	Name         string `json:"name,omitempty"`
	CategoryType string `json:"category_type,omitempty"` // "Product" "Contact" "Employee" "SalesInvoice" "Expenditure"
	BgColor      string `json:"bg_color,omitempty"`
	TextColor    string `json:"text_color,omitempty"`
}
