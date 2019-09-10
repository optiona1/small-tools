package conversion

import (
	"strings"
	"unicode"
)

func ToUnderline(str string) string {
	var temp []string
	for _, v := range str {
		if unicode.IsUpper(v) {
			temp = append(temp, "_"+string(v))
		} else {
			temp = append(temp, string(v))
		}
	}
	return strings.ToLower(strings.Trim(strings.Join(temp, ""), "_"))
}

func ToHump(str string) string {
	temp := strings.Split(str, "_")
	for i, ele := range temp {
		temp[i] = strings.Title(ele)
	}
	temp_str := strings.Join(temp, "")
	return temp_str
}
