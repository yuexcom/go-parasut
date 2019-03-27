package parasut

type categoryData struct {
	ID            string         `json:"id,omitempty"`
	Type          string         `json:"type,omitempty"`
	Attributes    *category      `json:"attributes,omitempty"`
	Relationships *relationships `json:"relationships,omitempty"`
}

type category struct {
	Name         string `json:"name,omitempty"`
	CategoryType string `json:"category_type,omitempty"` // "Product" "Contact" "Employee" "SalesInvoice" "Expenditure"
	BgColor      string `json:"bg_color,omitempty"`
	TextColor    string `json:"text_color,omitempty"`
}
