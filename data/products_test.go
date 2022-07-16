package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := Product{
		Name:  "bitch",
		Price: 1.00,
		SKU:   "abs-sdf-sff",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
