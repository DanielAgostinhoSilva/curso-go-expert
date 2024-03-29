package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{0, 0},
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, calcTaxItem := range table {
		result := CalculateTax(calcTaxItem.amount)
		if result != calcTaxItem.expected {
			t.Errorf("Expected %f but got %f", calcTaxItem.expected, result)
		}
	}
}
