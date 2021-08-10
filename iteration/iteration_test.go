package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("x")
	want := "xxxxx"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}