package lib

import "testing"

func TestDayOneExec(t *testing.T) {
	got := DayOneExec("../assets/day1.test.txt")
	want := 24000

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
