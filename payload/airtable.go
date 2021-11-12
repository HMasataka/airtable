package payload

type Tables struct {
	Tables []Table `json:"tables"`
}

type Table struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	PrimaryFieldId string  `json:"primaryFieldId"`
	Fields         []Field `json:"fields"`
	Views          []View  `json:"views"`
}

type Field struct {
	Id      string      `json:"id"`
	Name    string      `json:"name"`
	Type    string      `json:"type"`
	Options interface{} `json:"options,omitempty"`
}

type View struct {
	Id      string      `json:"id"`
	Name    string      `json:"name"`
	Type    string      `json:"type"`
	Options interface{} `json:"options,omitempty"`
}
