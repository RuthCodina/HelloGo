package functions

import (
	"testing"
)

func Test_doMath(t *testing.T) {
	num := doMath(1, 2, func(i1, i2 int) int { return 1 + 2 })
	if num != 3 {
		t.Errorf("incorrect, got: %d, want:%d.", num, 3)
	}
}

func Test_add(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want:%d.", total, 10)
	}
}

func Test_subtract(t *testing.T) {
	total := subtract(5, 5)
	if total != 0 {
		t.Errorf("Sum was incorrect, got: %d, want:%d.", total, 0)
	}
}
