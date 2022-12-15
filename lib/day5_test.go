package lib

import (
	"testing"
)

func TestDayFifthExec(t *testing.T) {
    mapping := make(map[string][]interface{})

    mapping["1"] = append(mapping["1"], "Z", "N")
    mapping["2"] = append(mapping["2"], "M", "C", "D")
    mapping["3"] = append(mapping["3"], "P")

	got := DayFifthExec("../assets/day5.test.txt", mapping)
	want := "CMZ"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
