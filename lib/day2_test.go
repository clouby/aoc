package lib

import "testing"

func TestDayTwoExec(t *testing.T) {
    got := DayTwoExec("../assets/day2.test.txt")
    want := 15
    
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}
