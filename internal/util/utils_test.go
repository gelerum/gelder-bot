package util

import "testing"

func TestCategoryIsValid(t *testing.T) {
	if CategoryIsValid("food", "expenses") != true {
		t.Error(`CategoryIsValid("food", "expenses") = false`)
	}
	if CategoryIsValid("job", "income") != true {
		t.Error(`CategoryIsValid("job", "income") = false`)
	}
	if CategoryIsValid("food", "income") != false {
		t.Error(`CategoryIsValid("food", "income") = true`)
	}
	if CategoryIsValid("job", "expenses") != false {
		t.Error(`CategoryIsValid("job", "expenses") = true`)
	}
}
