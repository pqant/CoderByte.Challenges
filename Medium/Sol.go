package Medium

import (
	"fmt"
	"sort"
	"strconv"
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
		val = ""
		for _, value := range results {
			val += fmt.Sprintf("%v\n", value)
		}
	}
	return val
}

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

func duplicate_count(s1 string) int {
	temp := make(map[string]int, 0)
	items := strings.Split(s1, "")
	for _, v := range items {
		if val, exist := temp[string(v)]; exist {
			temp[string(v)] = val + 1
		} else {
			temp[string(v)] = 1
		}
	}
	count := 0
	for _, v := range temp {
		if v > 1 {
			count++
		}
	}
	return count
}

func BinarySearch(arr []int, isArrSorted bool, search int) bool {
	if len(arr) == 0 {
		return false
	}
	if !isArrSorted {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	}
	if len(arr) < 2 {
		if arr[0] == search {
			return true
		}
		return false
	}
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == search {
			return true
		} else if arr[mid] > search {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func TwoNumberSum(array []int, target int) []int {
	if len(array) < 2 {
		return make([]int, 0)
	}

	binarySearch := func(arr []int, isArrSorted bool, search int) bool {
		//fmt.Printf("%v\n",arr)
		if len(arr) == 0 {
			return false
		}
		if !isArrSorted {
			sort.Slice(arr, func(i, j int) bool {
				return arr[i] < arr[j]
			})
		}
		if len(arr) < 2 {
			if arr[0] == search {
				return true
			}
			return false
		}
		left := 0
		right := len(arr) - 1
		for left <= right {
			mid := (left + right) / 2
			if arr[mid] == search {
				return true
			} else if arr[mid] > search {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return false
	}
	for i := 0; i < len(array)-1; i++ {
		newTarget := target - array[i]
		if binarySearch(array[i+1:], false, newTarget) {
			if array[i] > newTarget {
				return []int{array[i], newTarget}
			} else {
				return []int{newTarget, array[i]}
			}
		}
	}
	return make([]int, 0)
}

func TwoNumberSum_V2(array []int, target int) []int {
	if len(array) < 2 {
		return make([]int, 0)
	}
	sorter := func(arr *[]int) {
		sort.Slice(*arr, func(i, j int) bool {
			return (*arr)[i] < (*arr)[j]
		})
	}
	binarySearch := func(arr []int, isArrSorted bool, search int) bool {
		if len(arr) == 0 {
			return false
		}
		if !isArrSorted {
			sorter(&arr)
		}
		if len(arr) < 2 {
			if arr[0] == search {
				return true
			}
			return false
		}
		left := 0
		right := len(arr) - 1
		for left <= right {
			mid := (left + right) / 2
			if arr[mid] == search {
				return true
			} else if arr[mid] > search {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return false
	}
	for i := 0; i < len(array)-1; i++ {
		newTarget := target - array[i]
		if binarySearch(array[i+1:], false, newTarget) {
			result := []int{array[i], newTarget}
			sorter(&result)
			return result
		}
	}
	return make([]int, 0)
}

func TwoNumberSum_V3(array []int, target int) []int {
	kv := make(map[int]interface{}, 0)
	for _, value := range array {
		lookingFor := target - value
		if _, isExists := kv[lookingFor]; isExists {
			result := make([]int, 2)
			result[0], result[1] = value, lookingFor
			sort.Slice(result, func(i, j int) bool {
				return result[i] < result[j]
			})
			return result
		} else {
			kv[value] = new(interface{})
		}
	}

	return make([]int, 0)
}

func MissingDigitII(str string) string {
	type MyResult string
	const (
		noanswer MyResult = ""
	)
	equalSign, questionMark := "=", "?"
	if !strings.Contains(str, equalSign) {
		return string(noanswer)
	}
	str = strings.Replace(str, " ", "", -1)
	parts := strings.Split(str, "=")
	if len(parts) != 2 {
		return string(noanswer)
	}
	splitter := []string{"*", "/", "-", "+"}
	calculator := func(num1, num2 int, operator string) int {
		if operator == splitter[0] {
			return num1 * num2
		} else
		if operator == splitter[1] {
			return num1 / num2
		} else
		if operator == splitter[2] {
			return num1 - num2
		} else
		if operator == splitter[3] {
			return num1 + num2
		}
		return -1
	}
	indexSign := 0
	indexPart, leftPart := 0, true
out:
	for {
		if strings.Contains(parts[indexPart], splitter[indexSign]) {
			partInside := strings.Split(parts[indexPart], splitter[indexSign])
			if len(partInside) != 2 {
				return string(noanswer)
			}
			indexForIn, left := 0, true
			for {
				if strings.Contains(partInside[indexForIn], questionMark) {
					var newVal1, newVal2, resVal string
					if left {
						newVal1 = partInside[0]
						newVal2 = partInside[1]
					} else {
						newVal1 = partInside[1]
						newVal2 = partInside[0]
					}
					if leftPart {
						resVal = parts[1]
					} else {
						resVal = parts[0]
					}
					newVal1Bck, newVal2Back, resValBck := newVal1, newVal2, resVal

					for u := 0; u < 10; u++ {
						newVal1, newVal2 = newVal1Bck, newVal2Back
						if strings.Split(newVal1, "")[0] == questionMark && u == 0 {
							continue
						}
						newVal1 = strings.Replace(newVal1, "?", strconv.Itoa(u), -1)
						intnewVal1, err1 := strconv.Atoi(newVal1)
						if err1 != nil {
							return string(noanswer)
						}
						newVal2 = strings.Replace(newVal2, "?", strconv.Itoa(u), -1)
						intnewVal2, err2 := strconv.Atoi(newVal2)
						if err2 != nil {
							return string(noanswer)
						}
						result := calculator(intnewVal1, intnewVal2, splitter[indexSign])

						for i := 0; i < 10; i++ {
							resVal = resValBck
							if strings.Split(resVal, "")[0] == questionMark && i == 0 {
								continue
							}
							resVal = strings.Replace(resVal, "?", strconv.Itoa(i), -1)
							intResVal, err3 := strconv.Atoi(resVal)
							if err3 != nil {
								return string(noanswer)
							}
							if result == intResVal {
								if leftPart {
									return fmt.Sprintf("%v %v", u, i)
								} else {
									return fmt.Sprintf("%v %v", i, u)
								}
							}
						}
					}
					break
				}
				indexForIn++
				left = false
			}
			break out
		}
		indexSign++
		if indexSign == len(splitter) {
			leftPart = !leftPart
			if indexPart == 0 {
				indexPart = 1
			} else {
				indexPart = 0
			}
			indexSign = 0
		}
	}
	return string(noanswer)
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

const (
	none result = "none"
)
