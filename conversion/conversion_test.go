package conversion

import (
    "testing"
    "github.com/optiona1/assert"
)

func TestToHump(t *testing.T) {
    str1 := "abc_def"
    str2 := "a_b_c_d"

    assert.Assert(ToHump(str1) == "AbcDef", "This test failed")
    assert.Assert(ToHump(str1) == "ABCD", "This test failed")
}

func TestToUnderline(t *testing.T) {
    str1 := "AbcDef"
    str1 := "ABCD"

    assert.Assert(ToUnderline(str1) == "abc_def", "This test failed")
    assert.Assert(ToUnderline(str2) == "a_b_c_d", "This test failed")
}
