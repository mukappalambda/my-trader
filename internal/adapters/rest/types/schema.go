package types

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Doc  string `json:"doc"`
}

type Schema struct {
	Id      string  `json:"id,omitempty"`
	Subject string  `json:"subject"`
	Name    string  `json:"name"`
	Doc     string  `json:"doc"`
	Fields  []Field `json:"fields"`
}

var DefaultSchema = &Schema{
	Subject: "",
	Name:    "",
	Doc:     "",
	Fields:  []Field{{Name: "", Type: "", Doc: ""}},
}

func (s *Schema) SetId(id string) {
	s.Id = id
}
