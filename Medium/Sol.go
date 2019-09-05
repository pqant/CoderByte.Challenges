package Medium

import "strings"

func PalindromicSubstring(str string) string {
	result := string(none)
	if len(str) == 0 {
		return result
	}
	chars := strings.Split(str, "")
	reverse := func(val string) string {
		if len(val) == 0 {
			return ""
		}
		MAX := len(val) / 2
		valChars := strings.Split(val, "")
		for u := 0; u < MAX; u++ {
			temp := valChars[u]
			valChars[u] = valChars[len(valChars)-u-1]
			valChars[len(valChars)-u-1] = temp
		}
		return strings.Join(valChars, "")
	}
	result = ""
	for u := 0; u < len(chars); u++ {
		for y := u + 1; y <= len(chars); y++ {
			temp := strings.Join(chars[u:y], "")
			if len(temp) > 2 && reverse(temp) == temp {
				if len(temp) > len(result) {
					result = temp
				}
			}
		}
	}
	if len(result) == 0 {
		result = string(none)
	}
	return result
}

type result string

const (
	none result = "none"
)
