package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "Nics",
		Price: 1.00,
		SKU:   "abc-dfs-dfg",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
