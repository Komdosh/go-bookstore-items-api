package queries

type EsQuery struct {
	Equals    []FieldValue `json:"equals"`
	NotEquals []FieldValue `json:"not_equals"`
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
