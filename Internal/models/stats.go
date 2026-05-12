package models

type Summary struct {
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
	Balance float64 `json:"balance"`
}

type TopCategory struct {
	CategoryName string  `json:"category_name"`
	Total        float64 `json:"total"`
}
