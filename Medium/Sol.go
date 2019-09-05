package Medium

import "strings"

func PalindromicSubstringReducingWayRec(str string) string {
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
	results := make([][]string, 0)
	selectedSoFar := make([]string, 0)
	var permutation func([]string, int, []string, *[][]string)
	permutation = func(items []string, breakPoint int, selectedSoFar []string, results *[][]string) {
		if len(items) == breakPoint {
			str := strings.Join(selectedSoFar, "")
			if len(selectedSoFar) > 2 && reverse(str) == str && len(str) > len(result) {
				result = str
				temp := make([]string, len(selectedSoFar))
				copy(temp, selectedSoFar)
				*results = [][]string{temp}
			}
			return
		}
		selectedSoFar = append(selectedSoFar, items[breakPoint])
		permutation(items, breakPoint+1, selectedSoFar, results)
		selectedSoFar = selectedSoFar[:len(selectedSoFar)-1]
		permutation(items, breakPoint+1, selectedSoFar, results)
	}
	permutation(chars, 0, selectedSoFar, &results)
	if len(results) != 0 {
		if len(results[0]) != 0 {
			return strings.Join(results[0], "")
		} else {
			return string(none)
		}
	}
	return string(none)
}

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
