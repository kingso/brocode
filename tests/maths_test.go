package tests

import (
	"testing"

	. "github.com/kingso/brocode/internal"
)

func TestDivideByZero(t *testing.T) {
	want := "Divide by zero"
	_, got := DivideBy(1, 0)
	if want != got.Error() {
		t.Error("DivideBy Zero should return an error")
	}
}
func TestDivideByNumber(t *testing.T) {
	want := 0.1
	got, _ := DivideBy(1, 10)
	if want != got {
		t.Error("DivideBy Number 1 / 10 should return 0.1")
	}
}
