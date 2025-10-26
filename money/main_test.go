package money

import "testing"

func TestCreateDollarObject(t *testing.T) {
	five := NewDollar(5)
	if five.Amount() != 5 {
		t.Errorf("Expected 5, got %v", five.Amount())
	}
}
