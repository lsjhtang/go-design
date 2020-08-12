package models

const (
	ORDERBY_SCORE_ASC  = 1
	ORDERBY_SCORE_DESC = 2
)

type OrderBy struct {
	Score      bool `json:"score"`
	PriceOrder int  `json:"price_order" binding:"oneof=0 1 2"`
}

type Search struct {
	BookName   string  `json:"book_name"`
	BookPress  string  `json:"book_press"`
	BookPrice1 float64 `json:"book_price1"`
	BookPrice2 float64 `json:"book_price2"`
	OrderBy    `json:"OrderBy" binding:"required,dive"`
}

func NewSearch() *Search {
	return &Search{}
}
