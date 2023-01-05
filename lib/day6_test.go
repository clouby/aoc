package lib

import "testing"

func TestDaySixExec(t *testing.T) {
    got := DaySixExec("../assets/day6.test.txt")
    want := 29

    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}
