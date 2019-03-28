package parasut

// TagData ...
type TagData struct {
	ID            string         `json:"id,omitempty"`
	Type          string         `json:"type,omitempty"`
	Attributes    *Tag           `json:"attributes,omitempty"`
	Relationships *Relationships `json:"relationships,omitempty"`
}

// Tag ...
type Tag struct {
	Name string `json:"name,omitempty"`
}
