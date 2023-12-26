package product

type CreateProductRequestPayload struct {
	Name  string `json:"name"`
	Stock int16  `json:"stock"`
	Price int    `json:"price"`
}

type ListProductRequestPayload struct {
	Cursor int `query:"cursor" json:"cursor"`
	Size   int `query:"size" json:"size"`
}

func (l ListProductRequestPayload) GenerateDefaultValue() ListProductRequestPayload {
	if l.Cursor < 0 {
		l.Cursor = 0
	}
	if l.Size <= 0 {
		l.Size = 10
	}
	return l
}
