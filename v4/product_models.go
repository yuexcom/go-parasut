package parasut

type productData struct {
	ID            string         `json:"id,omitempty"`
	Type          string         `json:"type,omitempty"`
	Attributes    *product       `json:"attributes,omitempty"`
	Relationships *relationships `json:"relationships,omitempty"`
}

type product struct {
	Name string `json:"name"`
	Code string `json:"code"`
	// ...
}
