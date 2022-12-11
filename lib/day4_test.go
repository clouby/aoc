package lib

import "testing"

func TestDayFourthExec(t *testing.T) {
    got := DayFourthExec("../assets/day4.test.txt")
    want := 2

    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}
