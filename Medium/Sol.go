package Medium

import (
	"fmt"
	"strings"
)

func PalindromicSubstringReducingWayRec(str string) string {
	result := string(none)
	if len(str) == 0 {
		return result
	}
	chars := strings.Split(str, "")
	reverse := func(val *string) string {
		if len(*val) == 0 {
			return ""
		}
		MAX := len(*val) / 2
		valChars := strings.Split(*val, "")
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
	var powers func([]string, int, []string, *[][]string)
	powers = func(items []string, breakPoint int, selectedSoFar []string, results *[][]string) {
		if len(items) == breakPoint {
			str := strings.Join(selectedSoFar, "")
			if len(selectedSoFar) > 2 && len(str) > len(result) && reverse(&str) == str {
				result = str
				temp := make([]string, len(selectedSoFar))
				copy(temp, selectedSoFar)
				*results = [][]string{temp}
			}
			return
		}
		selectedSoFar = append(selectedSoFar, items[breakPoint])
		powers(items, breakPoint+1, selectedSoFar, results)
		selectedSoFar = selectedSoFar[:len(selectedSoFar)-1]
		powers(items, breakPoint+1, selectedSoFar, results)
	}
	powers(chars, 0, selectedSoFar, &results)
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
	reverse := func(val *string) string {
		if len(*val) == 0 {
			return ""
		}
		MAX := len(*val) / 2
		valChars := strings.Split(*val, "")
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
			if len(temp) > 2 && reverse(&temp) == temp {
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

func FindCouple(items []int, max int, coupleLenght int) string {
	val := "none"
	if len(items) == 0 {
		return val
	}
	results := make([][]int, 0)
	selectedSoFar := make([]int, 0)
	sum := func(all ...int) int {
		sum := 0
		for _, value := range all {
			sum += value
		}
		return sum
	}
	var powerset func([]int, int, []int, *[][]int, int)
	powerset = func(arr []int, breakPoint int, selectedSoFar []int, results *[][]int, maxItem int) {
		if breakPoint == len(arr) {
			if maxItem == len(selectedSoFar) && sum(selectedSoFar...) == max {
				temp := make([]int, len(selectedSoFar))
				copy(temp, selectedSoFar)
				*results = append(*results, temp)
			}
			return
		}
		selectedSoFar = append(selectedSoFar, arr[breakPoint])
		powerset(arr, breakPoint+1, selectedSoFar, results, coupleLenght)
		selectedSoFar = selectedSoFar[:len(selectedSoFar)-1]
		powerset(arr, breakPoint+1, selectedSoFar, results, coupleLenght)
	}
	powerset(items, 0, selectedSoFar, &results, coupleLenght)
	if len(results) != 0 {
		val =""
		for _, value := range results {
			val += fmt.Sprintf("%v\n", value)
		}
	}
	return val
}



const (
	none result = "none"
)

func FindWhichCharInserted(first, second string) string {
	if len(first) >= len(second) {
		return ""
	}
	firstArr := strings.Split(first, "")
	diff := second
	for u := 0; u < len(firstArr); u++ {
		if strings.Index(second, firstArr[u]) != -1 {
			diff = strings.Replace(diff, firstArr[u], "", 1)
		} else {
			return firstArr[u]
		}
	}
	return string(diff[0])
}

func WovelReverser(text string) string {
	if len(text) == 0 {
		return ""
	}
	vowels := []string{"a", "u", "i", "ö", "o"}
	isVowel := func(char string) bool {
		for _, value := range vowels {
			if value == strings.ToLower(char) {
				return true
			}
		}
		return false
	}
	reverse := func(val string) string {
		MAX := len(val) / 2
		strVal := strings.Split(val, "")
		for i := 0; i < MAX; i++ {
			if isVowel(strVal[i]) {
				temp := strVal[i]
				strVal[i] = strVal[len(strVal)-i-1]
				strVal[len(strVal)-i-1] = temp
			}
		}
		return strings.Join(strVal, "")
	}
	return reverse(text)
}


type result string




