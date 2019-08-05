package app

// Food is the model for some dishes
type Food struct {
	Title   string
	Price   int
	SoldOut bool
}

// GetFoods returns slice of dishes
func GetFoods() []Food {
	return []Food{
		{"Доширак по-гречески", 999, true},
		{"Курица в полете", 1342, false},
		{"Огурец среднепупырчатый", 634, true},
	}
}
