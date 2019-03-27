package parasut

type tagData struct {
	ID            string         `json:"id,omitempty"`
	Type          string         `json:"type,omitempty"`
	Attributes    *tag           `json:"attributes,omitempty"`
	Relationships *relationships `json:"relationships,omitempty"`
}

type tag struct {
	Name string `json:"name,omitempty"`
}
