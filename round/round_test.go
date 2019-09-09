package round

import (
    "testing"
    "strconv"
    "math"
)

func TestRoundFixed(t *testing.T) {
    t.Log(strconv.FormatFloat(RoundFixed(math.Pi, 1), 'g', -1, 64) == "3.1")
    t.Log(strconv.FormatFloat(RoundFixed(math.Pi, 2), 'g', -1, 64) == "3.14")
    t.Log(strconv.FormatFloat(RoundFixed(math.Pi, 3), 'g', -1, 64) == "3.142")
    t.Log(strconv.FormatFloat(RoundFixed(math.Pi, 4), 'g', -1, 64) == "3.1416")
    t.Log(strconv.FormatFloat(RoundFixed(math.Pi, 5), 'g', -1, 64) == "3.14159")
    t.Log(strconv.FormatFloat(RoundFixed(math.Pi, 6), 'g', -1, 64) == "3.141593")
    t.Log(strconv.FormatFloat(RoundFixed(math.Pi, 7), 'g', -1, 64) == "3.1415927")
}
