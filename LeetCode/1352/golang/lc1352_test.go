package lc1352

import "testing"

func Test_ProductOfNumbers(t *testing.T) {
	p := Constructor()
	var actualProduct int
	p.Add(3)
	p.Add(0)
	p.Add(2)
	p.Add(5)
	p.Add(4)
	actualProduct = p.GetProduct(2)
	if actualProduct != 20 {
		t.Errorf("Check #1 failed: want 20, got %d", actualProduct)
	}
	actualProduct = p.GetProduct(3)
	if actualProduct != 40 {
		t.Errorf("Check #2 failed: want 40, got %d", actualProduct)
	}
	actualProduct = p.GetProduct(4)
	if actualProduct != 0 {
		t.Errorf("Check #3 failed: want 0, got %d", actualProduct)
	}
	p.Add(8)
	actualProduct = p.GetProduct(2)
	if actualProduct != 32 {
		t.Errorf("Check #4 failed: want 32, got %d", actualProduct)
	}
}
