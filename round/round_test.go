package round

import (
    "testing"
    "strconv"
    "math"

    "github.com/optiona1/small-tools/assert"
)

func TestRoundFixed(t *testing.T) {
    assert.Assert(strconv.FormatFloat(RoundFixed(math.Pi, 1), 'g', -1, 64) == "3.1", "Not Equal")
    assert.Assert(strconv.FormatFloat(RoundFixed(math.Pi, 2), 'g', -1, 64) == "3.14", "Not Equal")
    assert.Assert(strconv.FormatFloat(RoundFixed(math.Pi, 3), 'g', -1, 64) == "3.142", "Not Equal")
    assert.Assert(strconv.FormatFloat(RoundFixed(math.Pi, 4), 'g', -1, 64) == "3.1416", "Not Equal")
    assert.Assert(strconv.FormatFloat(RoundFixed(math.Pi, 5), 'g', -1, 64) == "3.14159", "Not Equal")
    assert.Assert(strconv.FormatFloat(RoundFixed(math.Pi, 6), 'g', -1, 64) == "3.141593", "Not Equal")
    assert.Assert(strconv.FormatFloat(RoundFixed(math.Pi, 7), 'g', -1, 64) == "3.1415927", "Not Equal")
}
