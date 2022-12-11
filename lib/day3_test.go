package lib

import "testing"

func TestDayThreeExec(t *testing.T) {
	got := DayThreeExec("../assets/day3.test.txt")
	want := 157

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
