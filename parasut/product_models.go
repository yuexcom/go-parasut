package parasut

// ProductData ...
type ProductData struct {
	ID            string         `json:"id,omitempty"`
	Type          string         `json:"type,omitempty"`
	Attributes    *Product       `json:"attributes,omitempty"`
	Relationships *Relationships `json:"relationships,omitempty"`
}

// Product ...
type Product struct {
	Name string `json:"name"`
	Code string `json:"code"`
	// ...
}
