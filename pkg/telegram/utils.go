package telegram

func isCategoryValid(category string, kind string) bool {
	expensesCategories := [5]string{
		"food",
		"transportation",
		"savings",
		"others",
		"subscribtions",
	}
	incomeCategories := [5]string{
		"job",
		"freelancing",
		"buisness",
		"cashback",
		"others",
	}
	var categories [5]string
	if kind == "expenses" {
		categories = expensesCategories
	} else {
		categories = incomeCategories
	}
	for _, item := range categories {
		if category == item {
			return true
		}
	}
	return false
}
