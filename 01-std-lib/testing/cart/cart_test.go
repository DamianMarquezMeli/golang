package main_test

import (
	"testing"

	cart "github.com/devpablocristo/go-concepts/std-lib/testing/cart"
)

func TestSumItems(t *testing.T) {

	items := make(map[string]int)

	items["apple"] = 1
	items["orange"] = 2
	items["pear"] = 3

	total := cart.SumItems(items)

	if total != 6 {
		t.Errorf("Sum incorrect %d + %d + % d, got: %d, expected: %d.", items["apple"], items["orange"], items["pear"], total, 6)
	}
}
