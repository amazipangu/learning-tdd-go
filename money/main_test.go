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

func TestEquality(t *testing.T) {
	a := NewDollar(5)
	isTrue := NewDollar(5).Equal(a)
	isFalse := NewDollar(6).Equal(a)
	if isTrue != true {
		t.Errorf("Expected true, got %v", isTrue)
	}
	if isFalse != false {
		t.Errorf("Expected false, got %v", isFalse)
	}

}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	result := five.Times(2)
	if result.Amount() != 10 {
		t.Errorf("Expected 10, got %v", result.Amount())
	}
}
