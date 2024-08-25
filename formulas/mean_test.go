package formula

import (
	"testing"
)

func TestCalculateMean(t *testing.T) {
	got := CalculateMean([]float64{8.0, 5.0, 7.0, 4.0})
	want := 6.0
	if got != want {
		t.Errorf("Test failed")
	}
}
