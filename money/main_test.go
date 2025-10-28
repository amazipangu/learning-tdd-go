package money

import "testing"

func TestCreateDollarObject(t *testing.T) {
	five := NewDollar(5)
	if five.Amount() != 5 {
		t.Errorf("Expected 5, got %v", five.Amount())
	}
}

func TestMultiplyDollar(t *testing.T) {
	five := NewDollar(5)
	result := five.Times(2)
	if result.Amount() != 10 {
		t.Errorf("Expected 10, got %v", result.Amount())
	}
	product := five.Times(3)
	if product.Amount() != 15 {
		t.Errorf("Expected 15, got %v", product.Amount())
	}
}
