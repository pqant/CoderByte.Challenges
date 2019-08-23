package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"bufio"
	"context"
	"log"
	"os"
	"time"
)

var vowels = []rune{'a', 'e', 'i', 'o', 'u'}

func DivisionStringified(num1 int, num2 int) string {
	if num2 == 0 {
		return "0" // divbyZero??!!
	}
	result := fmt.Sprintf("%.0f", math.Round(float64(num1)/float64(num2)))
	if len(result) <= 3 {
		return result
	}
	t := ""
	index := 0
	for u := len(result) - 1; u >= 0; u-- {
		if index != 0 && index%3 == 0 {
			t += ","
		}
		t += string(result[u])
		index++
	}
	result = ""
	for j := len(t) - 1; j >= 0; j-- {
		result += string(t[j])
	}

	return result
}

//Mod (tepe değer)	3  (max frequency )
//Mean , ortlama değer
//Medyan, median (ortanca)	3

func MeanMode(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	for u := 0; u < len(arr); u++ {
		if arr[u] < 0 {
			return 0
		}
	}
	//mean
	mean := 0
	for u := 0; u < len(arr); u++ {
		mean += arr[u]
	}
	mean /= len(arr)
	//sorting for median!
	for u := 0; u < len(arr); u++ {
		for y := 0; y < len(arr); y++ {
			if arr[u] < arr[y] {
				arr[u], arr[y] = arr[y], arr[u]
			}
		}
	}
	//sorted
	median := 0.0
	if len(arr)%2 == 1 {
		median = float64(arr[len(arr)/2])
	} else {
		median = float64(arr[(len(arr)/2)-1]+arr[(len(arr) / 2)]) / 2
	}
	frequency := make(map[int]int, len(arr))
	for u := 0; u < len(arr); u++ {
		val, isExist := frequency[arr[u]]
		if isExist {
			frequency[arr[u]] = val + 1
		} else {
			frequency[arr[u]] = 1
		}
	}
	max := 0
	mode := 0
	for key, value := range frequency {
		if value > max {
			max = value
			mode = key
		}
	}
	if float64(mode) == median {
		return 1
	}
	return 0
}

func rec(items []string) (values []string) {
	keepSearch := false
	for _, value := range items {
		isAdded := false
	keepLooking:
		for _, char := range punchItems {
			tSplitter := strings.Split(value, string(char))
			if len(tSplitter) != 1 {
				isAdded = true
				for _, valueSplit := range tSplitter {
					values = append(values, string(valueSplit))
				}
				keepSearch = true
				break keepLooking
			}
		}
		if !isAdded {
			values = append(values, string(value))
			isAdded = true
		}
	}
	if keepSearch {
		return rec(values)
	} else {
		return values
	}
}

func ThreeNumbers(str string) bool {
	if len(str) == 0 {
		return false
	}
	isNumber := func(value string) bool {
		if value == "0" || value == "1" || value == "2" || value == "3" || value == "4" ||
			value == "5" || value == "6" || value == "7" || value == "8" || value == "9" {
			return true
		}
		return false
	}
	totalNumberCount := 0
	for u := 0; u < len(str); u++ {
		if isNumber(string(str[u])) {
			totalNumberCount++
		}
	}
	if totalNumberCount < 3 {
		return false
	}

	items := strings.Split(str, " ")

	if len(items) == 0 {
		return false
	}
	for _, value := range items {
		xVal := string(value)
		n0, n1, n2, n3, n4, n5, n6, n7, n8, n9 := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
		for u := 0; u < len(xVal); u++ {
			if isNumber(string(xVal[u])) {
				if string(xVal[u]) == "0" {
					n0 += 1
					if n0 > 1 {
						return false
					}
				} else if string(xVal[u]) == "1" {
					n1 += 1
					if n1 > 1 {
						return false
					}
				} else if string(xVal[u]) == "2" {
					n2 += 1
					if n2 > 1 {
						return false
					}
				} else if string(xVal[u]) == "3" {
					n3 += 1
					if n3 > 1 {
						return false
					}
				} else if string(xVal[u]) == "4" {
					n4 += 1
					if n4 > 1 {
						return false
					}
				} else if string(xVal[u]) == "5" {
					n5 += 1
					if n5 > 1 {
						return false
					}
				} else if string(xVal[u]) == "6" {
					n6 += 1
					if n6 > 1 {
						return false
					}
				} else if string(xVal[u]) == "7" {
					n7 += 1
					if n7 > 1 {
						return false
					}
				} else if string(xVal[u]) == "8" {
					n8 += 1
					if n8 > 1 {
						return false
					}
				} else if string(xVal[u]) == "9" {
					n9 += 1
					if n9 > 1 {
						return false
					}
				}
			}
			if u > 0 && u < len(xVal)-1 {
				//fmt.Printf("%v ->>>  %v - %v - %v\n",xVal,string(xVal[u-1]),string(xVal[u]),string(xVal[u+1]))
			}

			if u > 0 && u < len(xVal)-1 &&
				isNumber(string(xVal[u-1])) &&
				isNumber(string(xVal[u])) &&
				isNumber(string(xVal[u+1])) {
				return false
			}
		}
	}
	return true
}

func TwoSum(arr []int) string {
	if len(arr) < 3 {
		return "-1"
	}
	sum := arr[0]
	var pairs []string
	for u := 1; u < len(arr); u++ {
		for y := u + 1; y < len(arr); y++ {
			if sum == arr[u]+arr[y] {
				pairs = append(pairs, fmt.Sprintf("%v,%v", arr[u], arr[y]))
			}
		}
	}
	if len(pairs) == 0 {
		return "-1"
	}
	return strings.Join(pairs, " ")
}

/*
"cats AND*Dogs-are Awesome -> "CatsAndDogsAreAwesome"
"Daniel LikeS-coding"      -> "DanielLikesCoding"

*/
//noinspection ALL
func DifferentCases(str string) string {
	var punchItems = []string{
		" ",
		"-",
		"%",
		"’'",
		"()[]{}<>",
		":",
		",",
		"+",
		"‒–—―",
		"…",
		"!",
		".",
		"«»",
		"-‐",
		"?",
		"‘’“”",
		";",
		"/",
		"⁄",
		"␠",
		"·",
		"&",
		"@",
		"*",
		"\\",
		"•",
		"^",
		"¤¢$€£¥₩₪",
		"†‡",
		"°",
		"¡",
		"¿",
		"¬",
		"#",
		"№",
		"%‰‱",
		"¶",
		"′",
		"§",
		"~",
		"¨",
		"_",
		"|¦",
		"⁂",
		"☞",
		"∴",
		"‽",
		"※"}

	if len(str) == 0 {
		return ""
	}
	isIllegalChar := func(val string) bool {
		for _, value := range punchItems {
			if string(value) == val {
				return true
			}
		}
		return false
	}

	splitted := strings.Split(str, " ")
	splitted = rec(splitted)
	var words []string
	for _, value := range splitted {
		if !isIllegalChar(string(value)) {
			var temp []rune
			t := string(value)
			for u := 0; u < len(t); u++ {
				temp = append(temp, rune(t[u]))
			}
			for u := 0; u < len(temp); u++ {
				if u == 0 {
					if temp[u] >= 97 && temp[u] <= 122 {
						temp[u] = rune(temp[u] - 32)
					}
				} else {
					if temp[u] >= 65 && temp[u] < 97 {
						temp[u] = rune(temp[u] + 32)
					}
				}
			}
			newWord := ""
			for _, value := range temp {
				newWord += fmt.Sprintf("%c", rune(value))
			}

			words = append(words, newWord)
		}
	}
	return strings.Join(words, "")

}

func EqualSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	return true
}

func HammingDistance(strArr []string) int {
	if len(strArr) != 2 {
		return 0
	}
	if len(strArr[0]) != len(strArr[1]) {
		return 0
	}
	result := 0
	for u := 0; u < len(strArr[0]); u++ {
		if strArr[0][u] != strArr[1][u] {
			result++
		}
	}
	return result
}

//[]int {1,2,3,4}
func Superincreasing(arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	for u := 0; u < len(arr); u++ {
		internalSum := 0
		if u != 0 {
			for y := 0; y < u; y++ {
				internalSum += arr[y]
			}
			if internalSum >= arr[u] {
				return false
			}
		}
	}
	return true
}

func OverlappingRanges(arr []int) bool {
	if len(arr) != 5 {
		return false
	}
	a := arr[0]
	b := arr[1]
	c := arr[2]
	d := arr[3]
	x := arr[4]
	var numbers1 []int
	var numbers2 []int
	for u := a; u <= b; u++ {
		numbers1 = append(numbers1, u)
	}
	for u := c; u <= d; u++ {
		numbers2 = append(numbers2, u)
	}
	counter := 0
	for _, value1 := range numbers1 {
		for _, value2 := range numbers2 {
			if value1 == value2 {
				counter++
				if counter >= x {
					return true
				}
			}
		}
	}
	return false
}

//["5","4","6","E","1","7","E","E","3","2"] ->  4,1,5
//["1","2","E","E","3"]                     ->  1,2
func OffLineMinimum(strArr []string) string {
	if len(strArr) == 0 {
		return ""
	}
	for _, value := range strArr {
		if string(value) == "E" {
			break
			// at least there is one E in the array!
		}
	}
	var items []string
	var tempArr []string
	for _, value := range strArr {
		if string(value) == "E" {
			if len(tempArr) != 0 {
				for u := 0; u < len(tempArr); u++ {
					for y := 0; y < len(tempArr); y++ {
						if tempArr[u] < tempArr[y] {
							tempArr[u], tempArr[y] = tempArr[y], tempArr[u]
						}
					}
				}
				var tempValue = tempArr[0]
			exit:
				for _, value := range tempArr {
					isFound := false
					for _, valueItem := range items {
						if string(value) == string(valueItem) {
							isFound = true
							break
						}
					}
					if !isFound {
						tempValue = value
						break exit
					}
				}
				items = append(items, tempValue)
			}
		} else {
			tempArr = append(tempArr, string(value))
		}
	}
	return strings.Join(items, ",")
}

func MultiplicativePersistence(num int) int {
	val := fmt.Sprintf("%v", num)
	if len(val) < 2 {
		return 0
	}
	index := 0
do:
	var numbers []int
	for _, value := range val {
		numberVal, _ := strconv.Atoi(string(value))
		numbers = append(numbers, numberVal)
	}
	resultTemp := 1
	for _, value := range numbers {
		resultTemp *= value
	}
	index++
	if len(fmt.Sprintf("%v", resultTemp)) != 1 {
		val = fmt.Sprintf("%v", resultTemp)
		goto do
	}

	return index
}

func ChangingSequence(arr []int) int {
	if len(arr) < 3 {
		return -1
	}
	inc := false
	if arr[1] > arr[0] {
		inc = true
	}
	index := -1
exit:
	for u := 0; u < 1; u++ {
		for y := u + 1; y < len(arr); y++ {
			if inc {
				if arr[y] > arr[y-1] {
					// increasing..!!!
					index = y
				} else {
					break exit
				}
			} else {
				if arr[y] < arr[y-1] {
					// decreasing..!!!
					index = y
				} else {
					break exit
				}
			}
		}
	}
	if len(arr)-1 == index {
		return -1
	}
	return index
}

func IntSliceToStrSlice(items []int) (str []string) {
	for _, value := range items {
		str = append(str, strconv.Itoa(value))
	}
	return
}

func OtherProducts(arr []int) string {
	if len(arr) < 1 && len(arr) > 10 {
		return ""
	}
	var results []string
	for u := 0; u < len(arr); u++ {
		partialMul := 1
		for y := 0; y < len(arr); y++ {
			if u != y {
				partialMul *= arr[y]
			}
		}
		results = append(results, fmt.Sprintf("%d", partialMul))
	}
	return strings.Join(results, "-")
}

func FoodDistribution(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	//fmt.Printf("INPUT  : [%v]\n", strings.Join(IntSliceToStrSlice(arr), ","))
	mean := 0
	var newArr []int
	for index, value := range arr {
		if index == 0 {
			continue
		}
		mean += int(value)
		newArr = append(newArr, int(value))
	}
	mean /= len(arr) - 1

	stock := arr[0]
	if stock == 0 {
		return stock
	}

	sort.Slice(newArr, func(i, j int) bool {
		return newArr[i] > newArr[j]
	})

	sum := func(arr []int) (sum int) {
		for _, value := range arr {
			sum += int(value)
		}
		return
	}

	if sum(newArr) <= stock {
		return 0
	}

	diffCalc := func(myArr []int) (diff float64) {
		for u := 0; u < len(newArr); u++ {
			for y := u + 1; y < len(newArr); {
				diff += math.Abs(float64(newArr[y] - newArr[u]))
				break
			}
		}
		return
	}
	var result []float64
	mean += 1
reCalc:
	for u := 0; u < len(newArr); u++ {
		if stock == 0 {
			break
		}
		if newArr[u] > mean {
			if stock >= newArr[u]-mean {
				stock -= newArr[u] - mean
				newArr[u] = mean
			} else {
				newArr[u] -= stock
				stock = 0
			}
		}
	}

	calc := diffCalc(newArr)
	printOut := func(sign string, meanVal int, calcRes float64, active bool) {
		if !active {
			return
		}
		var temp []int
		temp = append(temp, stock)
		for u := 0; u < len(newArr); u++ {
			temp = append(temp, newArr[u])
		}
		fmt.Printf("%v - M [%v] - OUTPUT : %v [diff : %v]\n", sign, meanVal, temp, calcRes)
	}

	if calc != 0 {
		result = append(result, calc)
		mean--
		if mean >= 0 {
			printOut("*", mean, calc, print)
			goto reCalc
		}
	} else {
		printOut("|", mean, 0, print)
		result = nil
	}

	if len(result) != 0 {
		sort.Slice(result, func(i, j int) bool {
			return result[i] < result[j]
		})
		calc = result[0]
	}

	return int(calc)

}

var (
	print = true
)

//a1 > a2 < a3 > a4 < a5 > a6 ...

func NextPalindrome(num int) int {
	if num < 0 {
		return 0
	}
	checkPalindrome := func(text string) bool {
		reverse := func(val string) string {
			if len(val) == 0 {
				return ""
			}
			var items []string
			for _, value := range val {
				items = append(items, string(value))
			}
			MAX := len(items) / 2
			for u := 0; u < MAX; u++ {
				temp := items[u]
				items[u] = items[len(items)-u-1]
				items[len(items)-u-1] = temp
			}
			return strings.Join(items, "")
		}
		if reverse(text) == text {
			return true
		}
		return false
	}

	if len(strconv.Itoa(num)) == 1 {
		num++
		if len(strconv.Itoa(num)) == 1 {
			return num
		}
	}

	if checkPalindrome(strconv.Itoa(num)) {
		num++
	}

	for !checkPalindrome(strconv.Itoa(num)) {
		num++
	}
	return num
}

func WaveSorting(arr []int) string {
	if len(arr) == 0 {
		return "false"
	}
	res := "true"
	var utility func([]int, int)

	check := func(resItem []int) (r string) {
		sign := true
		r = "true"
		for k := 0; k < len(resItem); k++ {
			if sign {
				if k < len(resItem)-1 && resItem[k] <= resItem[k+1] {
					r = "false"
					break
				}
				sign = false
			} else {
				if k < len(resItem)-1 && resItem[k] >= resItem[k+1] {
					r = "false"
					break
				}
				sign = true
			}
		}
		return
	}
	state := false

	utility = func(arr []int, n int) {
		if n == 1 {
			res = check(arr)
			if res == "true" {
				state = true
				return
			} else {
				fmt.Printf("Trying..%v \n", arr)
			}
		} else {
			for i := 0; i < n; i++ {
				utility(arr, n-1)
				if n%2 == 1 {
					temp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = temp
				} else {
					temp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = temp
				}
			}
		}
	}
	utility(arr, len(arr))

	return fmt.Sprintf("%v", res)
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	var res [][]int

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func NonrepeatingCharacter(str string) string {
	if len(str) < 1 {
		return ""
	}
	for u := 0; u < len(str); u++ {
		found := false
		for y := 0; y < len(str); y++ {
			if u != y && str[u] == str[y] {
				found = true
				break
			}
		}
		if !found {
			return string(str[u])
		}
	}
	return string(str[0])
}

func ProductDigits(num int) int {
	var div []int
	for u := 2; u < num; u++ {
		if num%u == 0 {
			div = append(div, u)
		}
	}
	if len(div) == 0 {
		div = append(div, 1, num)
		return len(strconv.Itoa(num)) + 1
	}

	var lenOps []int
	for u := 0; u < len(div); u++ {
		for y := 0; y < len(div); y++ {
			if u != y && div[u]*div[y] == num {
				lenOps = append(lenOps, len(strconv.Itoa(div[u]))+len(strconv.Itoa(div[y])))
			}
		}
	}
	sort.Slice(lenOps, func(i, j int) bool {
		return lenOps[i] < lenOps[j]
	})
	return lenOps[0]
}

func BinaryReversal(str string) int {
	if len(str) == 0 {
		return 0
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	n := int64(num)
	binary := strconv.FormatInt(n, 2)

	if len(binary)%8 != 0 {
		upperBound := len(binary)
		for upperBound%8 != 0 {
			upperBound++
		}
		numGenerator := func(number string, count int) (r string) {
			if count == 0 {
				return
			}
			for u := 0; u < count; u++ {
				r += number
			}
			return
		}
		binary = numGenerator("0", upperBound-len(binary)) + binary
	}
	var bin []string
	for _, value := range binary {
		bin = append(bin, string(value))
	}

	MAX := len(binary) / 2
	for u := 0; u < MAX; u++ {
		temp := bin[u]
		bin[u] = bin[len(bin)-u-1]
		bin[len(bin)-u-1] = temp
	}

	binToDecimal := func(bin string) (r float64) {
		if len(bin) == 0 {
			r = 0
		}
		for u := 0; u < len(bin); u++ {
			index := len(bin) - 1 - u
			val, _ := strconv.Atoi(string(bin[index]))
			res := math.Pow(2.0, float64(u)) * float64(val)
			r += res
		}
		return
	}

	return int(binToDecimal(strings.Join(bin, "")))
}

func ArrayMatching(strArr []string) string {
	if len(strArr) != 2 {
		return ""
	}
	var coll []string
	cleaner := func(text string) (r []string) {
		rT := strings.Replace(text, "[", "", -1)
		rT = strings.Replace(rT, "]", "", -1)
		rT = strings.Replace(rT, " ", "", -1)
		r = strings.Split(rT, ",")
		return
	}
	strArr1 := cleaner(strArr[0])
	strArr2 := cleaner(strArr[1])

	nConverter := func(number string) int {
		val, _ := strconv.Atoi(number)
		return val
	}

	if len(strArr1) == len(strArr2) {
		for u := 0; u < len(strArr1); u++ {
			coll = append(coll, fmt.Sprintf("%d", nConverter(strArr1[u])+nConverter(strArr2[u])))
		}
	} else {
		if len(strArr1) > len(strArr2) {
			for u := 0; u < len(strArr2); u++ {
				coll = append(coll, fmt.Sprintf("%d", nConverter(strArr1[u])+nConverter(strArr2[u])))
			}
			for u := len(strArr2); u < len(strArr1); u++ {
				coll = append(coll, fmt.Sprintf("%d", nConverter(strArr1[u])))
			}
		} else {
			for u := 0; u < len(strArr1); u++ {
				coll = append(coll, fmt.Sprintf("%d", nConverter(strArr1[u])+nConverter(strArr2[u])))
			}
			for u := len(strArr1); u < len(strArr2); u++ {
				coll = append(coll, fmt.Sprintf("%d", nConverter(strArr2[u])))
			}
		}
	}

	return strings.Join(coll, "-")
}

func RemoveBrackets(str string) int {
	if len(str) == 0 {
		return 0
	}
	leftPCount := 0
	rightPCount := 0
	diffPlus := 0
	j := 0
	for j = 0; j < len(str); j++ {
		if string(str[j:j+1]) == ")" {
			diffPlus++
		}
		if j == 0 && diffPlus == 0 {
			break
		}
	}
	if diffPlus != 0 {
		str = str[diffPlus:]
	}
	for u := 0; u < len(str); u++ {
		val := string(str[u])
		if val == "(" {
			leftPCount++
		}
		if val == ")" {
			rightPCount++
		}
	}
	diff := float64(leftPCount - rightPCount + diffPlus)
	if diff != 0 {
		return int(math.Abs(diff))
	} else {
		return 0
	}
}

func LargestFour(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	sum := 0
	for u := 0; u < len(arr); u++ {
		sum += arr[u]
		if u == 3 {
			break
		}
	}
	return sum
}

func StringPeriods(str string) string {
	if len(str) == 1 || len(str) == 0 {
		return "-1"
	}
	var chars []string
	for _, value := range str {
		chars = append(chars, string(value))
	}
	count := func(val string, search string) (total int) {
		total = 0
		if len(val) == 0 || len(search) == 0 {
			return
		}
		res := 0
		for u := 0; u < len(val); u++ {
			res = strings.Index(val[u:], search)
			if res != -1 {
				total++
				u += res
			}
		}
		return
	}
	_ = count

	kv := make([]string, 0)
	for u := 0; u < len(chars); u++ {
		for y := u + 1; y < len(chars); y++ {
			keyword := strings.Join(chars[u:y], "")
			kv = append(kv, keyword)
		}
	}
	kvNext := make(map[string]int, 0)
	for _, value := range kv {
		newVal := string(value)
		m := 1
		for ; len(newVal) <= len(str); {
			if newVal == str {
				kvNext[string(value)] = m
			}
			newVal += string(value)
			m++
		}
	}

	var keys []string

	for key := range kvNext {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) > len(keys[j])
	})

	if len(keys) != 0 {
		return fmt.Sprintf("%v", keys[0])
	} else {
		return "-1"
	}
}

func VowelSquare_OLD(strArr []string) string {
	if len(strArr) < 2 {
		return "not found"
	}
	for _, value := range strArr {
		if len(value) < 2 {
			return "not found"
		}
	}
	vowelCheck := func(char string) bool {
		if strings.ToLower(char) == "a" || strings.ToLower(char) == "e" ||
			strings.ToLower(char) == "u" || strings.ToLower(char) == "i" ||
			strings.ToLower(char) == "o" || strings.ToLower(char) == "ü" ||
			strings.ToLower(char) == "ö" {
			return true
		}
		return false
	}

	matrixController := func(items []string, sizeOfx int, sizeOfy int) bool {
		colSize := len(items[0])
		rowSize := len(items)
		if colSize < sizeOfx || rowSize < sizeOfy {
			return false
		}

		for u := 0; u < colSize-sizeOfx; u++ {
			for y := 0; y < rowSize-sizeOfy; y++ {
				fmt.Printf("[%v]\n", items[u][y])
			}
		}
		return true
	}

	_ = matrixController

	var newArr []string
	for k := 0; k < len(strArr); k++ {
		newLine := ""
		for n := 0; n < len(strArr[k]); n++ {
			if vowelCheck(string(strArr[k][n])) {
				newLine += "1"
			} else {
				newLine += "0"
			}
		}
		newArr = append(newArr, newLine)
	}

	if len(newArr) == 0 {
		return "not found"
	}
	lenArr := len(newArr[0])

	for u := 0; u < lenArr; u++ {
		first := string(newArr[u])
		for y := 0; y < lenArr; y++ {
			if first == "1" {
				if u < lenArr-1 {
					//secondRowVal := string(newArr[u+1])

				}
			}

		}
	}

	return "0-0"
}

func DistinctCharacters(str string) string {
	if len(str) < 10 {
		return "false"
	}
	kv := make(map[string]int, len(str))
	for _, value := range str {
		newVal := string(value)
		if val, isExist := kv[newVal]; isExist {
			kv[newVal] = val + 1
		} else {
			kv[newVal] = 1
		}
	}
	if len(kv) >= 10 {
		return "true"
	}
	return "false"
}

func ThreeSum(arr []int) string {
	if len(arr) < 4 {
		return "false"
	}
	first := arr[:1][0]
	results := make([][]int, 0)
	soFarList := make([]int, 0)
	var powerSetGenerator func([]int, int, []int, *[][]int)

	powerSetGenerator = func(subset []int, breakPoint int, soFarList []int, result *[][]int) {
		if breakPoint == len(subset) {
			if len(soFarList) == 3 { // powerset just for 3 items!
				sum := 0
				for _, value := range soFarList {
					sum += int(value)
				}
				if sum == first {
					clone := make([]int, 3)
					copy(clone, soFarList)
					*result = append(*result, clone)
				}
			}
			return
		}
		soFarList = append(soFarList, subset[breakPoint])
		powerSetGenerator(subset, breakPoint+1, soFarList, result)
		soFarList = soFarList[:len(soFarList)-1]
		powerSetGenerator(subset, breakPoint+1, soFarList, result)
	}
	powerSetGenerator(arr[1:], 0, soFarList, &results)

	for _, value := range results {
		sum := 0
		for _, val := range value {
			sum += val
		}
		if sum == first {
			return "true"
		}
	}
	return "false"
}

func QuestionsMarks(str string) string {
	if len(str) == 0 {
		return "false"
	}
	isNumberFunc := func(val rune) bool {
		switch val {
		case '0':
			return true
		case '1':
			return true
		case '2':
			return true
		case '3':
			return true
		case '4':
			return true
		case '5':
			return true
		case '6':
			return true
		case '7':
			return true
		case '8':
			return true
		case '9':
			return true
		default:
			return false
		}
	}
	questionCount := 0
	bef, aft := "-1", "-1"
	for _, value := range str {
		val := string(value)
		if val == "?" {
			questionCount++
		} else {
			if isNumberFunc(rune(value)) {
				if bef == "-1" {
					questionCount = 0
					bef = string(value)
				} else {
					aft = string(value)
				}
				if bef != "-1" && aft != "-1" && questionCount == 3 {
					befVal, _ := strconv.Atoi(bef)
					aftVal, _ := strconv.Atoi(aft)
					if befVal+aftVal == 10 {
						questionCount = 0
						bef = "-1"
						aft = "-1"
						return "true"
					} else {
						return "false"
					}
				}
			} else {
				questionCount = 0
			}
		}
	}
	return "false"
}

func BasicRomanNumerals(str string) string {
	if len(str) == 0 {
		return "0"
	}
	romanChars := []string{"M", "D", "C", "L", "X", "V", "I"}

	RtoD := func(rom string) int {
		if len(rom) == 0 {
			return 0
		}
		switch rom {
		case "I":
			return 1
		case "V":
			return 5
		case "X":
			return 10
		case "L":
			return 50
		case "C":
			return 100
		case "D":
			return 500
		case "M":
			return 1000
		default:
			return 0
		}
	}

	var numberList []string

do:
	for u := 0; u < len(romanChars) && len(str) != 0; u++ {
		index := strings.Index(str, string(romanChars[u]))
		if index != -1 {
			numberList = append(numberList, str[:index+1])
			str = str[index+1:]
			if len(str) != 0 {
				goto do
			}
		}
	}

	total := 0
	for u := 0; u < len(numberList); u++ {
		soFar := 0
		for y := 0; y < len(numberList[u])-1; y++ {
			soFar += RtoD(string(numberList[u][y]))
		}
		total += RtoD(string(numberList[u][len(numberList[u])-1]))
		total -= soFar
	}
	return fmt.Sprintf("%v", total)
}

func PowerSetCount(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var powerSetGenerator func([]int, int, []int, *[][]int)
	powerSetGenerator = func(subset []int, breakPoint int, internalList []int, allList *[][]int) {
		if len(subset) == breakPoint {
			*allList = append(*allList, internalList)
			return
		}
		internalList = append(internalList, subset[breakPoint])
		powerSetGenerator(subset, breakPoint+1, internalList, allList)
		internalList = internalList[:len(internalList)-1]
		powerSetGenerator(subset, breakPoint+1, internalList, allList)
	}
	allList := make([][]int, 0)
	var internalList []int
	powerSetGenerator(arr, 0, internalList, &allList)

	return len(allList)
}

func BitwiseTwo(strArr []string) string {
	if len(strArr) != 2 {
		return ""
	}
	if len(strArr[0]) != len(strArr[1]) || (len(strArr[0]) == 0 && len(strArr[2]) == 0) {
		return ""
	}

	andOp := func(val1, val2 string) (result string) {
		result = "0"
		if val1 == "1" && val2 == "1" {
			result = "1"
		}
		return
	}

	strToSlice := func(val string) (result []string) {
		for _, value := range val {
			result = append(result, string(value))
		}
		return
	}

	var res []string
	str1, str2 := strToSlice(strArr[0]), strToSlice(strArr[1])
	for u := 0; u < len(str1); u++ {
		res = append(res, andOp(str1[u], str2[u]))
	}
	return strings.Join(res, "")
}

func LargestPair(num int) int {
	strVal := fmt.Sprintf("%d", num)
	if len(strVal) == 2 {
		return num
	}
	max := 0
	for n := 0; n < len(strVal); n++ {
		for y := n + 1; y < len(strVal)-1; y++ {
			number, _ := strconv.Atoi(strVal[y : y+2])
			if number > max {
				max = number
			}
		}
	}
	return max
}

func LongestIncreasingSequence(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var base []int
	for range arr {
		base = append(base, 1)
	}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if j >= i {
				break
			}
			if arr[j] < arr[i] {
				if base[i] <= base[j] {
					base[i] = base[j] + 1
				}
				//fmt.Printf("%v \n", base)
			}
		}
	}
	max := base[0]
	for u := 1; u < len(base); u++ {
		if base[u] > max {
			max = base[u]
		}
	}
	//fmt.Printf("ARR - %v  - %v \n\n", arr, base)
	return max
}

func EvenPairs(str string) string {
	if len(str) == 0 {
		return "false"
	}

	numberConverter := func(str string) (r bool, v int) {
		r = false
		v = 0
		if len(str) == 0 {
			return
		}
		v, e := strconv.Atoi(str)
		if e != nil {
			return
		}
		r = true
		return
	}

	var all []int
	check := ""
	for _, value := range str {
		opSuccess, number := numberConverter(fmt.Sprintf("%c", rune(value)))
		if opSuccess {
			check = fmt.Sprintf("%v%v", check, number)
		} else {
			if len(check) != 0 {
				num, err := strconv.Atoi(check)
				if err == nil {
					all = append(all, num)
				}
				check = ""
			}
		}
	}
	if len(check) != 0 {
		num, err := strconv.Atoi(check)
		if err == nil {
			all = append(all, num)
		}
		check = ""
	}

	for u := 0; u < len(all); u++ {
		strVal := fmt.Sprintf("%d", all[u])
		index := 1
		for n := 0; n < len(strVal); n++ {
			first := ""
			for _, value := range strVal {
				first += string(value)
				if len(first) == index {
					break
				}
			}
			second := strVal[len(first):]
			if len(second) != 0 {
				p1, e1 := strconv.Atoi(first)
				p2, e2 := strconv.Atoi(second)
				if e1 == nil && e2 == nil {
					if p1%2 == 0 && p2%2 == 0 {
						return "true"
					}
				}
			}
			index++
		}
	}
	return "false"
}

//1, 2, 3, 7, 4, 5

//2, 10, 3, 9, 11, 5
func LongestIncreasingSequence_old(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	kv := make(map[int][]int, len(arr))
	key := -1
	for u := 0; u < len(arr); u++ {
		key = -1
		search := arr[u]
		for y := 1; y < len(arr); y++ {
			if y > u {
				newItem := arr[y]
				if newItem > search {
					if key == -1 {
						key = u // sequence starter!
						if newItem > arr[u] {
							isSmallerFound := false
							for n := y + 1; n < len(arr); n++ {
								temp := arr[n]
								if temp-arr[u] < newItem-arr[u] {
									isSmallerFound = true
									break
								}
							}
							if !isSmallerFound {
								kv[u] = append(kv[u], arr[u], newItem)
							} else {
								kv[u] = append(kv[u], arr[u])
							}
						}
					} else {
						exist, ok := kv[key]
						if ok && len(exist) != 0 {
							last := kv[key][len(kv[key])-1]
							if newItem > last {
								isSmallerFound := false
								for n := y + 1; n < len(arr); n++ {
									temp := arr[n]
									if temp-last < newItem-last {
										isSmallerFound = true
										break
									}
								}
								if !isSmallerFound {
									kv[key] = append(kv[key], newItem)
								} else {
									// if there is smaller , don't add this NEW item!
								}
							}
						}
					}
				}
			}
		}
	}
	max := 0
	fmt.Printf("ARR - %v\n\n", arr)
	for k, value := range kv {
		fmt.Printf("  %v - %v\n", arr[k], value)
		if len(value) > max {
			max = len(value)
		}
	}
	return max
}

func LoveStory(itemsArr []int, sum int) bool {
	result := false
	sort.Slice(itemsArr, func(i, j int) bool {
		return itemsArr[i] < itemsArr[j]
	})
	fmt.Printf("%v\n\n\n", itemsArr)
	var already []string
	isAlreadyChecked := func(items []string, val string) bool {
		for _, value := range items {
			if value == val {
				return true
			}
		}
		return false
	}
exit:
	for u := 0; u < len(itemsArr); u++ {
		for y := u + 1; y < len(itemsArr); y++ {
			if itemsArr[u]+itemsArr[y] < sum && !isAlreadyChecked(already, fmt.Sprintf("%v,%v", itemsArr[u], itemsArr[y])) {
				fmt.Printf("%v-%v -> [%v]\n", itemsArr[u], itemsArr[y], sum-(itemsArr[u]+itemsArr[y]))
				result = BinarySearch(itemsArr, sum-(itemsArr[u]+itemsArr[y]))
				if result {
					break exit
				} else {
					already = append(already, fmt.Sprintf("%v,%v", itemsArr[u], itemsArr[y]))
				}
			}
		}
	}
	return result
}

func BinarySearch(items []int, search int) bool {
	if len(items) == 0 {
		return false
	}
	left := 0
	right := len(items) - 1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if items[mid] == search {
			return true
		} else if items[mid] > search {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func Fact(a int) int {
	if a == 0 {
		return 1
	}
	return a * Fact(a-1)
}

func CharPermutation(items []string) [][]string {
	var result [][]string
	var utility func([]string, int)
	utility = func(items []string, n int) {
		if n == 1 {
			temp := make([]string, len(items))
			copy(temp, items)
			result = append(result, temp)
		} else {
			for i := 0; i < n; i++ {
				utility(items, n-1)
				if n%2 == 1 {
					temp := items[i]
					items[i] = items[n-1]
					items[n-1] = temp
				} else {
					temp := items[0]
					items[0] = items[n-1]
					items[n-1] = temp
				}
			}
		}
	}
	utility(items, len(items))
	return result
}

func CallMe(original []string, items []string, index int, activeIndex int, show bool, limiter int, totalExec int) []string {
	var item string
	var send []string
	if len(items) == 0 {
		return []string{}
	}
	if totalExec > limiter {
		return []string{}
	}
	for _, value := range items {
		item += string(value)
	}
	if show {
		fmt.Printf("%v \n", item)
	}
	if index == 0 {
		index = len(items) - 1
		start := len(items) - index
		send = append(send, original[activeIndex])
		for u := start; u < len(items); u++ {
			send = append(send, string(items[u]))
		}
		if totalExec != 0 && limiter%totalExec == 0 {
			activeIndex++
		}
	} else {
		item = item[1:] + item[:1]
		for _, value := range item {
			send = append(send, fmt.Sprintf("%c", rune(value)))
		}
		index--
	}
	totalExec++
	return CallMe(original, send, index, activeIndex, true, limiter, totalExec)
}

func CalcPermutation(items []int) [][]int {
	var result [][]int
	var util func([]int, int)
	util = func(items []int, n int) {
		if n == 1 {
			temp := make([]int, len(items))
			copy(temp, items)
			result = append(result, temp)
		} else {
			for i := 0; i < n; i++ {
				util(items, n-1)
				if n%2 == 1 {
					temp := items[i]
					items[i] = items[n-1]
					items[n-1] = temp
				} else {
					temp := items[0]
					items[0] = items[n-1]
					items[n-1] = temp
				}
			}
		}
	}
	util(items, len(items))
	return result
}

func Perm(items []int) [][]int {
	var result [][]int
	var utility func([]int, int)
	utility = func(items []int, n int) {
		if n == 1 {
			temp := make([]int, len(items))
			copy(temp, items)
			result = append(result, temp)
		} else {
			for i := 0; i < n; i++ {
				utility(items, n-1)
				if n%2 == 1 {
					temp := items[i]
					items[i] = items[n-1]
					items[n-1] = temp
				} else {
					temp := items[0]
					items[0] = items[n-1]
					items[n-1] = temp
				}
			}
		}
	}
	utility(items, len(items))
	return result
}

func PowerSet(inputSet []int) [][]int {
	powerSet := make([][]int, 0)
	var selectedSoFar []int
	generatePowerItems(inputSet, 0, selectedSoFar, &powerSet)
	return powerSet
}

func generatePowerItems(inputSet []int, decisionPoint int, selectedSoFar []int, powerSet *[][]int) {
	if decisionPoint == len(inputSet) {
		*powerSet = append(*powerSet, selectedSoFar)
		//fmt.Printf("%v\n",*powerSet)
		return
	}
	selectedSoFar = append(selectedSoFar, inputSet[decisionPoint])
	generatePowerItems(inputSet, decisionPoint+1, selectedSoFar, powerSet)
	fmt.Printf("--- %v\n", selectedSoFar)
	selectedSoFar = selectedSoFar[:len(selectedSoFar)-1]
	//fmt.Printf("### %v\n", selectedSoFar)
	generatePowerItems(inputSet, decisionPoint+1, selectedSoFar, powerSet)
}

func AlphabetSearching(str string) string {
	engAlphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r",
		"s", "t", "u", "v", "w", "x", "y", "z"}
	var all []string
	for _, value := range str {
		all = append(all, string(value))
	}
	kv := make(map[string]int, len(engAlphabet))
	for _, value := range engAlphabet {
		kv[string(value)] = 0
	}
	for _, value := range all {
		if count, isExist := kv[string(value)]; isExist {
			kv[string(value)] = count + 1
		} else {
			kv[string(value)] = 1
		}
	}
	for _, value := range kv {
		if value == 0 {
			return "false"
		}
	}
	return "true"
}

//9876541110 --> 6
func OneDecremented(num int) int {
	var nums []int
	for _, value := range fmt.Sprintf("%d", num) {
		nums = append(nums, int(value)-48)
	}
	index := 0
	for u := 1; u < len(nums); u++ {
		if nums[u]-nums[u-1] == -1 {
			index++
		}
	}
	return index
}

func SimpleEvens(num int) bool {
	var nums []int
	for _, value := range fmt.Sprintf("%d", num) {
		nums = append(nums, int(value)-48)
	}
	for u := 0; u < len(nums); u++ {
		if nums[u]%2 != 0 {
			return false
		}
	}
	return true
}

//"cats AND*Dogs-are Awesome" --> "catsAndDogsAreAwesome"
func CamelCase(str string) string {
	if len(str) == 0 {
		return ""
	}
	var (
		illItems = []string{
			" ",
			"’'",
			"-",
			"%",
			"()[]{}<>",
			":",
			",",
			"+",
			"‒",
			"…",
			"!",
			".",
			"«»",
			"-‐",
			"?",
			"‘’“”",
			";",
			"/",
			"⁄",
			"␠",
			"·",
			"&",
			"@",
			"*",
			"\\",
			"•",
			"^",
			"¤¢$€£¥₩₪",
			"†‡",
			"°",
			"¡",
			"¿",
			"¬",
			"#",
			"№",
			"%‰‱",
			"¶",
			"′",
			"§",
			"~",
			"¨",
			"_",
			"|¦",
			"⁂",
			"☞",
			"∴",
			"‽",
			"※"}
	)
	illegalReplacer := func(text string) string {
		for _, char := range illItems {
			text = strings.Replace(text, string(char), " ", -1)
		}
		return text
	}
	str = illegalReplacer(str)
	var items = strings.Split(str, " ")
	result := ""
	index := 0
	for _, value := range items {
		if index == 0 {
			result += strings.ToLower(value[:])
		} else {
			result += strings.ToUpper(value[:1]) + strings.ToLower(value[1:])
		}
		index++
	}
	return result
}

func SumMultiplier(arr []int) string {
	if len(arr) == 0 {
		return "false"
	}
	sum := 0
	for _, value := range arr {
		sum += int(value)
	}
	sum *= 2
	for u := 0; u < len(arr); u++ {
		for y := u + 1; y < len(arr); y++ {
			if arr[u]*arr[y] > sum {
				return "true"
			}
		}
	}
	return "false"
}

func ClosestEnemy(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var enemies []int
	hero := 0
	oneCount := 0
	for key, value := range arr {
		val := int(value)
		if val == 1 {
			oneCount++
			hero = key
		}
		if val == 2 {
			enemies = append(enemies, key)
		}
	}
	if oneCount != 1 {
		return 0
	}
	if len(enemies) == 0 {
		return 0
	}
	min := int(math.Abs(float64(hero) - float64(enemies[0])))
	for u := 1; u < len(enemies); u++ {
		if int(math.Abs(float64(enemies[u])-float64(hero))) < min {
			min = int(math.Abs(float64(enemies[u]) - float64(hero)))
		}
	}
	return int(math.Abs(float64(min)))
}

func NumberStream(str string) string {
	if len(str) == 0 {
		return "false"
	}
	var nums []int
	for _, value := range str {
		val := int(value) - 48
		if val == 0 || val == 1 {
			return "false"
		}
		nums = append(nums, val)
	}

	multiplex := func(text string, count int) string {
		if count == 0 {
			return text
		}
		if len(text) == 0 {
			return ""
		}
		result := ""
		for u := 0; u < count; u++ {
			result += text
		}
		return result
	}

	for u := 0; u < len(nums); u++ {
		result := multiplex(fmt.Sprintf("%d", nums[u]), nums[u])
		if strings.Contains(str, result) {
			return "true"
		}
	}
	return "false"
}

func GCF(arr []int) int {
	if len(arr) != 2 {
		return 0
	}
	if arr[0] < 0 || arr[1] < 0 {
		return 0
	}
	result := 1
	min := arr[0]
	if arr[1] < arr[0] {
		min = arr[1]
	}
	for u := 1; u <= min; u++ {
		if arr[0]%u == 0 && arr[1]%u == 0 {
			result = u
		}
	}
	return result
}

//Input:"cats AND*Dogs-are Awesome" --> cats_and_dogs_are_awesome
func SnakeCase(str string) string {
	if len(str) == 0 {
		return ""
	}
	var (
		illItems = []string{
			" ",
			"’'",
			"-",
			"%",
			"()[]{}<>",
			":",
			",",
			"+",
			"‒",
			"…",
			"!",
			".",
			"«»",
			"-‐",
			"?",
			"‘’“”",
			";",
			"/",
			"⁄",
			"␠",
			"·",
			"&",
			"@",
			"*",
			"\\",
			"•",
			"^",
			"¤¢$€£¥₩₪",
			"†‡",
			"°",
			"¡",
			"¿",
			"¬",
			"#",
			"№",
			"%‰‱",
			"¶",
			"′",
			"§",
			"~",
			"¨",
			"_",
			"|¦",
			"⁂",
			"☞",
			"∴",
			"‽",
			"※"}
	)
	illegalReplacer := func(text string) string {
		for _, char := range illItems {
			text = strings.Replace(text, string(char), " ", -1)
		}
		return text
	}
	str = illegalReplacer(str)
	var items = strings.Split(str, " ")
	var result []string
	for _, value := range items {
		result = append(result, strings.ToLower(value[:]))
	}
	return strings.Join(result, "_")
}

//[]string {"X:-1", "Y:1", "X:-4", "B:3", "X:5"} -> "B:3,Y:1"
func GroupTotals(strArr []string) string {
	if len(strArr) == 0 {
		return ""
	}
	kv := make(map[string]int, 0)
	for _, value := range strArr {
		vals := strings.Split(string(value), ":")
		if len(vals) != 2 {
			return ""
		}
		newVal, err := strconv.Atoi(vals[1])
		if err != nil {
			return ""
		}
		if data, isExist := kv[vals[0]]; isExist {
			kv[vals[0]] = data + newVal
		} else {
			kv[vals[0]] = newVal
		}
	}
	result := ""
	var keys []string
	for key := range kv {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprintf("%d", []rune(keys[i])[0]) < fmt.Sprintf("%d", []rune(keys[j])[0])
	})

	for _, value := range keys {
		if kv[value] != 0 {
			result += fmt.Sprintf("%v:%v,", value, kv[value])
		}
	}
	if len(result) == 0 {
		return ""
	}
	if result[len(result)-1:] == "," {
		result = result[0 : len(result)-1]
	}
	return result
}

func HDistance(strArr []string) int {
	if len(strArr) == 0 {
		return 0
	}
	if len(strArr[0]) != len(strArr[1]) {
		return 0
	}
	var a, b []string
	for _, value := range strArr[0] {
		a = append(a, string(value))
	}
	for _, value := range strArr[1] {
		b = append(b, string(value))
	}
	diff := 0
	for u := 0; u < len(a); u++ {
		if a[u] != b[u] {
			diff++
		}
	}
	return diff
}

func StringMerge(str string) string {
	if len(str) == 0 {
		return ""
	}
	if !strings.Contains(str, "*") {
		return ""
	}
	items := strings.Split(str, "*")
	if len(items) != 2 {
		return ""
	}
	text := ""
	for u := 0; u < len(items[0]); u++ {
		text += string(items[0][u]) + string(items[1][u])
	}
	return text
}

func ASCIIConversion(str string) string {
	if len(str) == 0 {
		return ""
	}

	var all []string
	all = strings.Split(str, " ")

	text := ""
	for _, value := range all {
		for _, value := range string(value) {
			text += fmt.Sprintf("%d", rune(value))
		}
		text += " "
	}
	return strings.TrimRight(text, " ")
}

func ElementMerger(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var newList []int

do:
	for u := 1; u < len(arr); u++ {
		diff := int(math.Abs(float64(arr[u]) - (float64(arr[u-1]))))
		newList = append(newList, diff)
		if u == len(arr)-1 {
			if len(newList) != 1 {
				arr = make([]int, len(newList))
				copy(arr, newList)
				newList = nil
				goto do
			}
		}
	}
	if len(arr) == 2 {
		return int(math.Abs(float64(arr[0]) - (float64(arr[1]))))
	}
	return 0
}

func MatrixMultiply(m1, m2 [][]float32) ([][]float32, error) {
	if len(m1[0]) != len(m2) {
		return [][]float32{}, errors.New("matrix can't be multiply!Dimensions are different!")
	}
	res := make([][]float32, len(m1))
	for i := 0; i < len(m1); i++ {
		res[i] = make([]float32, len(m2[0]))
		for u := 0; u < len(m2[0]); u++ {
			for y := 0; y < len(m2); y++ {
				res[i][u] += m1[i][y] * m2[y][u]
			}
		}
	}
	return res, nil
}

func SerialNumber(str string) string {
	if len(str) == 0 {
		return "false"
	}
	if len(str) > 11 {
		return "false"
	}
	split := strings.Split(str, ".")
	if len(split) != 3 {
		return "false"
	}
	index := 0
	checkValIsLess := func(valIn string) bool {
		if len(valIn) == 0 {
			return false
		}
		last, errLast := strconv.Atoi(valIn[:len(valIn)-1])
		if errLast != nil {
			return false
		}
		for u := 0; u < len(valIn)-1; u++ {
			inVal, errIn := strconv.Atoi(string(valIn[u]))
			if errIn != nil {
				return false
			}
			if inVal >= last {
				return false
			}
		}
		return true
	}
	for _, value := range split {
		valIn := 0
		for _, value := range string(value) {
			val := int(value) - 48
			if val < 0 || val > 9 {
				return "false"
			}
			valIn += val
		}
		if index == 0 {
			if valIn%2 != 0 {
				return "false"
			}
		} else if index == 1 {
			if valIn%2 == 0 {
				return "false"
			}
		}
		if !checkValIsLess(string(value)) {
			return "false"
		}
		index++
	}
	return "true"
}

func StringChanges(str string) string {
	if len(str) == 0 {
		return ""
	}
	deleteNext := false
	var text []string
	for u := 0; u < len(str); u++ {
		if string(str[u]) == "M" {
			deleteNext = false
			if len(text) != 0 {
				text = append(text, text[len(text)-1:][0])
			}
			continue
		} else if string(str[u]) == "N" {
			if u != 0 && string(str[u-1]) != "N" {
				deleteNext = true
			} else {
				deleteNext = false
			}
			continue
		} else {
			if !deleteNext {
				text = append(text, strings.ToLower(string(str[u])))
			}
			deleteNext = false
		}
	}
	return strings.Join(text, "")
}

func FizzBuzz(num int) string {
	var result []string
	for u := 1; u <= num; u++ {
		if u%3 == 0 && u%5 != 0 {
			result = append(result, "Fizz")
		} else if u%3 != 0 && u%5 == 0 {
			result = append(result, "Buzz")
		} else if u%3 == 0 && u%5 == 0 {
			result = append(result, "FizzBuzz")
		} else {
			result = append(result, fmt.Sprintf("%d", u))
		}
	}
	return strings.Join(result, " ")
}

//"0.38" -> half empty empty empty empty
//"4.5"  -> full full full full half
func StarRating(str string) string {
	half, empty, full, min, max, halfFactor, fullFactor := "half", "empty", "full", 0.0, 5.0, 0.5, 1.0
	generate := func(which string, count int) string {
		var res []string
		for u := 0; u < count; u++ {
			res = append(res, which)
		}
		return strings.Join(res, " ")
	}
	if len(str) == 0 {
		return generate(empty, int(max))
	}

	num64, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return generate(empty, int(max))
	}

	round := func(x, unit float64) float64 {
		return math.Round(x/unit) * unit
	}
	num64 = round(num64, halfFactor)

	if num64 < min && num64 > max {
		return generate(empty, int(max))
	}
	if num64 == min {
		return generate(empty, int(max))
	}
	sum := 0.0
	hCount, fCount, eCount := 0, 0, 0
	for sum <= num64 {
		if sum+fullFactor <= num64 {
			sum += fullFactor
			fCount++
			if sum == num64 {
				break
			}
			continue
		} else if sum <= num64 {
			sum += num64 - sum
			hCount++
			if sum == num64 {
				break
			}
			continue
		} else {
			break
		}
	}
	eCount = int(max - float64(fCount))
	if hCount != 0 {
		eCount -= 1
	}
	result := ""
	if fCount > 0 {
		result += generate(full, fCount)
	}
	if hCount > 0 {
		if len(result) != 0 {
			result += " "
		}
		result += generate(half, hCount)
	}
	if eCount > 0 {
		if len(result) != 0 {
			result += " "
		}
		result += generate(empty, eCount)
	}
	return result
}

func CommandLine(str string) string {
	if len(str) == 0 {
		return ""
	}
	var templates []string
	templates = strings.Split(str, "=")
	kv := make(map[int]map[string]string, 0)
	indexMe := 0
	reservedKey := ""
	for _, value := range templates {
		if indexMe == 0 {
			reservedKey = string(value)
		} else {
			findLastEmpty := strings.LastIndex(string(value), " ")
			kv[indexMe] = make(map[string]string, 1)
			val := kv[indexMe]
			if findLastEmpty != -1 && indexMe != len(templates)-1 {
				val[reservedKey] = string(value)[:findLastEmpty]
				reservedKey = string(value)[findLastEmpty+1:]
			} else {
				val[reservedKey] = string(value)[:]
			}
		}
		indexMe++
	}
	text := ""
	var keys []int
	for key := range templates {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, keyIndex := range keys {
		if data, isExist := kv[keyIndex]; isExist {
			for key, value := range data {
				text += fmt.Sprintf("%v=%v ", len(key), len(value))
			}
		}
	}
	if text[len(text)-1:] == " " {
		text = text[:len(text)-1]
	}
	return text
}

func PalindromeSwapperExtended(str string) string {
	if len(str) == 0 {
		return "-1"
	}
	var allItems []string
	for _, value := range str {
		allItems = append(allItems, string(value))
	}
	kv := make(map[string]int, 0)
	for _, value := range str {
		if val, isExist := kv[string(value)]; isExist {
			kv[string(value)] = val + 1
		} else {
			kv[string(value)] = 1
		}
	}
	threshold := 0
	if len(str)%2 != 0 {
		threshold++
	}
	detectedCount := 0
	for _, value := range kv {
		if value != 2 {
			detectedCount++
			if detectedCount > threshold {
				return "-1"
			}
		}
	}

	var permutation func([]string, int)
	result := make([][]string, 0)
	reverse := func(val string) string {
		if len(val) == 0 {
			return ""
		}
		var text []string
		for _, value := range val {
			text = append(text, string(value))
		}
		MAX := len(text) / 2
		for u := 0; u < MAX; u++ {
			temp := text[u]
			text[u] = text[len(text)-u-1]
			text[len(text)-u-1] = temp
		}
		return strings.Join(text, "")
	}
	palindromeRes := "-1"
	permutation = func(items []string, n int) {
		if n == 1 {
			arr := make([]string, len(items))
			copy(arr, items)
			result = append(result, arr)
			data := strings.Join(arr, "")
			if palindromeRes == "-1" && data == reverse(data) {
				palindromeRes = data
				return
			}
		} else {
			for i := 0; i < n; i++ {
				permutation(items, n-1)
				if n%2 == 1 {
					temp := items[i]
					items[i] = items[n-1]
					items[n-1] = temp
				} else {
					temp := items[0]
					items[0] = items[n-1]
					items[n-1] = temp
				}
			}
		}
	}
	permutation(allItems, len(allItems))

	return palindromeRes
}

func PalindromeSwapper(str string) string {
	if len(str) == 0 {
		return "-1"
	}
	var result = "-1"
	reverse := func(val string) string {
		if len(val) == 0 {
			return ""
		}
		var text []string
		for _, value := range val {
			text = append(text, string(value))
		}
		MAX := len(text) / 2
		for u := 0; u < MAX; u++ {
			temp := text[u]
			text[u] = text[len(text)-u-1]
			text[len(text)-u-1] = temp
		}
		return strings.Join(text, "")
	}
	var allItems []string
	for _, value := range str {
		allItems = append(allItems, string(value))
	}
	temp := make([]string, len(allItems))
	copy(temp, allItems)
	in := false
out:
	for u := 0; u < len(allItems); u++ {
		for y := u + 1; y < len(allItems); y++ {
			tempChar := temp[u]
			temp[u] = temp[y]
			temp[y] = tempChar
			data := strings.Join(temp, "")
			if data == reverse(data) {
				result = data
				in = true
				break out
			} else {
				temp = make([]string, len(allItems))
				copy(temp, allItems)
			}
			u++
		}
	}
	if !in {
		if str == reverse(str) {
			return result
		}
	}
	return result
}

func TriangleRow(num int) int {
	sum := 0
	if num < 0 {
		return sum
	}
	num++
	rowGenerator := func(row int) []int {
		var result []int
		for u := 0; u < row; u++ {
			if u == 0 {
				result = []int{1}
			} else if u == 1 {
				result = []int{1, 1}
			} else if u == 2 {
				result = []int{1, 2, 1}
			} else {
				var temp []int
			do:
				temp = append(temp, result[0])
				last := result[len(result)-1]
				for y := 0; y < len(result); y++ {
					for k := y + 1; k < len(result); k++ {
						temp = append(temp, result[y]+result[k])
						y++
					}
				}
				temp = append(temp, last)
				if len(temp) < num {
					result = make([]int, len(temp))
					copy(result, temp)
					temp = make([]int, 0)
					goto do
				} else {
					return temp
				}
			}
		}
		return result
	}
	for _, value := range rowGenerator(num) {
		sum += value
	}

	return sum
}

func TimeDifference(strArr []string) int {
	if len(strArr) < 2 {
		return 0
	}
	padLeft := func(base string, what string, count int) string {
		if len(base) >= count {
			return base
		}
		add := ""
		for u := len(base); u < count; u++ {
			add += what
		}
		return add + base
	}
	hourAsMinute := 60
	dayAsHour := 24
	suffixAm, suffixPm := "am", "pm"
	timeIndexGenerator := func() map[int]string {
		hour := 12
		minute := 0
		timeIndex := make(map[int]string, hourAsMinute*dayAsHour)
		for u := 0; u < hourAsMinute*dayAsHour; u++ {
			if u%hourAsMinute == 0 {
				minute = 0
				if u == 0 {
					hour = dayAsHour / 2
					timeIndex[u] = fmt.Sprintf("%v:%v%v", padLeft(strconv.Itoa(hour), "0", 2), padLeft(strconv.Itoa(minute), "0", 2), suffixAm)
				} else if u == hourAsMinute {
					hour = 1
					timeIndex[u] = fmt.Sprintf("%v:%v%v", padLeft(strconv.Itoa(hour), "0", 2), padLeft(strconv.Itoa(minute), "0", 2), suffixAm)
				} else {
					hour++
					suffix := suffixAm
					if u >= hourAsMinute*(dayAsHour/2) {
						suffix = suffixPm
					}
					if u == hourAsMinute*(dayAsHour/2) {
						hour = dayAsHour / 2
					} else if u == (hourAsMinute*(dayAsHour/2))+hourAsMinute {
						hour = 1
					}
					timeIndex[u] = fmt.Sprintf("%v:%v%v", padLeft(strconv.Itoa(hour), "0", 2), padLeft(strconv.Itoa(minute), "0", 2), suffix)
				}
			} else {
				minute++
				suffix := suffixAm
				if u >= hourAsMinute*(dayAsHour/2) {
					suffix = suffixPm
				}
				timeIndex[u] = fmt.Sprintf("%v:%v%v", padLeft(strconv.Itoa(hour), "0", 2), padLeft(strconv.Itoa(minute), "0", 2), suffix)
			}
		}
		return timeIndex
	}

	res := timeIndexGenerator()
	var keys []int
	for key := range res {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	timeIndexFind := func(val string) int {
		for key, valueOfIndex := range keys {
			if res[valueOfIndex] == val {
				return key
			}
		}
		return 0
	}
	timeMinuteDiffCalc := func(time []string) int {
		if len(time) == 0 {
			return 0
		}
		min := hourAsMinute * dayAsHour
		for u := 0; u < len(time); u++ {
			for y := 0; y < len(time); y++ {
				if y != u {
					suff1Pm := strings.Contains(time[u], suffixPm)
					suff2Am := strings.Contains(time[y], suffixAm)
					t1 := timeIndexFind(padLeft(time[u], "0", 7))
					t2 := timeIndexFind(padLeft(time[y], "0", 7))
					op := 0
					diffTwo := 0
					if suff1Pm && suff2Am {
						op = hourAsMinute*(dayAsHour) - t1
						t1 = 0
						diffTwo = int(math.Abs(float64(t1)+float64(t2))) + op
					} else {
						diffTwo = int(math.Abs(float64(t1)-float64(t2))) + op
					}
					if diffTwo < min {
						min = diffTwo
					}
				}
			}
		}
		return min
	}
	return timeMinuteDiffCalc(strArr)
}

func MovingMedian(arr []int) string {
	if len(arr) == 0 {
		return "0"
	}
	medianFinder := func(array []int) (result int) {
		result = 0
		sort.Slice(array, func(i, j int) bool {
			return array[i] < array[j]
		})
		if len(array)%2 != 0 {
			result = array[len(array)/2]
		} else {
			result = (array[len(array)/2-1] + array[len(array)/2]) / 2
		}
		return
	}

	var result []string
	if len(arr) < 2 {
		return strconv.Itoa(arr[0])
	}
	first := arr[0]
	arr = arr[1:]
	for u := 0; u < len(arr); u++ {
		var subset []int
		for j := u; j >= 0 && j > (u-first); j-- {
			if j < len(arr) {
				subset = append(subset, arr[j])
			}
		}
		result = append(result, fmt.Sprintf("%v", medianFinder(subset)))
	}
	return strings.Join(result, ",")
}

func MoneyDistribution(money float64, maxDistributionCount int) (float64, map[int]float64, error) {
	if money <= 0 {
		return -1, nil, errors.New("There is no money to share!")
	}
	if maxDistributionCount <= 0 {
		return -1, nil, errors.New("There is no one to share the money!")
	}
	round := func(x, unit float64) float64 {
		return math.Round(x/unit) * unit
	}
	_ = round
	max := 0.0
	percent := 1.0
	people := make(map[int]float64, maxDistributionCount)
	for u := 1; u <= maxDistributionCount; u++ {
		allowance := money * float64(u) / float64(maxDistributionCount)
		people[u] = allowance
		if allowance > max {
			max = allowance
		}
		money -= allowance
		percent += 1.0
	}
	return max, people, nil
}

// "[3, 4]", "[1, 2, 7, 7]
func ScaleBalancing(strArr []string) string {
	notPossible := "not possible"
	if len(strArr) != 2 {
		return notPossible
	}
	index := 0
	var pool []int
	var left int
	var right int
	var maxSelectionCount = 2
	for _, value := range strArr {
		value = strings.Replace(value, "[", "", -1)
		value = strings.Replace(value, "]", "", -1)
		values := strings.Split(value, ",")
		if index == 0 {
			if len(values) != 2 {
				return notPossible
			}
			var errLeft error
			var errRight error
			left, errLeft = strconv.Atoi(strings.TrimSpace(values[0]))
			if errLeft != nil {
				return notPossible
			}
			right, errRight = strconv.Atoi(strings.TrimSpace(values[1]))
			if errRight != nil {
				return notPossible
			}
		} else {
			for _, value := range values {
				val, errInt := strconv.Atoi(strings.TrimSpace(value))
				if errInt != nil {
					return notPossible
				}
				pool = append(pool, val)
			}
		}
		index++
	}
	binarySearch := func(items []int, search int) (isFound bool, itemIndex int) {
		itemIndex = -1
		isFound = false
		left := 0
		right := len(items) - 1
		for left <= right {
			middle := (left + right) / 2
			if items[middle] == search {
				itemIndex = middle
				isFound = true
				break
			} else if items[middle] > search {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		return
	}
	var limitedPowerSet func([]int, int, []int, *[][]int, int)
	limitedPowerSet = func(subset []int, breakPoint int, selectedFar []int, result *[][]int, maxSliceLen int) {
		if len(subset) == breakPoint {
			if len(selectedFar) == maxSliceLen {
				clone := make([]int, len(selectedFar))
				copy(clone, selectedFar)
				*result = append(*result, clone)
			}
			return
		} else {
			selectedFar = append(selectedFar, subset[breakPoint])
			limitedPowerSet(subset, breakPoint+1, selectedFar, result, maxSliceLen)
			selectedFar = selectedFar[:len(selectedFar)-1]
			limitedPowerSet(subset, breakPoint+1, selectedFar, result, maxSliceLen)
		}
	}
	index = 0
	sort.Slice(pool, func(i, j int) bool {
		return pool[i] < pool[j]
	})
	var selected []string
	selectedSoFar := make([]int, 0)
	results := make([][]int, 0)
out:
	for {
		if left == right {
			break
		} else if left > right {
			if index == 0 {
				diff := left - right
				isFound, itemIndex := binarySearch(pool, diff)
				if isFound {
					selected = append(selected, strconv.Itoa(pool[itemIndex]))
					break
				}
			} else if index == 1 {
				limitedPowerSet(pool, 0, selectedSoFar, &results, maxSelectionCount)
				for _, value := range results {
					sum := 0
					for _, inValue := range value {
						sum += inValue
					}
					if sum+right == left {
						for _, inValue := range value {
							selected = append(selected, strconv.Itoa(inValue))
						}
						break out
					}
				}
			}
		} else if right > left {
			if index == 0 {
				diff := right - left
				isFound, itemIndex := binarySearch(pool, diff)
				if isFound {
					selected = append(selected, strconv.Itoa(pool[itemIndex]))
					break
				}
			} else if index == 1 {
				limitedPowerSet(pool, 0, selectedSoFar, &results, maxSelectionCount)
				for _, value := range results {
					sum := 0
					for _, inValue := range value {
						sum += inValue
					}
					if sum+left == right {
						for _, inValue := range value {
							selected = append(selected, strconv.Itoa(inValue))
						}
						break out
					}
				}
			}
		}
		if index == 2 {
			var leftSideItems, rightSideItems []int
			for _, value := range pool {
				leftSideItems = append(leftSideItems, left+value)
			}
			for _, value := range pool {
				rightSideItems = append(rightSideItems, right+value)
			}
			for u := 0; u < len(leftSideItems); u++ {
				for y := 0; y < len(rightSideItems); y++ {
					if leftSideItems[u] == rightSideItems[y] {
						selected = append(selected, strconv.Itoa(pool[u]), strconv.Itoa(pool[y]))
						break out
					}
				}
			}
		} else {
			if index > 2 {
				break out
			}
		}
		index++
	}
	if len(selected) != 0 {
		sort.Slice(selected, func(i, j int) bool {
			val1, _ := strconv.Atoi(selected[i])
			val2, _ := strconv.Atoi(selected[j])
			return val1 < val2
		})
		return strings.Join(selected, ",")
	}
	return notPossible
}

func PalindromeCreator(str string) string {
	notPossible := "not possible"
	if len(str) == 0 {
		return notPossible
	}
	minLength := 3
	reverse := func(text string) string {
		if len(text) == 0 {
			return ""
		}
		var allInside []string
		for _, value := range text {
			allInside = append(allInside, string(value))
		}
		MAX := len(allInside) / 2
		for u := 0; u < MAX; u++ {
			temp := allInside[u]
			allInside[u] = allInside[len(allInside)-u-1]
			allInside[len(allInside)-u-1] = temp
		}
		return strings.Join(allInside, "")
	}
	if len(str) >= minLength && reverse(str) == str {
		return "palindrome"
	}
	var all []string
	for _, value := range str {
		all = append(all, string(value))
	}

	chars := make([]string, 0)
	index := 0
	isDone := false
do:
	for u := 0; u < len(all); u++ {
		if index == 0 {
			newText := strings.Join(all, "")
			if u > 0 {
				newText = strings.Join(all[:u], "") + strings.Join(all[u+1:], "")
			} else {
				newText = strings.Join(all[1:], "")
			}
			if len(newText) >= minLength && newText == reverse(newText) {
				chars = append(chars, all[u])
				isDone = true
				break do
			}
			if u == len(all)-1 {
				index++
			}
		} else if index == 1 {
			for y := u + 1; y < len(all); y++ {
				if u != y {
					newText := ""
					for j := 0; j < len(all); j++ {
						if j != u && j != y {
							newText += all[j]
						}
					}
					if len(newText) >= minLength && newText == reverse(newText) {
						chars = append(chars, all[u], all[y])
						isDone = true
						break do
					}
				}
			}
			if u == len(all)-1 {
				index++
			}
		}
	}
	if !isDone {
		if index < 2 {
			goto do
		}
		return notPossible
	} else {
		if len(chars) != 0 {
			return strings.Join(chars, "")
		}
		return notPossible
	}
}

func ClosestEnemyII_old(strArr []string) int {
	if len(strArr) == 0 {
		return 0
	}
	lengthOfOne := len(strArr[0])
	for u := 1; u < len(strArr); u++ {
		if len(strArr[u]) != lengthOfOne {
			return 0
		}
	}
	type Coordinate struct {
		X, Y int
	}
	transposition := func(rawMatrix []string) (allMatrix [][]string, firsts []Coordinate, seconds []Coordinate) {
		allMatrix = make([][]string, len(rawMatrix[0]))
		firsts = make([]Coordinate, 0)
		seconds = make([]Coordinate, 0)
		if len(rawMatrix) == 0 {
			return
		}
		for u := 0; u < len(rawMatrix[0]); u++ {
			newLine := make([]string, len(rawMatrix))
			for y := 0; y < len(rawMatrix); y++ {
				val := string(rawMatrix[y][u])
				if val == "1" {
					firsts = append(firsts, Coordinate{X: u, Y: y})
				} else if val == "2" {
					seconds = append(seconds, Coordinate{X: u, Y: y})
				}
				newLine[y] = val
			}
			allMatrix[u] = newLine
		}
		return
	}

	result, firsts, seconds := transposition(strArr)
	if len(result) != 0 {
		if len(firsts) != 1 {
			return 0
		}
		if len(seconds) == 0 {
			return 0
		}
		var distance []int
		action := 0
		sort.Slice(seconds, func(i, j int) bool {
			return seconds[i].X < seconds[j].X
		})
		for _, second := range seconds {
			if second.X != 0 && second.X != len(result)-1 && firsts[0].X != 0 && firsts[0].X != len(result)-1 {
				action = int(math.Abs(float64(firsts[0].X) - float64(second.X)))
			}
			if firsts[0].Y != second.Y {
				actionFirst := int(math.Abs(float64(firsts[0].Y) - float64(second.Y)))
				actionSecondOne := firsts[0].Y
				actionSecondTwo := len(result) - second.Y - 1
				actionFinal := actionSecondOne + actionSecondTwo
				if actionFirst == len(result)-1 && actionFinal == 0 {
					action = 1
				}
				if actionFinal <= actionFirst {
					action += actionFinal
				} else {
					action += actionFirst
				}
			}
			distance = append(distance, action)
			break
		}

		if len(distance) != 0 {
			sort.Slice(distance, func(i, j int) bool {
				return distance[i] < distance[j]
			})
			return distance[0]
		}
		return 0
	}
	return 0
}

func ClosestEnemyII(strArr []string) int {
	result := 1000
	pos1X, pos1Y := -1, -1
	distanceCalc := func(pos1X, pos1Y, pos2X, pos2Y, height, width int) int {
		movesX, movesX2, movesY, movesY2 := 0, 0, 0, 0
		if pos1X != pos2X {
			movesX = pos1X - pos2X
			if movesX < 0 {
				movesX = movesX * -1
			}
			movesX2 = pos1X + (width - pos2X)
			if movesX > movesX2 {
				movesX = movesX2
			}
		}
		if pos1Y != pos2Y {
			movesY = pos1Y - pos2Y
			if movesY < 0 {
				movesY = movesY * -1
			}
			movesY2 = pos1Y + (height - pos2Y)
			if movesY > movesY2 {
				movesY = movesY2
			}
		}
		return movesX + movesY
	}
	var pos2 = make(map[int][]int)
	for index, value := range strArr {
		if pos1X < 0 {
			index1 := strings.Index(string(value), "1")
			if index1 >= 0 {
				pos1X = index1
				pos1Y = index
			}
		}
		for i, v := range string(value) {
			if string(v) == "2" {
				pos2[i] = append(pos2[i], index)
			}
		}
	}
	if len(pos2) < 1 {
		return 0
	}
	for index, value := range pos2 {
		for _, v := range value {
			moves := distanceCalc(pos1X, pos1Y, index, v, len(strArr), len(strArr[0]))
			if moves < result {
				result = moves
			}
		}
	}
	return result
}

func VowelSquare(strArr []string) string {
	notFound := "not found"
	if len(strArr) < 2 {
		return notFound
	}
	lenBase := len(strArr[0])
	for u := 1; u < len(strArr); u++ {
		if len(strArr[u]) != lenBase {
			return notFound
		}
	}
	var vowels = []rune{'a', 'e', 'i', 'o', 'u'}
	isVowel := func(val rune) bool {
		for u := 0; u < len(vowels); u++ {
			if vowels[u] == val {
				return true
			}
		}
		return false
	}
	indexer := make([]string, 0)
	kv := make(map[string][]int, 0)
	for u := 0; u < len(strArr); u++ {
		blockSize := 0
		startIndex := -1
		lineCount := 0
		for y := 0; y < len(strArr[u]); y++ {
			if isVowel(rune(strArr[u][y])) {
				if startIndex == -1 {
					startIndex = y
				}
				blockSize++
			} else {
				if startIndex != -1 {
					if blockSize >= 2 {
						text := fmt.Sprintf("%v#%v", u, lineCount)
						kv[text] = []int{startIndex, blockSize}
						indexer = append(indexer, text)
						lineCount++
					}
					startIndex = -1
					blockSize = 0
				} else {
					continue
				}
			}
		}
		if blockSize >= 2 {
			text := fmt.Sprintf("%v#%v", u, lineCount)
			kv[text] = []int{startIndex, blockSize}
			indexer = append(indexer, text)
		}
	}
	beforePoint := -1
	before, beforeMe := -1, -1
	lineCheck := 0
	result := ""
	if len(indexer) < 2 {
		return notFound
	}
	lock := false
	kvCheck := make(map[int][]int, 0)
	for key, value := range indexer {
		if points, isOk := kv[value]; isOk {
			splitter := strings.Split(value, "#")
			if len(splitter) == 2 {
				if before == -1 {
					var err error
					before, err = strconv.Atoi(splitter[0])
					if err != nil {
						return notFound
					}
					beforeMe = before
					lineCheck++
					beforePoint = points[0]
					result = fmt.Sprintf("%v-%v", splitter[0], points[0])
					var add []int
					for n := points[0]; n < points[0]+points[1]; n++ {
						add = append(add, n)
					}
					kvCheck[key] = add
				} else {
					newVal, err := strconv.Atoi(splitter[0])
					if err != nil {
						return notFound
					}
					if newVal-beforeMe != 1 {
						lineCheck = 0
					} else {
						lineCheck++
					}
					if points[0] > beforePoint && !lock {
						result = fmt.Sprintf("%v-%v", before, points[0])
					} else {
						lock = true
					}
					beforeMe = newVal
					var add []int
					for n := points[0]; n < points[0]+points[1]; n++ {
						add = append(add, n)
					}
					kvCheck[key] = add
				}
			}
		}
	}
	if lineCheck < 2 {
		return notFound
	}

	cpy := make([]int, 0)
	index := 0
	isOk := false
outx:
	for _, value := range kvCheck {
		if index == 0 {
			cpy = make([]int, len(value))
			copy(cpy, value)
		} else {
			matchCount := 0
			if len(cpy) > len(value) {
				for _, out := range cpy {
					for _, in := range value {
						if out == in {
							matchCount++
							if matchCount == 2 {
								isOk = true
								break outx
							}
							break
						}
					}
				}
			} else if len(cpy) <= len(value) {
				for _, out := range value {
					for _, in := range cpy {
						if out == in {
							matchCount++
							if matchCount == 2 {
								isOk = true
								break outx
							}
							break
						}
					}
				}
			}

		}
		index++
	}
	if !isOk {
		return notFound
	}
	return result
}

func TestX(all []string, maxLength int) {
	var permx func([]string, int, []string, *[][]string, int)
	permx = func(subset []string, breakPoint int, selectedSoFar []string, results *[][]string, maxLen int) {
		if len(subset) == breakPoint {
			if maxLen == len(selectedSoFar) {
				res := make([]string, len(selectedSoFar))
				copy(res, selectedSoFar)
				*results = append(*results, res)
			}
			return
		}
		selectedSoFar = append(selectedSoFar, subset[breakPoint])
		permx(subset, breakPoint+1, selectedSoFar, results, maxLength)
		selectedSoFar = selectedSoFar[:len(selectedSoFar)-1]
		permx(subset, breakPoint+1, selectedSoFar, results, maxLength)
	}

	selectedSF := make([]string, 0)
	resultAll := make([][]string, 0)
	permx(all, 0, selectedSF, &resultAll, 2)
	for _, value := range resultAll {
		fmt.Printf("%v\n", value)
	}

}

func withTimeOut() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	sleepAndTalk(ctx, 5*time.Second, "Hello withTimeOut!")
}

func withCancel() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	sleepAndTalk(ctx, 5*time.Second, "Hello withCancel!")
}

func background() {
	ctx := context.Background()
	sleepAndTalk(ctx, 5*time.Second, "Hello background!")
}

func sleepAndTalk(ctx context.Context, d time.Duration, s string) {
	select {
	case <-time.After(d):
		fmt.Println(s)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}

func Calc(values ...int) int {
	base := 1
	for _, value := range values {
		if value != 0 {
			base *= value
		}
	}
	return base
}

// N üzeri P  - N  % P == 1

func PrimeTime(num int) string {
	if num < 2 {
		return "false"
	}
	if num == 2 || num == 3 {
		return "true"
	}
	A := num - 1
	for A > 2 {
		calc := int64(math.Pow(float64(A), float64(num))-float64(A)) % int64(num)
		//fmt.Printf("%v\n", calc)
		if calc == 0 {
			return "true"
		}
		A--
	}
	return "false"
}

func PrimeNew(num int) string {
	for i := 2; i <= int(math.Floor(float64(num)/2)); i++ {
		if num%i == 0 {
			return "false"
		}
	}
	if num < 1 {
		return "false"
	}
	return "true"
}

func RunLength(str string) string {
	if len(str) == 0 {
		return ""
	}
	result := ""
	first := string(str[0])
	index := 1
	inside := false
	if len(str) == 1 {
		return fmt.Sprintf("%d%v", len(str), string(str[0]))
	}
	for u := 1; u < len(str); u++ {
		if string(str[u]) == first {
			index++
			inside = true
		} else {
			result += fmt.Sprintf("%d%v", index, string(str[u-1]))
			index = 1
			first = string(str[u])
			inside = false
		}
	}
	if (!inside && string(result[len(result)-1:]) != first) || (index > 1 && inside) {
		result += fmt.Sprintf("%d%v", index, string(str[len(str)-1]))
	}
	return result
}

func PalindromeTwo(str string) string {
	var punchItems = []string{
		" ",
		"-",
		"%",
		"’'",
		"()[]{}<>",
		":",
		",",
		"+",
		"‒–—―",
		"…",
		"!",
		".",
		"«»",
		"-‐",
		"?",
		"‘’“”",
		";",
		"/",
		"⁄",
		"␠",
		"·",
		"&",
		"@",
		"*",
		"\\",
		"•",
		"^",
		"¤¢$€£¥₩₪",
		"†‡",
		"°",
		"¡",
		"¿",
		"¬",
		"#",
		"№",
		"%‰‱",
		"¶",
		"′",
		"§",
		"~",
		"¨",
		"_",
		"|¦",
		"⁂",
		"☞",
		"∴",
		"‽",
		"※"}
	text := ""
	for _, value := range str {
		isFound := false
		for _, char := range punchItems {
			if string(value) == string(char) {
				isFound = true
				break
			}
		}
		if !isFound {
			text += strings.ToLower(string(value))
		}
	}
	reverse := func(val string) string {
		if len(val) == 0 {
			return ""
		}
		result := make([]string, len(val))
		for k := 0; k < len(val); k++ {
			result[k] = string(val[k])
		}
		max := len(val) / 2
		for u := 0; u < max; u++ {
			temp := result[u]
			result[u] = result[len(result)-u-1]
			result[len(result)-u-1] = temp
		}
		return strings.Join(result, "")
	}
	if text == reverse(text) {
		return "true"
	}
	return "false"
}

func Division(num1 int, num2 int) int {
	max := num1
	if num2 > max {
		max = num2
	}
	result := 1
	for u := 1; u <= max; u++ {
		if num1%u == 0 && num2%u == 0 {
			result = u
		}
	}
	return result
}

func StringScramble(str1 string, str2 string) string {
	if len(str1) == 0 && len(str2) == 0 {
		return "false"
	}

	var findIndexList []int
	findMe := func(val int) bool {
		for _, value := range findIndexList {
			if val == value {
				return true
			}
		}
		return false
	}
	for _, value := range str2 {
		for index, char := range str1 {
			if string(char) == string(value) && !findMe(index) {
				findIndexList = append(findIndexList, index)
				break
			}
		}
	}
	if len(str2) == len(findIndexList) {
		return "true"
	}

	return "false"
}

func ArithGeoII(arr []int) string {
	if len(arr) < 2 {
		return "-1"
	}
	first := arr[0]
	returnMe := true
	for _, value := range arr {
		if value == 0 {
			return "-1"
		}
		if value != first {
			returnMe = false
		}
	}
	if returnMe {
		return "-1"
	}
	diffSub := math.Abs(float64(arr[0]) - float64(arr[1]))
	diffMul := math.Abs(float64(arr[1]) / float64(arr[0]))
	isArithmetic, isGeoMetric := true, true
	for u := 1; u < len(arr); u++ {
		if isArithmetic && math.Abs(float64(arr[u])-float64(arr[u-1])) != diffSub {
			isArithmetic = false
		}
		if isGeoMetric && math.Abs(float64(arr[u])/float64(arr[u-1])) != diffMul {
			isGeoMetric = false
		}
		if !isGeoMetric && !isArithmetic {
			return "-1"
		}
	}
	if isArithmetic {
		return "Arithmetic"
	}
	if isGeoMetric {
		return "Geometric"
	}
	return "-1"
}

func LetterCount(str string) string {
	if len(str) == 0 {
		return "-1"
	}
	words := strings.Split(str, " ")
	kvOut := make(map[int]int, len(words))
	breakDown := true
	for key, value := range words {
		kvIn := make(map[string]int, 0)
		for u := 0; u < len(value); u++ {
			if val, isExist := kvIn[string(value[u])]; isExist {
				kvIn[string(value[u])] = val + 1
			} else {
				kvIn[string(value[u])] = 1
			}
		}
		max := 0
		for _, value := range kvIn {
			if value > max {
				max = value
			}
		}
		kvOut[key] = max
		if breakDown && max != 1 {
			breakDown = false
		}
	}
	if breakDown {
		return "-1"
	}
	var indexSort []int
	for key, _ := range kvOut {
		indexSort = append(indexSort, key)
	}
	sort.Slice(indexSort, func(i, j int) bool {
		return indexSort[i] < indexSort[j]
	})
	max := 0
	keyIndex := 0
	for key := range indexSort {
		if kvOut[key] > max {
			max = kvOut[key]
			keyIndex = key
		}
	}
	maxCount := 0
	for _, value := range kvOut {
		if value == max {
			maxCount++
		}
	}
	if maxCount == 1 {
		return words[keyIndex]
	} else {
		for key := range indexSort {
			if kvOut[key] == max {
				max = kvOut[key]
				keyIndex = key
				break
			}
		}
		return words[keyIndex]
	}
}

func DashInsertIIRev(num string) string {
	if len(num) == 0 {
		return ""
	}
	result := ""
	for u := 1; u < len(num); u++ {
		setVal := func(char string) {
			result = fmt.Sprintf("%v%v%v", result, char, string(num[u]))
		}
		if string(num[u]) != "0" {
			numVal, _ := strconv.Atoi(string(num[u]))
			numValBef, _ := strconv.Atoi(string(num[u-1]))
			if numValBef != 0 {
				if numVal%2 == 0 {
					if numValBef%2 == 0 {
						setVal("*")
					} else {
						setVal("")
					}
				} else {
					if numValBef%2 != 0 {
						setVal("-")
					} else {
						setVal("")
					}
				}
			} else {
				setVal("")
			}
		} else {
			setVal("")
		}
	}
	return string(num[0]) + result
}

func DashInsertII(num int) string {
	numbers := fmt.Sprintf("%v", num)
	if len(numbers) == 0 {
		return ""
	}
	result := ""
	for u := 1; u < len(numbers); u++ {
		add := func(char string) {
			result = fmt.Sprintf("%v%v%v", result, char, string(numbers[u]))
		}
		if string(numbers[u]) != "0" {
			numVal, _ := strconv.Atoi(string(numbers[u]))
			numValBef, _ := strconv.Atoi(string(numbers[u-1]))
			if numValBef != 0 {
				if numVal%2 == 0 {
					if numValBef%2 == 0 {
						add("*")
					} else {
						add("")
					}
				} else {
					if numValBef%2 != 0 {
						add("-")
					} else {
						add("")
					}
				}
			} else {
				add("")
			}
		} else {
			add("")
		}
	}
	return string(numbers[0]) + result
}

func ArrayAddition(arr []int) string {
	if len(arr) == 0 {
		return "false"
	}
	first := arr[0]
	isOk := false
	for u := 1; u < len(arr); u++ {
		if arr[u] != first {
			isOk = true
		}
	}
	if !isOk {
		return "false"
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	max := arr[0]
	arr = arr[1:]
	var powerset func([]int, int, []int, *[][]int)
	powerset = func(subset []int, breakPoint int, selectedSoFar []int, result *[][]int) {
		if len(subset) == breakPoint {
			newArr := make([]int, len(selectedSoFar))
			copy(newArr, selectedSoFar)
			*result = append(*result, newArr)
			return
		} else {
			selectedSoFar = append(selectedSoFar, subset[breakPoint])
			powerset(subset, breakPoint+1, selectedSoFar, result)
			selectedSoFar = selectedSoFar[:len(selectedSoFar)-1]
			powerset(subset, breakPoint+1, selectedSoFar, result)
		}
	}

	resultSet := make([][]int, 0)
	selectedSoFar := make([]int, 0)
	powerset(arr, 0, selectedSoFar, &resultSet)
	for _, value := range resultSet {
		sum := 0
		for _, innerValue := range value {
			sum += innerValue
		}
		if sum == max {
			return "true"
		}
	}
	return "false"
}

func BinaryConverter(str string) string {
	if len(str) == 0 {
		return "0"
	}
	sum := 0.0
	for u := len(str) - 1; u >= 0; u-- {
		val, _ := strconv.Atoi(string(str[len(str)-u-1]))
		if val == 1 {
			sum += math.Pow(float64(2), float64(u))
		}
		fmt.Printf("%v-%v\n", u, val)
	}
	return fmt.Sprintf("%v", sum)
}

func CaesarCipher(str string, num int) string {
	if len(str) == 0 {
		return ""
	}
	if num <= 0 {
		return str
	}
	var text []string
	for _, value := range str {
		text = append(text, string(value))
	}
	for u := 0; u < len(text); u++ {
		charIndex := []rune(text[u])[0]
		if charIndex >= 65 && charIndex <= 122 {
			valx := int(charIndex) + num
			diff := 0
			if valx > 90 && valx < 97 {
				diff = valx - 91
				valx = 65
			}
			if valx > 122 {
				diff = valx - 123
				valx = 97
			}
			text[u] = fmt.Sprintf("%c", valx+diff)
		}
	}
	return strings.Join(text, "")
}

func SimpleMode(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	kv := make(map[int]int, 0)
	max := -1

	for u := 0; u < len(arr); u++ {
		if val, isOk := kv[arr[u]]; isOk {
			kv[arr[u]] = val + 1
			if val+1 > max {
				max = val + 1
			}
		} else {
			kv[arr[u]] = 1
		}
	}

	for _, value := range arr {
		if kv[value] == max {
			return value
		}
	}
	return -1
}

func Consecutive(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var numbers []int
	for u := 0; u < len(arr); u++ {
		numbers = append(numbers, arr[u])
	}

	index := 0
	for u := arr[0]; u <= arr[len(arr)-1]; u++ {
		isFound := false
		for j := 0; j < len(numbers); j++ {
			if numbers[j] == u {
				isFound = true
				break
			}
		}
		if !isFound {
			index++
		}
	}
	return index
}

func NumberSearch(str string) string {
	if len(str) == 0 {
		return "0"
	}
	roundme := func(x, unit float64) float64 {
		return math.Round(x/unit) * unit
	}
	sum, letterCount := 0, 0
	words := strings.Split(str, " ")
	for u := 0; u < len(words); u++ {
		for j := 0; j < len(words[u]); j++ {
			if rune(words[u][j]) >= 48 && rune(words[u][j]) <= 57 {
				sum += int(rune(words[u][j]) - 48)
			} else if (rune(words[u][j]) >= 65 && rune(words[u][j]) <= 89) ||
				(rune(words[u][j]) >= 97 && rune(words[u][j]) <= 122) {
				letterCount++
			}
		}
	}
	if letterCount == 0 {
		return "0"
	}
	return fmt.Sprintf("%0.f", roundme(float64(sum)/float64(letterCount), 1))
}

func ThreeFiveMultiples(num int) int {
	if num == 0 {
		return num
	}
	sum := 0
	for u := 0; u < num; u++ {
		if u%3 == 0 || u%5 == 0 {
			sum += u
		}
	}
	return sum
}

func StringReductionOld(str string) int {
	if len(str) == 0 {
		return 0
	}
	letters := []rune{'a', 'b', 'c'}
	for _, value := range str {
		hasFound := false
		for _, letter := range letters {
			if value == letter {
				hasFound = true
				break
			}
		}
		if !hasFound {
			return 0
		}
	}
	check := func(where string, what []string) bool {
		for _, value := range what {
			if strings.Index(where, value) != -1 {
				return true
			}
		}
		return false
	}
	reverse := func(data string) string {
		if len(data) == 0 {
			return ""
		}
		dataBody := make([]string, len(data))
		for key, value := range data {
			dataBody[key] = string(value)
		}
		MAX := len(dataBody) / 2
		for u := 0; u < MAX; u++ {
			temp := dataBody[u]
			dataBody[u] = dataBody[len(dataBody)-u-1]
			dataBody[len(dataBody)-u-1] = temp
		}
		return strings.Join(dataBody, "")
	}
	bodyGenerator := func(letters []rune) []string {
		if len(letters) == 0 {
			return []string{}
		}
		var result []string
		for _, letter := range letters {
			replaceBody := ""
			for _, value := range letters {
				if string(letter) != string(value) {
					replaceBody += string(value)
				}
			}
			result = append(result, replaceBody, reverse(replaceBody))
		}
		return result
	}
	replacer := func(what, text string) string {
		if len(what) != 1 {
			return ""
		}
		replaceBody := ""
		for _, value := range letters {
			if what != string(value) {
				replaceBody += string(value)
			}
		}
		text = strings.Replace(text, replaceBody, what, 1)
		text = strings.Replace(text, reverse(replaceBody), what, 1)
		return text
	}
	text := make([]string, len(str))
	for key, value := range str {
		text[key] = string(value)
	}
	lBlockMaxCount := func(letter string) int {
		max, count := 0, 0
		for u := 1; u < len(str); u++ {
			if string(str[u]) == letter {
				if string(str[u-1]) == letter {
					if u == 0 {
						count = 1
					}
					count++
				} else {
					count = 1
				}
				if count > max {
					max = count
				}
			}
		}
		return max
	}
	var letterList []string
	var values []int
	a := lBlockMaxCount("a")
	b := lBlockMaxCount("b")
	c := lBlockMaxCount("c")
	values = append(values, a, b, c)
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	if a == values[0] {
		letterList = append(letterList, "a", "b", "c")
	} else if b == values[0] {
		letterList = append(letterList, "b", "a", "c")
	} else {
		letterList = append(letterList, "c", "a", "b")
	}
do:
	for _, letter := range letterList {
		str = replacer(letter, str)
	}
	if check(str, bodyGenerator(letters)) {
		goto do
	}
	return len(str)
}

func StringReduction(str string) int {
	if len(str) == 0 {
		return 0
	}
	letters := []rune{'a', 'b', 'c'}
	for _, value := range str {
		hasFound := false
		for _, letter := range letters {
			if value == letter {
				hasFound = true
				break
			}
		}
		if !hasFound {
			return 0
		}
	}
	check := func(where string, what []string) bool {
		for _, value := range what {
			if strings.Index(where, value) != -1 {
				return true
			}
		}
		return false
	}
	reverse := func(data string) string {
		if len(data) == 0 {
			return ""
		}
		dataBody := make([]string, len(data))
		for key, value := range data {
			dataBody[key] = string(value)
		}
		MAX := len(dataBody) / 2
		for u := 0; u < MAX; u++ {
			temp := dataBody[u]
			dataBody[u] = dataBody[len(dataBody)-u-1]
			dataBody[len(dataBody)-u-1] = temp
		}
		return strings.Join(dataBody, "")
	}
	bodyGenerator := func(letters []rune) []string {
		if len(letters) == 0 {
			return []string{}
		}
		var result []string
		for _, letter := range letters {
			replaceBody := ""
			for _, value := range letters {
				if string(letter) != string(value) {
					replaceBody += string(value)
				}
			}
			result = append(result, replaceBody, reverse(replaceBody))
		}
		return result
	}
	replacer := func(what, text string) string {
		if len(what) != 1 {
			return ""
		}
		replaceBody := ""
		for _, value := range letters {
			if what != string(value) {
				replaceBody += string(value)
			}
		}
		text = strings.Replace(text, replaceBody, what, 1)
		text = strings.Replace(text, reverse(replaceBody), what, 1)
		return text
	}
	text := make([]string, len(str))
	for key, value := range str {
		text[key] = string(value)
	}

	var helper func([]rune, int)
	var resultSet [][]rune

	helper = func(arr []rune, n int) {
		if n == 1 {
			tmp := make([]rune, len(arr))
			copy(tmp, arr)
			resultSet = append(resultSet, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(letters, len(letters))

	var lenList []int
	index := 0
	strBackup := str
index:
	for _, letter := range resultSet[index] {
		str = replacer(string(letter), str)
	}
	if check(str, bodyGenerator(letters)) {
		goto index
	}
	lenList = append(lenList, len(str))
	if index < len(letters) {
		str = strBackup
		index++
		goto index
	}
	sort.Slice(lenList, func(i, j int) bool {
		return lenList[i] < lenList[j]

	})
	return lenList[0]
}

func PermutationStep(num int) int {
	var helper func([]int, int)
	var results [][]int
	helper = func(items []int, n int) {
		if n == 1 {
			temp := make([]int, len(items))
			copy(temp, items)
			results = append(results, temp)
		} else {
			for i := 0; i < n; i++ {
				helper(items, n-1)
				if n%2 == 1 {
					temp := items[i]
					items[i] = items[n-1]
					items[n-1] = temp
				} else {
					temp := items[0]
					items[0] = items[n-1]
					items[n-1] = temp
				}
			}
		}
	}

	var numLetters []int
	text := strconv.Itoa(num)
	for _, value := range text {
		val, _ := strconv.Atoi(string(value))
		numLetters = append(numLetters, val)
	}
	helper(numLetters, len(numLetters))

	var itemsAll []int
	for _, value := range results {
		number := ""
		for _, val := range value {
			number += strconv.Itoa(val)
		}
		digits, _ := strconv.Atoi(number)
		itemsAll = append(itemsAll, digits)
	}
	sort.Slice(itemsAll, func(i, j int) bool {
		return itemsAll[i] < itemsAll[j]
	})
	for _, value := range itemsAll {
		if value > num {
			return value
		}
	}
	return -1
}

func FormattedDivision(num1 int, num2 int) string {
	if num2 == 0 {
		return fmt.Sprintf("0.0000")
	}
	roundme := func(x, unit float64) float64 {
		return math.Round(x/unit) * unit
	}
	text := fmt.Sprintf("%0.4f", roundme(float64(num1)/float64(num2), 0.0001))
	parts := strings.Split(text, ".")
	if len(parts) != 2 {
		return fmt.Sprintf("0.0000")
	}
	first := ""
	index := 0
	reverse := func(text string) string {
		var str []string
		for _, value := range text {
			str = append(str, string(value))
		}
		MAX := len(str) / 2
		for u := 0; u < MAX; u++ {
			temp := str[u]
			str[u] = str[len(str)-1-u]
			str[len(str)-1-u] = temp
		}
		return strings.Join(str, "")
	}
	info := parts[0]
	for u := len(info) - 1; u >= 0; u-- {
		if index%3 == 0 && index != 0 {
			first += fmt.Sprintf(",%v", string(info[u]))
		} else {
			first += fmt.Sprintf("%v", string(info[u]))
		}
		index++
	}
	return fmt.Sprintf("%v.%v", reverse(first), parts[1])
}

//"Hello -5LOL6"    ->     hELLO -6lol5
//"2S 6 du5d4e"     ->     2s 6 DU4D5E
func SwapII(str string) string {
	if len(str) == 0 {
		return ""
	}
	isNumber := func(val rune) (int, bool) {
		if val >= 48 && val <= 57 {
			return int(val) - 48, true
		}
		return -1, false
	}

	changer := func(strValue, with string) string {
		formatted := ""
		if strings.Contains(strValue, with) || len(strValue) != 0 {
			values := strings.Split(str, with)
			index := 0
			for _, value := range values {
				var text []string
				for _, in := range value {
					text = append(text, string(in))
				}
				firstIndex := 0
				firstF := false
				newVal := ""
				isCharDetected := 0
				for u := 0; u < len(text); u++ {
					addValue := ""
					add := func() {
						if []rune((text[u]))[0] >= 65 && []rune((text[u]))[0] < 97 {
							addValue = strings.ToLower(text[u])
						} else {
							addValue = strings.ToUpper(text[u])
						}
					}
					if _, isNum := isNumber([]rune((text[u]))[0]); isNum {
						if !firstF {
							firstF = true
							firstIndex = u
							addValue = text[u]
							isCharDetected = 1
						} else {
							if isCharDetected == 2 {
								temp := text[u]
								text[u] = text[firstIndex]
								var forNewVal []string
								for _, char := range newVal {
									forNewVal = append(forNewVal, string(char))
								}
								forNewVal = append(forNewVal, text[u])
								forNewVal[u] = text[firstIndex]
								text[firstIndex] = temp
								forNewVal[firstIndex] = temp
								newVal = strings.Join(forNewVal, "")
								firstF = false
								isCharDetected = 0
								continue
							} else {
								add()
							}
						}
					} else {
						if isCharDetected == 1 {
							isCharDetected = 2
						}
						add()
					}
					newVal += addValue
				}
				if index == len(values)-1 {
					formatted += newVal
				} else {
					formatted += fmt.Sprintf("%v%v", newVal, with)
				}
				index++
			}
			return formatted
		} else {
			return strValue
		}
	}
	result := changer(str, " ")
	return result
}

/*
input:[]string {"baseball", "a,all,b,ball,bas,base,cat,code,d,e,quit,z"}
Output:"base,ball"
*/
func WordSplit(strArr []string) string {
	if len(strArr) != 2 || (len(strArr[0]) == 0 || len(strArr[1]) == 0) {
		return "not possible"
	}
	var powerGenerator func([]string, int, []string, *[][]string, int)
	powerGenerator = func(subset []string, breakPoint int, selectedSoFar []string, results *[][]string, maxLength int) {
		if len(subset) == breakPoint {
			if len(selectedSoFar) == maxLength {
				arr := make([]string, maxLength)
				copy(arr, selectedSoFar)
				*results = append(*results, arr)
			}
			return
		}
		selectedSoFar = append(selectedSoFar, subset[breakPoint])
		powerGenerator(subset, breakPoint+1, selectedSoFar, results, maxLength)
		selectedSoFar = selectedSoFar[:len(selectedSoFar)-1]
		powerGenerator(subset, breakPoint+1, selectedSoFar, results, maxLength)
	}

	results := make([][]string, 0)
	selectedSoFar := make([]string, 0)
	items := strings.Split(strArr[1], ",")
	powerGenerator(items, 0, selectedSoFar, &results, 2)
	for _, value := range results {
		tempVal := ""
		var reverseEdition []string
		for _, val := range value {
			tempVal += val
			reverseEdition = append(reverseEdition, val)
		}
		if strArr[0] == tempVal {
			return strings.Join(value, ",")
		}
		MAX := len(reverseEdition) / 2
		for u := 0; u < MAX; u++ {
			temp := reverseEdition[u]
			reverseEdition[u] = reverseEdition[len(reverseEdition)-1-u]
			reverseEdition[len(reverseEdition)-1-u] = temp
		}
		tempVal = ""
		for _, val := range reverseEdition {
			tempVal += val
		}
		if strArr[0] == tempVal {
			return strings.Join(reverseEdition, ",")
		}
	}
	return "not possible"
}

func DistinctList(arr []int) int {
	duplicateCount := 0
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	kv := make(map[int]int, 0)
	for u := 0; u < len(arr); u++ {
		if val, isExist := kv[arr[u]]; isExist {
			kv[arr[u]] = val + 1
		} else {
			kv[arr[u]] = 1
		}
	}
	for _, value := range kv {
		if value != 1 {
			duplicateCount += value - 1
		}
	}
	return duplicateCount
}

//"(coder)(byte))" ->  "0"
func BracketMatcher(str string) string {
	if len(str) == 0 {
		return "0"
	}
	chars := []string{"(", ")"}
	var temp []string
	for _, value := range str {
		val := string(value)
		if val == chars[0] || val == chars[1] {
			temp = append(temp, val)
		}
	}
	valOfString := strings.Join(temp, "")
	for {
		if strings.Index(valOfString, strings.Join(chars, "")) != -1 {
			valOfString = strings.Replace(valOfString, strings.Join(chars, ""), "", -1)
		} else {
			break
		}
	}
	if len(valOfString) == 0 {
		return "1"
	}
	return "0"
}

func TripleDouble(num1 int, num2 int) int {
	kvGen := func(num int) map[int]int {
		kv := make(map[int]int, 0)
		index := 0
		strVal := strconv.Itoa(num)
		for _, value := range strVal {
			val, _ := strconv.Atoi(string(value))
			if valInKv, isExist := kv[val]; isExist {
				if string(strVal[index-1]) == string(value) {
					kv[val] = valInKv + 1
				}
			} else {
				kv[val] = 1
			}
			index++
		}
		return kv
	}
	kvFirst := kvGen(num1)
	kvSecond := kvGen(num2)
	for key, value := range kvFirst {
		if value >= 3 {
			for keyIn, valueIn := range kvSecond {
				if valueIn >= 2 && key == keyIn {
					return 1
				}
			}
		}
	}
	return 0
}

//"12:30pm-12:00am"
func CountingMinutes(str string) string {
	if len(str) == 0 {
		return "0"
	}
	parts := strings.Split(str, "-")
	if len(parts) != 2 {
		return "0"
	}
	partFirst := strings.Split(parts[0], ":")
	if len(partFirst) != 2 {
		return "0"
	}
	hFirst, _ := strconv.Atoi(partFirst[0])
	firstHType := "am"
	if strings.Contains(partFirst[1], "pm") {
		firstHType = "pm"
	}
	minuteFirst := strings.Replace(partFirst[1], "pm", "", -1)
	minuteFirst = strings.Replace(minuteFirst, "am", "", -1)
	mFirst, _ := strconv.Atoi(minuteFirst)

	partSecond := strings.Split(parts[1], ":")
	if len(partSecond) != 2 {
		return "0"
	}

	hSecond, _ := strconv.Atoi(partSecond[0])
	secondHType := "am"
	if strings.Contains(partSecond[1], "pm") {
		secondHType = "pm"
	}
	minuteSecond := strings.Replace(partSecond[1], "pm", "", -1)
	minuteSecond = strings.Replace(minuteSecond, "am", "", -1)
	mSecond, _ := strconv.Atoi(minuteSecond)

	tt := time.Now()
	if hFirst < 12 && firstHType == "pm" {
		hFirst += 12
	}
	if hSecond < 12 && secondHType == "pm" {
		hSecond += 12
	}
	diffFirst, diffSecond := 0, 0
	if hFirst == 12 && mFirst == 0 && firstHType == "am" {
		hFirst = 0
		diffFirst = 1
	}
	if hSecond == 12 && mSecond == 0 && secondHType == "am" {
		hSecond = 0
		diffSecond = 1
	}

	first := time.Date(tt.Year(), tt.Month(), tt.Day()+diffFirst, hFirst, mFirst, 0, 0, time.UTC)
	second := time.Date(tt.Year(), tt.Month(), tt.Day()+diffSecond, hSecond, mSecond, 0, 0, time.UTC)
	valFirst, _ := strconv.Atoi(fmt.Sprintf("%02d%02d", hFirst, mFirst))
	valSecond, _ := strconv.Atoi(fmt.Sprintf("%02d%02d", hSecond, mSecond))

	if firstHType == "pm" && secondHType == "pm" {
		if valFirst > valSecond {
			second = second.AddDate(0, 0, 1)
			return fmt.Sprintf("%v", second.Sub(first).Minutes())
		} else if valSecond > valFirst {
			return fmt.Sprintf("%v", second.Sub(first).Minutes())
		} else {
			return "0"
		}
	}
	if firstHType == "am" && secondHType == "am" {
		if valFirst > valSecond {
			second = second.AddDate(0, 0, 1)
			return fmt.Sprintf("%v", second.Sub(first).Minutes())
		} else if valSecond > valFirst {
			return fmt.Sprintf("%v", second.Sub(first).Minutes())
		} else {
			return "0"
		}
	}
	if firstHType == "am" && secondHType == "pm" {
		return fmt.Sprintf("%v", second.Sub(first).Minutes())
	}
	if firstHType == "pm" && secondHType == "am" {
		if second.Hour() == 0 && second.Minute() == 0 && secondHType == "am" {
			return fmt.Sprintf("%v", second.Sub(first).Minutes())
		} else {
			secondT := second.AddDate(0, 0, 1)
			secondT = time.Date(secondT.Year(), secondT.Month(), secondT.Day(), 0, 0, 0, 0, time.UTC)
			firstP := secondT.Sub(first).Minutes()
			secondT = secondT.AddDate(0, 0, -1)
			secondP := second.Sub(secondT).Minutes()
			return fmt.Sprintf("%v", firstP+secondP)
		}
	}
	return "0"
}

//"3x + 12 = 46"

func MissingDigit(str string) string {
	if len(str) == 0 {
		return ""
	}
	str = strings.Replace(str, " ", "", -1)
	parts := strings.Split(str, "=")
	if len(parts) != 2 {
		return ""
	}
	isNumber := func(val string) (bool, int) {
		for _, value := range val {
			if !(value >= 48 && value <= 57) {
				return false, -1
			}
		}
		num, _ := strconv.Atoi(val)
		return true, num
	}

	right := math.MinInt32
	for _, values := range parts {
		if isNum, val := isNumber(values); isNum {
			right = val
			break
		}
	}
	splitterFind := func(valueStr string) string {
		splitter := ""
		if strings.Contains(valueStr, "+") {
			splitter = "+"
		} else if strings.Contains(valueStr, "-") {
			splitter = "-"
		} else if strings.Contains(valueStr, "/") {
			splitter = "/"
		} else if strings.Contains(valueStr, "*") {
			splitter = "*"
		}
		return splitter
	}
	splitOp := func(numValL1, numValL2 int, splitter string) int {
		if splitter == "+" {
			numValL1 += numValL2
		} else if splitter == "-" {
			numValL1 -= numValL2
		} else if splitter == "*" {
			numValL1 *= numValL2
		} else if splitter == "/" {
			numValL1 /= numValL2
		}
		return numValL1
	}
	splitOpAlt := func(numValL1, numValL2 int, splitter string) int {
		if splitter == "+" {
			numValL1 -= numValL2
		} else if splitter == "-" {
			numValL1 += numValL2
		} else if splitter == "*" {
			numValL1 /= numValL2
		} else if splitter == "/" {
			numValL1 *= numValL2
		}
		return numValL1
	}
	if right == math.MinInt32 {
		eq := ""
		for _, values := range parts {
			if strings.Contains(values, "x") {
				eq = values
				break
			}
		}
		for _, values := range parts {
			if !strings.Contains(values, "x") {
				splitter := splitterFind(values)
				if len(splitter) != 0 {
					data := strings.Split(values, splitter)
					if len(data) != 2 {
						return ""
					}
					numValL1, _ := strconv.Atoi(data[0])
					numValL2, _ := strconv.Atoi(data[1])
					numValL1 = splitOp(numValL1, numValL2, splitter)
					numValL1Str := fmt.Sprintf("%v", numValL1)
					for u := 0; u < len(eq); u++ {
						if string(eq[u]) == "x" {
							return string(numValL1Str[u])
						}
					}
				} else {
					return ""
				}
			}
		}
	}

	for _, values := range parts {
		if isNum, _ := isNumber(values); isNum {
			continue
		} else {
			splitter, eq := splitterFind(values), ""
			if len(splitter) != 0 {
				data := strings.Split(values, splitter)
				numVal := 0
				if len(data) != 2 {
					return ""
				}
				neg := false
				if isNumVal, val := isNumber(data[0]); isNumVal {
					numVal = val
					eq = data[1]
					if splitter == "-" {
						neg = true
					}
				} else {
					var err error
					numVal, err = strconv.Atoi(data[1])
					if err != nil {
						return ""
					} else {
						eq = data[0]
					}
				}
				if neg {
					numVal *= -1
				}
				right = splitOpAlt(right, numVal, splitter)
				rightStr := fmt.Sprintf("%v", right)
				for u := 0; u < len(eq); u++ {
					if string(eq[u]) == "x" {
						return string(rightStr[u])
					}
				}
				return ""
			} else {
				return ""
			}
		}
	}
	return ""
}

func PrimeChecker(num int) int {
	isPrime := func(num int) bool {
		for i := 2; i <= int(math.Floor(float64(num)/2)); i++ {
			if num%i == 0 {
				return false
			}
		}
		if num <= 1 {
			return false
		}
		return true
	}
	reverse := func(text string) string {
		var str []string
		for _, value := range text {
			str = append(str, string(value))
		}
		MAX := len(str) / 2
		for u := 0; u < MAX; u++ {
			temp := str[u]
			str[u] = str[len(str)-1-u]
			str[len(str)-1-u] = temp
		}
		return strings.Join(str, "")
	}
	result := make([][]int, 0)
	var perm func([]int, int)
	perm = func(arr []int, n int) {
		if n == 1 {
			temp := make([]int, len(arr))
			copy(temp, arr)
			result = append(result, temp)
		} else {
			for i := 0; i < n; i++ {
				perm(arr, n-1)
				if n%2 == 1 {
					temp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = temp
				} else {
					temp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = temp
				}
			}
		}
	}
	var nums []int
	for _, value := range fmt.Sprintf("%v", num) {
		val, err := strconv.Atoi(string(value))
		if err != nil {
			return 0
		}
		nums = append(nums, val)
	}

	perm(nums, len(nums))

	for _, value := range result {
		norVal := ""
		for _, val := range value {
			norVal += fmt.Sprintf("%v", val)
		}
		valNor, errNor := strconv.Atoi(norVal)
		if errNor != nil {
			return 0
		}
		valRev, errRev := strconv.Atoi(reverse(norVal))
		if errRev != nil {
			return 0
		}
		if isPrime(valNor) || isPrime(valRev) {
			return 1
		}
	}
	return 0
}

func PrimeMover(num int) int {
	isPrime := func(num int) bool {
		for i := 2; i <= int(math.Floor(float64(num)/2)); i++ {
			if num%i == 0 {
				return false
			}
		}
		if num <= 1 {
			return false
		}
		return true
	}
	index := 0
	result := 0
	for u := 0; u < math.MaxInt32 && index != num; u++ {
		if isPrime(u) {
			index++
			result = u
		}
	}
	return result
}

func PairSearching_OLD(num int) int {
	generate := func(numVal int) []int {
		var result []int
		var temp []int
		for _, value := range fmt.Sprintf("%v", numVal) {
			val, _ := strconv.Atoi(string(value))
			temp = append(temp, val)
		}
		for u := 0; u < len(temp); u++ {
			result = append(result, temp[u]*numVal)
		}
		return result
	}

	var first []int
	for _, value := range fmt.Sprintf("%v", num) {
		val, _ := strconv.Atoi(string(value))
		first = append(first, val)
	}
	index := 1
	firstCpy := make([]int, len(first))
	copy(firstCpy, first)
	kv := make(map[int]int, 0)
	f := -1
do:
	resultList := generate(num)
	for j := 0; j < len(resultList); j++ {
		value := resultList[j]
		if _, isExist := kv[value]; isExist {
			continue
		} else {
			kv[value] = 1
		}
		first = make([]int, len(firstCpy))
		copy(first, firstCpy)
		for _, val := range fmt.Sprintf("%v", value) {
			valIn, _ := strconv.Atoi(string(val))
			first = append(first, valIn)
		}
		for u := 1; u < len(first); u++ {
			if first[u-1] == first[u] {
				fmt.Printf("-->>%v[*]\n", first)
				if f == -1 {
					if len(resultList) == 1 {
						return index
					} else {
						f = index
						index = 1
						num = resultList[1]
						goto do
					}
				} else {
					if index < f {
						return index
					}
					return f
				}
			}
		}
		fmt.Printf("%v\n", first)
		if j == len(resultList)-1 {
			temp := generate(resultList[0])
			for _, value := range temp {
				if resultList[0] == value {
					continue
				} else {
					num = value
					break
				}
			}
		}
		index++
		goto do
	}
	return 0
}

func PairSearching_OLD_2(num int) int {
	var first []int
	for _, value := range fmt.Sprintf("%v", num) {
		val, _ := strconv.Atoi(string(value))
		first = append(first, val)
	}
	index := 1
	firstCpy := make([]int, len(first))
	copy(firstCpy, first)
	kv := make(map[int]int, 0)
	isDone := false
	var generate func(int)
	generate = func(numVal int) {
		var resultList []int
		var temp []int
		for _, value := range fmt.Sprintf("%v", numVal) {
			val, _ := strconv.Atoi(string(value))
			temp = append(temp, val)
		}
		for u := 0; u < len(temp); u++ {
			resultList = append(resultList, temp[u]*numVal)
		}
		sort.Slice(resultList, func(i, j int) bool {
			return resultList[i] < resultList[j]
		})
		for j := 0; j < len(resultList) && !isDone; j++ {
			value := resultList[j]
			if _, isExist := kv[value]; isExist {
				continue
			} else {
				kv[value] = 1
			}
			first = make([]int, len(firstCpy))
			copy(first, firstCpy)
			for _, val := range fmt.Sprintf("%v", value) {
				valIn, _ := strconv.Atoi(string(val))
				first = append(first, valIn)
			}
			for u := 2; u < len(first); u++ {
				if first[u-2] == first[u-1] && first[u-1] == first[u] {
					fmt.Printf("%v[*]---[%v]\n", first, num)
					isDone = true
					return
				}
			}
			fmt.Printf("%v\n", first)
			generate(resultList[j])
		}
		index++
	}
	generate(num)
	return index
}

func PairSearching(num int) int {
	var first []int
	for _, value := range fmt.Sprintf("%v", num) {
		val, _ := strconv.Atoi(string(value))
		first = append(first, val)
	}
	firstCpy := make([]int, len(first))
	copy(firstCpy, first)
	itemsQ := make([]int, 0)
	index := 0
	list := make([][]string, 0)
	checkMe := func(results []int) bool {
		for j := 0; j < len(results); j++ {
			value := results[j]

			/* ******** */
			first = make([]int, 0)
			for _, q := range itemsQ {
				for _, value := range fmt.Sprintf("%v", q) {
					val, _ := strconv.Atoi(string(value))
					first = append(first, val)
				}
			}
			/* ******** */

			for _, val := range fmt.Sprintf("%v", value) {
				valIn, _ := strconv.Atoi(string(val))
				first = append(first, valIn)
			}
			for u := 1; u < len(first); u++ {
				if first[u-1] == first[u] {
					//fmt.Printf("%v[*]---[%v]----[%v]\n", first, num, index)
					list = append(list, []string{fmt.Sprintf("%v", first)})
					return true
				}
			}
		}
		return false
	}
	j := 0
	check := true
	var resultList []int
	var generate func(int, *[][] string)
	generate = func(numVal int, items *[][]string) {
	do:
		itemsQ = append(itemsQ, numVal)
		if check {
			resultList = make([]int, 0)
			index++
		}
		resultListIn := make([]int, 0)
		var temp []int
		for _, value := range fmt.Sprintf("%v", numVal) {
			val, _ := strconv.Atoi(string(value))
			temp = append(temp, val)
		}
		for u := 0; u < len(temp); u++ {
			if temp[u] != 1 {
				resultListIn = append(resultListIn, temp[u]*numVal)
				if check {
					resultList = append(resultList, temp[u]*numVal)
				}
			}
		}
		if checkMe(resultListIn) {
			if !check {
				index++
			}
			return
		}
		for k := j; k < len(resultList); {
			numVal = resultList[j]
			j++
			if j == len(resultList) {
				if j >= 2 {
					itemsQ = itemsQ[:len(itemsQ)-1]
				}
				j = 0
				check = true
			} else {
				check = false
			}
			goto do
		}
	}
	generate(num, &list)
	return index
}

func PlusMinus(num int) string {
	notPossible := "not possible"
	numStr := fmt.Sprintf("%v", num)
	if len(numStr) < 2 {
		return notPossible
	}
	allItems := make([][]string, 0)
	var perm func([]string, int)
	perm = func(items []string, n int) {
		if n == 1 {
			isFound := false
			temp := make([]string, len(items))
			copy(temp, items)
			for _, value := range allItems {
				if fmt.Sprintf("%v", value) == fmt.Sprintf("%v", temp) {
					isFound = true
					break
				}
			}
			if !isFound {
				allItems = append(allItems, temp)
			}
		} else {
			for i := 0; i < n; i++ {
				perm(items, n-1)
				if n%2 == 1 {
					temp := items[i]
					items[i] = items[n-1]
					items[n-1] = temp
				} else {
					temp := items[0]
					items[0] = items[n-1]
					items[n-1] = temp
				}
			}
		}
	}

	str := make([]string, 0)
	for _, value := range numStr {
		str = append(str, string(value))
	}
	charGenerator := func(val string, count int) string {
		result := ""
		for u := 0; u < count; u++ {
			result += val
		}
		return result
	}

	minSign, plusSign := "-", "+"
	generator := func(minVal, plusVal int, up bool) []string {
		minCount := minVal
		plusCount := plusVal
		list := make([]string, 0)
		for u := 0; u < len(numStr); u++ {
			minStr := charGenerator(minSign, minCount)
			plusStr := charGenerator(plusSign, plusCount)
			if up {
				list = append(list, fmt.Sprintf("%v%v", minStr, plusStr))
				minCount++
				plusCount--
			} else {
				list = append(list, fmt.Sprintf("%v%v", plusStr, minStr))
				minCount--
				plusCount++
			}
		}
		list = list[1:]
		list = list[:len(list)-1]
		return list
	}

	opListPlusFirst := generator(0, len(numStr)-1, true)
	opListMinFirst := generator(len(numStr)-1, 0, false)

	list := make([]string, 0)
	for _, value := range opListPlusFirst {
		list = append(list, value)
	}
	for _, value := range opListMinFirst {
		list = append(list, value)
	}

	for _, value := range list {
		perm(strings.Split(value, ""), len(value))
	}
	allItems = append(allItems, strings.Split(charGenerator(minSign, len(numStr)-1), ""))
	allItems = append(allItems, strings.Split(charGenerator(plusSign, len(numStr)-1), ""))
	for _, signs := range allItems {
		index, a, b, signIndex := 0, 0, 0, 0
		for _, valueIn := range numStr {
			valInt, _ := strconv.Atoi(string(valueIn))
			if index == 0 {
				a = valInt
				index++
			} else if index == 1 {
				b = valInt
				val := signs[signIndex]
				if val == minSign {
					a -= b
				} else {
					a += b
				}
				index++
				signIndex++
			} else {
				val := signs[signIndex]
				if val == minSign {
					a -= valInt
				} else {
					a += valInt
				}
				signIndex++
			}
		}
		if a == 0 {
			fmt.Printf("%v\n", fmt.Sprintf("%v", strings.Join(signs, "")))
			//return fmt.Sprintf("%v", strings.Join(signs, ""))
		}
	}
	return "" //notPossible
}

func FibonacciChecker(num int) string {
	var fib func(int) int
	fib = func(num int) int {
		if num == 0 || num == 1 {
			return 1
		}
		return fib(num-1) + fib(num-2)
	}
	result := 0
	for u := 0; u < math.MaxInt32 && result <= num; u++ {
		result = fib(u)
		if result == num {
			return "yes"
		}
	}
	return "no"
}

func MultipleBrackets(str string) string {
	if len(str) == 0 {
		return "0"
	}
	strArr := strings.Split(str, "")
	validBrackets := []string{"[", "]", "(", ")"}
	pruneStr := ""
	anyFound := false
	for _, value := range strArr {
		isFound := false
		tmpChar := ""
		for _, char := range validBrackets {
			if char == value {
				anyFound = true
				tmpChar = char
				isFound = true
				break
			}
		}
		if isFound {
			pruneStr += tmpChar
		}
	}
	if !anyFound {
		return "1"
	}
	index := 0
do:
	if len(pruneStr) != 0 {
		changer := fmt.Sprintf("%v%v", validBrackets[0], validBrackets[1])
		if strings.Count(pruneStr, changer) != 0 {
			index += strings.Count(pruneStr, changer)
			pruneStr = strings.Replace(pruneStr, changer, "", -1)
		}
		changer = fmt.Sprintf("%v%v", validBrackets[2], validBrackets[3])
		if strings.Count(pruneStr, changer) != 0 {
			index += strings.Count(pruneStr, changer)
			pruneStr = strings.Replace(pruneStr, changer, "", -1)
		}
	}
	if len(pruneStr) != 0 {
		if strings.Index(pruneStr, fmt.Sprintf("%v%v", validBrackets[0], validBrackets[1])) != -1 ||
			strings.Index(pruneStr, fmt.Sprintf("%v%v", validBrackets[2], validBrackets[3])) != -1 {
			goto do
		}
		return "0"
	} else {
		return fmt.Sprintf("%v %v", "1", index)
	}
}

func SimpleSwapper(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func Variadic(values ...int) int {
	res := 0
	for _, value := range values {
		res += value
	}
	return res
}

type LinkNode struct {
	Value    int
	Head     *LinkNode
	Last     *LinkNode
	NextItem *LinkNode
}

func NewLinkNode(value int) *LinkNode {
	return &LinkNode{
		Value:    value,
		NextItem: nil,
	}
}

func (l *LinkNode) Next() *LinkNode {
	return l.NextItem
}

func (l *LinkNode) Add(value int) {
	element := NewLinkNode(value)
	if l.Head == nil {
		l.Head = element
	} else {
		l.Last.NextItem = element
	}
	l.Last = element
}

func SelectionSort(elements *[]int) {
	for u := 0; u < len(*elements)-1; u++ {
		min := u
		for j := u + 1; j < len(*elements); j++ {
			if (*elements)[j] < (*elements)[min] {
				min = j
			}
		}
		temp := (*elements)[u]
		(*elements)[u] = (*elements)[min]
		(*elements)[min] = temp
	}
}

//[]string {"12:15PM-02:00PM","09:00AM-10:00AM","10:30AM-12:00PM"}
func MostFreeTime(strArr []string) string {
	if len(strArr) < 3 {
		return "00:00"
	}
	padLeft := func(val string, with string, length int) string {
		if len(val) >= length {
			return val
		}
		result := val
		for u := len(val); u < length; u++ {
			with += result
		}
		return with
	}

	hour := 60
	dayH := 24
	day := hour * dayH
	timeSlotsKV := make(map[string]int, day)
	hh, mm, suffix := 0, 0, "AM"
	indexes := make([]string, 0)
	for u := 0; u < day; u++ {
		if mm >= hour {
			mm = 0
		}
		if u >= 0 && u < hour || u >= hour*(dayH/2) && u < (hour*(dayH/2))+hour {
			hh = dayH / 2
		} else {
			hh = u / hour
		}
		if u >= hour*dayH/2 {
			suffix = "PM"
		}
		if hh > dayH/2 {
			hh -= dayH / 2
		}
		key := fmt.Sprintf("%v:%v%v", padLeft(strconv.Itoa(hh), "0", 2), padLeft(strconv.Itoa(mm), "0", 2), suffix)
		timeSlotsKV[key] = u
		indexes = append(indexes, key)
		mm++
	}
	/*
		for _, value := range indexes {
			fmt.Printf("%v-%v\n",value,timeSlotsKV[value])
		}
	*/

	timeToIndex := func(timeVal string) int {
		if len(timeVal) == 0 {
			return 0
		}
		if val, exist := timeSlotsKV[timeVal]; exist {
			return val
		} else {
			return 0
		}
	}

	var items []int
	for _, value := range strArr {
		splitted := strings.Split(value, "-")
		for _, split := range splitted {
			items = append(items, timeToIndex(split))
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
	if len(items) < 3 {
		return "00:00"
	}
	max := items[2] - items[1]
	for u := 4; u < len(items); u += 2 {
		if items[u]-items[u-1] > max {
			max = items[u] - items[u-1]
		}
	}
	hh, mm = 0, 0
	if max >= hour {
		hh = max / hour
		mm = max - (hh * hour)
		return fmt.Sprintf("%v:%v", padLeft(strconv.Itoa(hh), "0", 2), padLeft(strconv.Itoa(mm), "0", 2))
	} else {
		return fmt.Sprintf("00:%v", padLeft(strconv.Itoa(max), "0", 2))
	}
}

func PrintRec(to int) {
	if to == 0 {
		return
	}
	fmt.Printf("%v\n", to)
	PrintRec(to - 1)
}

func SumRec(val int) int {
	if val == 0 {
		return val
	}
	return val + SumRec(val-1)
}

func FindMaxValRec(val []int) int {
	if len(val) < 2 {
		return val[0]
	}
	if val[1] >= val[0] {
		return FindMaxValRec(val[1:])
	} else {
		return FindMaxValRec(append(val[:1], val[2:]...))
	}
}

//noinspection ALL,GoUnreachableCode
func CoinDeterminer(num int) int {
	coins := []int{1, 5, 7, 9, 11}
	results := make([][]int, 0)
	sF := make([]int, 0)
	var powerSet func([]int, int, []int, *[][]int)
	powerSet = func(items []int, breakPoint int, selectedSoFar []int, results *[][]int) {
		if breakPoint == len(items) {
			temp := make([]int, len(selectedSoFar))
			copy(temp, selectedSoFar)
			*results = append(*results, temp)
			return
		}
		selectedSoFar = append(selectedSoFar, items[breakPoint])
		powerSet(items, breakPoint+1, selectedSoFar, results)
		selectedSoFar = selectedSoFar[:len(selectedSoFar)-1]
		powerSet(items, breakPoint+1, selectedSoFar, results)
	}
	powerSet(coins, 0, sF, &results)

	sort.Slice(results, func(i, j int) bool {
		return len(results[i]) < len(results[j])
	})

	var sumIntsByRec func([]int) int
	sumIntsByRec = func(arr []int) int {
		if len(arr) == 0 {
			return 0
		}
		return arr[0] + sumIntsByRec(arr[1:])
	}
	for _, value := range results {
		if sumIntsByRec(value) == num {
			return len(value)
		}
	}

	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	resultsKV := make(map[int]int, len(coins))

	for j := 0; j < len(coins); j++ {
		val := j
		isClean := false
		arrNew := make([]int, 0)
	do:
		for u := val; u < len(coins); u++ {
			if isClean {
				arrNew = make([]int, 0)
			}
			for {
				tempArr := append(arrNew, coins[u])
				if sumIntsByRec(tempArr) == num {
					resultsKV[coins[j]] = len(tempArr)
					break
				}
				if sumIntsByRec(tempArr) > num {
					val = u + 1
					isClean = false
					goto do
				}
				arrNew = append(arrNew, coins[u])
			}
			isClean = true
		}
	}
	min, index := 0, 0
	for _, value := range resultsKV {
		if index == 0 {
			min = value
			index++
			continue
		}
		if value < min {
			min = value
		}
	}
	return min
}

func LookSaySequence(num int) int {
	items := strings.Split(strconv.Itoa(num), "")
	result := ""
	for u := 0; u < len(items); u++ {
		index := 1
		for j := u + 1; j < len(items); j++ {
			if items[u] == items[j] {
				index++
				u++
			} else {
				break
			}
		}
		result += fmt.Sprintf("%v%v", index, items[u])
	}

	val, err := strconv.Atoi(result)
	if err == nil {
		return val
	}
	return 0
}

//[]string {"(0,0),(0,-2),(3,0),(3,-2),(2,-1),(3,-1),(2,3),(3,3)"}
func OverlappingRectangles(strArr []string) string {
	if len(strArr) != 1 {
		return "0"
	}
	points := make([]string, 0)
	index := 0
	pos := 0
	for len(strArr[0]) != 0 {
		if string(strArr[0][pos]) == "," {
			if index%2 == 1 {
				points = append(points, strArr[0][:pos])
				strArr[0] = strArr[0][pos+1:]
				pos = 0
				index++
				continue
			} else {
				if strings.Count(strArr[0], ",") == 1 {
					points = append(points, strArr[0])
					break
				}
			}
			index++
		}
		pos++
	}
	squareMinDotCount := 4
	type Point struct {
		x int
		y int
	}
	type Rectangle struct {
		p1 Point
		p2 Point
		p3 Point
		p4 Point
	}

	var rect1, rect2 Rectangle
	for u := 0; u < len(points); u++ {
		val := strings.Replace(strings.Replace(points[u], "(", "", -1), ")", "", -1)
		values := strings.Split(val, ",")
		if len(values) != 2 {
			return "0"
		}
		x, errX := strconv.Atoi(values[0])
		if errX != nil {
			return "0"
		}
		y, errY := strconv.Atoi(values[1])
		if errY != nil {
			return "0"
		}
		switch u {
		case 0:
			{
				rect1.p1 = Point{x: x, y: y}
			}
		case 1:
			{
				rect1.p2 = Point{x: x, y: y}

			}
		case 2:
			{
				rect1.p3 = Point{x: x, y: y}

			}
		case 3:
			{
				rect1.p4 = Point{x: x, y: y}

			}
		case 4:
			{
				rect2.p1 = Point{x: x, y: y}

			}
		case 5:
			{
				rect2.p2 = Point{x: x, y: y}

			}
		case 6:
			{
				rect2.p3 = Point{x: x, y: y}

			}
		case 7:
			{
				rect2.p4 = Point{x: x, y: y}
			}
		}
	}
	dimensions := func(rect Rectangle) (int, int, Point) {
		h, w := 0, 0
		for {
			if rect.p1.x == rect.p2.x {
				h = int(math.Abs(float64(rect.p1.y) - float64(rect.p2.y)))
				w = int(math.Abs(float64(rect.p3.x) - float64(rect.p1.x)))
				break
			}
			if rect.p1.x == rect.p3.x {
				h = int(math.Abs(float64(rect.p1.y) - float64(rect.p3.y)))
				w = int(math.Abs(float64(rect.p2.x) - float64(rect.p1.x)))
				break
			}
			if rect.p1.x == rect.p4.x {
				h = int(math.Abs(float64(rect.p1.y) - float64(rect.p4.y)))
				w = int(math.Abs(float64(rect.p2.x) - float64(rect.p1.x)))
				break
			}
			break
		}
		return h, w, rect.p1
	}
	matrixGenerator := func(rect Rectangle) [][]string {
		h, w, p := dimensions(rect)
		result := make([][]string, 0)
		for u := p.y; u <= h; u++ {
			for j := p.x; j <= w; j++ {
				result = append(result, []string{fmt.Sprintf("[%v,%v]", j, u)})
			}
		}
		return result
	}
	matrisForRect1 := matrixGenerator(rect1)
	h, w, _ := dimensions(rect1)
	sofRect1 := h * w
	matrisForRect2 := matrixGenerator(rect2)
	fmt.Printf("%v\n", matrisForRect1)
	fmt.Printf("%v\n", matrisForRect2)
	count := 0
	for _, valFor1 := range matrisForRect1 {
		for _, valFor2 := range matrisForRect2 {
			if fmt.Sprintf("%v", valFor1) == fmt.Sprintf("%v", valFor2) {
				count++
			}
		}
	}
	if count < squareMinDotCount {
		return "0"
	} else if count >= squareMinDotCount {
		return fmt.Sprintf("%v", sofRect1/count)
	}
	return "0"
}

//3aabacbebebe -> cbebebebe
func KUniqueCharacters(str string) string {
	min, max := 1, 6
	if len(str) == 0 {
		return ""
	}
	count := string(str[0])
	countInt, err := strconv.Atoi(count)
	if err != nil {
		return ""
	}
	if countInt > max && countInt < min {
		return ""
	}
	raw := strings.Split(str, "")
	second := strings.Join(raw[1:], "")
	items := make([]string, 0)
	for u := 0; u < len(second); u++ {
		kvChars := make(map[string]int, 0)
		word := ""
		for j := u; j < len(second); j++ {
			if val, isExist := kvChars[string(second[j])]; isExist {
				kvChars[string(second[j])] = val + 1
				word += string(second[j])
			} else {
				if len(kvChars) == countInt {
					break
				}
				if len(kvChars) < countInt {
					kvChars[string(second[j])] = 1
					word += string(second[j])
				}
			}
		}
		if len(kvChars) < countInt {
			continue
		}
		items = append(items, word)
	}
	backup := make([]string, len(items))
	copy(backup, items)
	sort.Slice(items, func(i, j int) bool {
		return len(items[i]) > len(items[j])
	})
	if len(items) != 0 {
		tempLen := len(items[0])
		for u := 0; u < len(backup); u++ {
			if tempLen == len(backup[u]) {
				return backup[u]
			}
		}
		return items[0]
	}
	return ""
}

// "A", "B", "C", "D", "A", "E", "D", "Z"  ----> C-A-E-D-Z
func LRUCache(strArr []string) string {
	if len(strArr) == 0 {
		return ""
	}
	cacheLimit := 5
	items := make([]string, 0)
	kvItems := make(map[string]int, 0)
	for u := 0; u < len(strArr); u++ {
		if _, isExist := kvItems[strArr[u]]; isExist {
			tmpItems := make([]string, 0)
			for j := 0; j < len(items); j++ {
				if items[j] != strArr[u] {
					tmpItems = append(tmpItems, items[j])
				}
			}
			if len(tmpItems) == cacheLimit {
				tmpItems = tmpItems[1:]
			}
			tmpItems = append(tmpItems, strArr[u])
			items = make([]string, len(tmpItems))
			copy(items, tmpItems)
		} else {
			if len(items) == cacheLimit {
				items = items[1:]
			}
			items = append(items, strArr[u])
		}
		kvItems[strArr[u]] = 1
	}
	return strings.Join(items, "-")
}

func ArrayMinJumps_V1(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	fmt.Printf("\n\n--------- >>>>>>> %v\n\n", arr)
	hasher := 10000
	getRealValue := func(val, index int) int {
		hasherIndex := val / hasher
		if hasherIndex == 0 {
			hasherIndex = 1
		}
		if val >= hasher*hasherIndex {
			return val - (hasher * hasherIndex) - (index * index * index)
		}
		return val
	}
	alter := func(items []int) []int {
		for u := 0; u < len(items); u++ {
			foundIndexes := make([]int, 0)
			for j := u + 1; j < len(items); j++ {
				if items[u] == items[j] {
					foundIndexes = append(foundIndexes, j)
				}
			}
			if len(foundIndexes) != 0 {
				for _, value := range foundIndexes {
					items[value] = items[value] + hasher + (value * value * value)
				}
			}
		}
		return items
	}
	arr = alter(arr)
	kv := make(map[string][]string, 0)
	indexes := make([]int, 0)
	indexForDones := make(map[int]int, 0)
	dones := make([]int, 0)
	isFound := false
	for u := 0; u < len(arr); u++ {
		var items []string
		index := 0
		j := u + 1
		for index < getRealValue(arr[u], u) && j < len(arr) {
			items = append(items, fmt.Sprintf("%v", getRealValue(arr[j], j)))
			j++
			index++
		}
		if j == len(arr) {
			dones = append(dones, arr[u])
			indexForDones[arr[u]] = u
			items = append(items, "*")
			isFound = true
		}
		kv[fmt.Sprintf("%v", arr[u])] = items
		indexes = append(indexes, arr[u])
		fmt.Printf("[%v]-[%v]\n", getRealValue(arr[u], u), items)
	}
	if !isFound {
		return -1
	}

	fmt.Printf("\n%v\n", "-----------------------------------------")

	kvLock := make(map[int]int, 0)
	min := make([]int, 0)
	for u := 0; u < len(dones); u++ {
		values := kv[fmt.Sprintf("%v", dones[u])]
		total := 1
		once := false
		isKeyUsed := false
		fmt.Printf("\n***************** %v *********************\n", getRealValue(dones[u], indexForDones[dones[u]]))
		for y := len(indexes) - 1; y >= 0; y-- {
			if fmt.Sprintf("%v", kv[fmt.Sprintf("%v", indexes[y])]) != fmt.Sprintf("%v", values) && !once {
				continue
			}
			if !once {
				once = !once
				continue
			}

			fmt.Printf("-->>[%v]-[%v]\n", getRealValue(indexes[y], y), kv[fmt.Sprintf("%v", indexes[y])])

			if !isKeyUsed {
				isKeyUsed = true
				for _, valueIn := range kv[fmt.Sprintf("%v", indexes[y])] {
					if valueIn == fmt.Sprintf("%v", dones[u]) {
						total++
						break
					}
				}
				continue
			} else {
				isKeyFound := false
				for _, valueIn := range kv[fmt.Sprintf("%v", indexes[y])] {
					if valueIn == fmt.Sprintf("%v", dones[u]) {
						isKeyFound = true
						break
					}
				}
				if !isKeyFound {
					total++
				}
			}
		}
		if _, isExist := kvLock[getRealValue(dones[u], indexForDones[dones[u]])]; !isExist {
			kvLock[getRealValue(dones[u], indexForDones[dones[u]])] = total
			min = append(min, total)
		}
	}
	sort.Slice(min, func(i, j int) bool {
		return min[i] < min[j]
	})
	if len(min) != 0 {
		return min[0]
	}
	return -1
}

func ArrayMinJumps(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return 0
	}
	hasher := 10000
	getRealValue := func(val, index int) int {
		hasherIndex := val / hasher
		if hasherIndex == 0 {
			hasherIndex = 1
		}
		if val >= hasher*hasherIndex {
			return val - (hasher * hasherIndex) - (index * index * index)
		}
		return val
	}
	alter := func(items []int) []int {
		for u := 0; u < len(items); u++ {
			foundIndexes := make([]int, 0)
			for j := u + 1; j < len(items); j++ {
				if items[u] == items[j] {
					foundIndexes = append(foundIndexes, j)
				}
			}
			if len(foundIndexes) != 0 {
				for _, value := range foundIndexes {
					items[value] = items[value] + hasher + (value * value * value)
				}
			}
		}
		return items
	}
	arr = alter(arr)
	kv := make(map[string][]string, 0)
	indexes := make([]int, 0)
	indexForSearch := make(map[int]int, 0)
	isFound := false
	for u := 0; u < len(arr); u++ {
		var items []string
		index := 0
		j := u + 1
		for index < getRealValue(arr[u], u) && j < len(arr) {
			items = append(items, fmt.Sprintf("%v", getRealValue(arr[j], j)))
			j++
			index++
		}
		if j == len(arr) {
			items = append(items, "*")
			isFound = true
		}
		kv[fmt.Sprintf("%v", arr[u])] = items
		indexes = append(indexes, arr[u])
		indexForSearch[arr[u]] = u
	}
	if !isFound {
		return -1
	}

	total := 1
	index := 0
do:
	for y := index; y < len(indexes); y++ {
		//fmt.Printf("-->>[%v]-[%v]\n", getRealValue(indexes[y], y), kv[fmt.Sprintf("%v", indexes[y])])
		max := 0
		realIndex := 0
		sub := 0
		for _, valueIn := range kv[fmt.Sprintf("%v", indexes[y])] {
			if valueIn == "*" {
				return total
			}
			val, _ := strconv.Atoi(valueIn)
			if val+realIndex > max {
				max = val + realIndex
				sub = realIndex
			}
			realIndex++
		}
		max -= sub

		indexList := make([]int, 0)
		indis := 0
		for _, valueIn := range kv[fmt.Sprintf("%v", indexes[y])] {
			if valueIn != "*" {
				val, _ := strconv.Atoi(valueIn)
				if val == max {
					indexList = append(indexList, indis)
				}
			}
			indis++
		}
		if max == 0 {
			return -1
		}
		maxCount := len(indexList)
		for j := 0; j < len(indexes); j++ {
			if getRealValue(indexes[j], indexForSearch[indexes[j]]) == max && j > index {
				if maxCount > 1 {
					maxCount--
					continue
				}
				index = j
				total++
				goto do
			}
		}
	}
	if total == 0 {
		return -1
	}
	return total
}

// [5, 2, 8, 3, 9, 12]  - >  -1 -1 2 2 3 9
func NearestSmallerValues(arr []int) string {
	if len(arr) == 0 {
		return ""
	}
	not := "-1"
	result := make([]string, 0)
	result = append(result,fmt.Sprintf("%v",not))
	for u := 1; u < len(arr); u++ {
		isFound := false
		for j := u-1; j>=0; j-- {
			if arr[j]<=arr[u] {
				result = append(result,fmt.Sprintf("%v",arr[j]))
				isFound = true
				break
			}
		}
		if !isFound {
			result = append(result,fmt.Sprintf("%v",not))
		}
	}
	return strings.Join(result, " ")
}

func CharacterRemoval(strArr []string) string {
	if len(strArr)!=2 {
		return "-1"
	}

	var powerSet func([]string,int,[]string,*[][]string)
	powerSet = func(items []string,breakPoint int,selectedSoFar []string,results *[][]string) {
		if len(items) == breakPoint {
			tempSf := make([]string,len(selectedSoFar))
			copy(tempSf,selectedSoFar)
			*results = append(*results,tempSf)
			return
		}
		selectedSoFar = append(selectedSoFar,items[breakPoint])
		powerSet(items,breakPoint+1,selectedSoFar,results)
		selectedSoFar = selectedSoFar[:len(selectedSoFar)-1]
		powerSet(items,breakPoint+1,selectedSoFar,results)
	}
	search := strArr[0]
	for {
		if len(search)!=0 {

			selectedSoFar:=make([]string,0)
			results := make([][]string,0)
			powerSet()


		}
	}
	return ""
}

//noinspection ALL
func main() {
	//fmt.Printf("%v\n", KUniqueCharacters("2aabbaaccbbaaccaabb"))
	//fmt.Printf("%v\n", KUniqueCharacters("2aabbcbbbadef"))

	fmt.Printf("%v\n", NearestSmallerValues([]int{5, 2, 8, 3, 9, 12}))
	fmt.Printf("%v\n", NearestSmallerValues([]int{5, 3, 1, 9, 7, 3, 4, 1}))
	fmt.Printf("%v\n", NearestSmallerValues([]int{2,4,5,1,7}))

	return
	fmt.Printf("%v\n", ArrayMinJumps([]int{4, 5, 2, 1, 5, 3, 1, 4, 6, 2, 1, 0, 1, 0, 4, 3, 0, 1, 2, 4, 5}))
	fmt.Printf("%v\n", ArrayMinJumps([]int{3, 4, 2, 1, 1, 100}))
	fmt.Printf("%v\n", ArrayMinJumps([]int{4}))
	fmt.Printf("%v\n", ArrayMinJumps([]int{1, 0, 0, 2}))
	fmt.Printf("%v\n", ArrayMinJumps([]int{1, 3, 6, 8, 2, 7, 1, 2, 1, 2, 6, 1, 2, 1, 2}))
	return
	fmt.Printf("%v\n", LRUCache([]string{"A", "B", "C", "D", "A", "E", "D", "Z"}))
	fmt.Printf("%v\n", LRUCache([]string{"A", "B", "A", "C", "A", "B"}))
	fmt.Printf("%v\n", LRUCache([]string{"A", "B", "C", "D", "E", "D", "Q", "Z", "C"}))

	return
	fmt.Printf("%v\n", KUniqueCharacters("1aabb"))
	fmt.Printf("%v\n", KUniqueCharacters("4aaffaacccerrfffaacca"))

	return

	fmt.Printf("%v\n", OverlappingRectangles([]string{"(0,0),(0,-2),(3,0),(3,-2),(2,-1),(3,-1),(2,3),(3,3)"})) // 6
	fmt.Printf("%v\n", OverlappingRectangles([]string{"(0,0),(2,2),(2,0),(0,2),(1,0),(1,2),(6,0),(6,2)"}))     // 2
	//!!editing
	return
	fmt.Printf("%v\n", LookSaySequence(110))
	fmt.Printf("%v\n", LookSaySequence(15))
	fmt.Printf("%v\n", LookSaySequence(101))
	fmt.Printf("%v\n", LookSaySequence(1211))
	fmt.Printf("%v\n", LookSaySequence(2466))
	fmt.Printf("%v\n", LookSaySequence(1))
	fmt.Printf("%v\n", LookSaySequence(100))
	fmt.Printf("%v\n", LookSaySequence(1001))
	return
	fmt.Printf("%v\n", CoinDeterminer(100))
	fmt.Printf("%v\n", CoinDeterminer(14))
	fmt.Printf("%v\n", CoinDeterminer(37))
	fmt.Printf("%v\n", CoinDeterminer(34))
	fmt.Printf("%v\n", CoinDeterminer(2))
	fmt.Printf("%v\n", CoinDeterminer(5))
	fmt.Printf("%v\n", CoinDeterminer(4))

	return

	fmt.Printf("%v\n", CoinDeterminer(6))
	fmt.Printf("%v\n", CoinDeterminer(16))
	fmt.Printf("%v\n", CoinDeterminer(25))

	return

	fmt.Printf("%v", FindMaxValRec([]int{100, 32, 444443, 55, 22, 19999, 43, 2}))
	return

	fmt.Printf("%v", SumRec(5))

	return

	fmt.Printf("%v \n", MostFreeTime([]string{"10:00AM-12:30PM", "02:00PM-02:45PM", "09:10AM-09:50AM"}))
	fmt.Printf("%v \n", MostFreeTime([]string{"12:15PM-02:00PM", "09:00AM-12:11PM", "02:02PM-04:00PM"}))
	fmt.Printf("%v \n", MostFreeTime([]string{"12:15PM-02:00PM", "09:00AM-10:00AM", "10:30AM-12:00PM"}))
	return

	arrx := []int{5, 1, 3, 2, 1}
	SelectionSort(&arrx)
	fmt.Printf("%v", arrx)
	return

	ll := &LinkNode{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	item := ll.Head
	for {
		fmt.Printf("%v\n", item.Value)
		if item.Next() != nil {
			item = item.Next()
		} else {
			break
		}
	}

	return

	fmt.Printf("%v \n", Variadic(1, 2, 3, 4))
	tempxx := []int{1, 4, 5, 3}
	fmt.Printf("%v \n", Variadic(tempxx...))

	a, b := 100, 9
	SimpleSwapper(&a, &b)
	fmt.Printf("%v %v", a, b)
	return
	//fmt.Printf("%v\n", MultipleBrackets("(coder)[byte)]"))
	fmt.Printf("%v\n", MultipleBrackets("(c([od]er)) b(yt[e])"))
	return
	fmt.Printf("%v\n", FibonacciChecker(34))
	fmt.Printf("%v\n", FibonacciChecker(54))

	return
	//fmt.Printf("%v\n", PlusMinus(35132))
	//fmt.Printf("%v\n", PlusMinus(199))
	fmt.Printf("%v\n", PlusMinus(26712))
	return
	fmt.Printf("%v\n", PairSearching(8))   // 3
	fmt.Printf("%v\n", PairSearching(15))  // 2
	fmt.Printf("%v\n", PairSearching(2))   // 5
	fmt.Printf("%v\n", PairSearching(46))  // 2
	fmt.Printf("%v\n", PairSearching(5))   // 2
	fmt.Printf("%v\n", PairSearching(134)) // 1
	fmt.Printf("%v\n", PairSearching(198)) // 2
	return

	fmt.Printf("%v\n", PrimeMover(9))
	fmt.Printf("%v\n", PrimeMover(16))
	fmt.Printf("%v\n", PrimeMover(100))
	return
	fmt.Printf("%v\n", PrimeChecker(30))
	fmt.Printf("%v\n", PrimeChecker(98))
	fmt.Printf("%v\n", PrimeChecker(598))
	fmt.Printf("%v\n", PrimeChecker(666))

	return
	//fmt.Printf("%v\n", MissingDigit("4356 * 23 = 1001x8"))
	//fmt.Printf("%v\n", MissingDigit("1 + 1111 = x112"))
	fmt.Printf("%v\n", MissingDigit("2004 / 6 = 33x"))
	fmt.Printf("%v\n", MissingDigit("10 - x = 10"))
	fmt.Printf("%v\n", MissingDigit("4 - 2 = x"))
	fmt.Printf("%v\n", MissingDigit("3x + 12 = 46"))
	fmt.Printf("%v\n", MissingDigit("10x * 12 = 1200"))
	fmt.Printf("%v\n", MissingDigit("1x0 * 10 = 1300"))
	return

	fmt.Printf("%v\n", CountingMinutes("11:00am-2:10pm"))
	fmt.Printf("%v\n", CountingMinutes("5:00pm-5:01pm"))
	fmt.Printf("%v\n", CountingMinutes("5:01pm-5:00pm"))
	fmt.Printf("%v\n", CountingMinutes("2:00pm-3:00pm"))
	fmt.Printf("%v\n", CountingMinutes("2:03pm-1:39pm"))
	fmt.Printf("%v\n", CountingMinutes("12:30pm-12:00am"))
	fmt.Printf("%v\n", CountingMinutes("1:00pm-11:00am"))
	fmt.Printf("%v\n", CountingMinutes("1:23am-1:08am"))

	return

	fmt.Printf("%v\n", TripleDouble(451999277, 41177722899))
	fmt.Printf("%v\n", TripleDouble(465555, 5579))
	fmt.Printf("%v\n", TripleDouble(67844, 66237))
	fmt.Printf("%v\n", TripleDouble(556668, 556886))

	return

	fmt.Printf("%v\n", BracketMatcher("(coder)(byte))"))
	fmt.Printf("%v\n", BracketMatcher("(c(oder)) b(yte)"))

	return
	fmt.Printf("%v\n", DistinctList([]int{0, -2, -2, 5, 5, 5}))
	fmt.Printf("%v\n", DistinctList([]int{1, 2, 2, 2, 3}))
	fmt.Printf("%v\n", DistinctList([]int{100, 2, 101, 4}))

	return
	fmt.Printf("%v\n", WordSplit([]string{"baseball", "a,all,b,ball,bas,base,cat,code,d,e,quit,z"}))
	fmt.Printf("%v\n", WordSplit([]string{"abcgefd", "a,ab,abc,abcg,b,c,dog,e,efd,zzzz"}))

	return
	fmt.Printf("%v\n", SwapII("123gg))(("))
	fmt.Printf("%v\n", SwapII("2S 6 du5d4e"))
	fmt.Printf("%v\n", SwapII("6coderbyte5"))

	return

	fmt.Printf("%v\n", FormattedDivision(101077282, 21123))

	return

	fmt.Printf("%v\n", PermutationStep(23514))
	fmt.Printf("%v\n", PermutationStep(897654321))

	return

	fmt.Printf("%v\n", StringReduction("aaa"))
	fmt.Printf("%v\n", StringReduction("abc"))
	fmt.Printf("%v\n", StringReduction("bcab"))
	fmt.Printf("%v\n", StringReduction("aabc"))
	fmt.Printf("%v\n", StringReduction("abcabc"))
	fmt.Printf("%v\n", StringReduction("abb"))
	fmt.Printf("%v\n", StringReduction("aa"))
	fmt.Printf("%v\n", StringReduction("ccaa"))

	return
	fmt.Printf("%v\n", ThreeFiveMultiples(1))
	return
	fmt.Printf("%v\n", NumberSearch("Hello6 9World 2, Nic8e D7ay!"))
	fmt.Printf("%v\n", NumberSearch("H3ello9-9"))
	fmt.Printf("%v\n", NumberSearch("i love cake 9 8 7"))
	fmt.Printf("%v\n", NumberSearch("3ko6"))

	return

	fmt.Printf("%v\n", Consecutive([]int{5, 10, 15}))
	fmt.Printf("%v\n", Consecutive([]int{-2, 10, 4}))

	return

	fmt.Printf("%v\n", SimpleMode([]int{10, 4, 5, 2, 4}))
	return

	fmt.Printf("%v\n", BinaryConverter("100101"))
	return
	fmt.Printf("%v\n", ArrayAddition([]int{3, 5, -1, 8, 12}))

	return
	fmt.Printf("%v\n", DashInsertII(667488958374553)) // 454*67-9-3
	fmt.Printf("%v\n", DashInsertII(99946))
	fmt.Printf("%v\n", DashInsertII(56647304))

	return

	fmt.Printf("%v\n", LetterCount("Today, is the greatest day ever!"))
	fmt.Printf("%v\n", LetterCount("Hello apple pie!"))
	fmt.Printf("%v\n", LetterCount("No Words"))

	return

	fmt.Printf("%v\n", ArithGeoII([]int{5, 10, 15}))
	fmt.Printf("%v\n", ArithGeoII([]int{5, 10, 11, 20}))
	fmt.Printf("%v\n", ArithGeoII([]int{2, 4, 6, 8}))
	fmt.Printf("%v\n", ArithGeoII([]int{2, 6, 18, 54}))
	fmt.Printf("%v\n", ArithGeoII([]int{2, 4, 16, 24}))

	return

	fmt.Printf("%v\n", StringScramble("cdore", "coder"))
	fmt.Printf("%v\n", StringScramble("h3llko", "hello"))
	fmt.Printf("%v\n", StringScramble("rkqodlw", "world"))

	return

	fmt.Printf("%v\n", Division(7, 3))
	fmt.Printf("%v\n", Division(36, 54))
	fmt.Printf("%v\n", Division(12, 16))

	return
	fmt.Printf("%v\n", PalindromeTwo("Noel - sees Leon"))
	fmt.Printf("%v\n", PalindromeTwo("Anne, I vote more cars race Rome-to-Vienna"))
	fmt.Printf("%v\n", PalindromeTwo("A war at Tarawa!"))

	return
	fmt.Printf("%v\n", RunLength("abqq"))
	fmt.Printf("%v\n", RunLength("wwwbbbw"))
	fmt.Printf("%v\n", RunLength("f"))

	return

	for _, value := range []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 26, 27, 28, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73,
		79,
		83, 89, 97, 98, 99, 100} {
		fmt.Printf("%v-%v\n", value, PrimeNew(value))

	}

	return

	background()
	withCancel()
	withTimeOut()
}

//noinspection ALL
func main2() {

	TestX([]string{"r", "d", "u", "l"}, 3)
	return

	fmt.Printf("%v\n", VowelSquare([]string{"aeeekmoo", "kmnouvoo", "frrsfsto"}))
	fmt.Printf("%v\n", VowelSquare([]string{"fghik", "mnoos", "tueae", "ffgei"}))

	fmt.Printf("%v\n", VowelSquare([]string{"aqrst", "ukaei", "fanuo"}))
	return

	return

	fmt.Printf("%v\n", VowelSquare([]string{"gg", "ff"}))
	fmt.Printf("%v\n", VowelSquare([]string{"abcd", "eikr", "oufj"}))

	fmt.Printf("%v\n", PalindromeCreator("annak"))

	return

	fmt.Printf("%v\n", ClosestEnemyII([]string{"0000", "1000", "0002", "0002"}))
	fmt.Printf("%v\n", ClosestEnemyII([]string{"01000", "00020", "00000", "00002", "02002"}))
	fmt.Printf("%v\n", ClosestEnemyII([]string{"00000", "10020", "00000", "00002", "02002"}))
	fmt.Printf("%v\n", ClosestEnemyII([]string{"000", "100", "200"}))
	fmt.Printf("%v\n", ClosestEnemyII([]string{"0000", "2010", "0000", "2002"}))

	return

	fmt.Printf("%v\n", PalindromeCreator("aajgmaa"))
	fmt.Printf("%v\n", PalindromeCreator("racecar"))
	fmt.Printf("%v\n", PalindromeCreator("aaabaaaj"))
	fmt.Printf("%v\n", PalindromeCreator("annak"))
	fmt.Printf("%v\n", PalindromeCreator("lolkm"))
	fmt.Printf("%v\n", PalindromeCreator("aaaaaa"))
	fmt.Printf("%v\n", PalindromeCreator("mmjmmhmm"))
	fmt.Printf("%v\n", PalindromeCreator("mmop"))
	fmt.Printf("%v\n", PalindromeCreator("kjjjhjjj"))

	fmt.Printf("%v\n", ScaleBalancing([]string{"[13, 4]", "[1, 2, 3, 3, 4]"}))
	fmt.Printf("%v\n", ScaleBalancing([]string{"[5, 9]", "[1, 2, 6, 7]"}))
	fmt.Printf("%v\n", ScaleBalancing([]string{"[3, 4]", "[1, 2, 7, 7]"}))
	fmt.Printf("%v\n", ScaleBalancing([]string{"[13, 4]", "[1, 2, 3, 6, 14]"}))
	return

	return

	getMax, people, _ := MoneyDistribution(100, 2)
	who := 0
	for key, value := range people {
		if value == getMax {
			who = key
			break
		}
	}
	fmt.Printf("Max Money Owner : %v - Person Queue Index :%v", getMax, who)

	return

	fmt.Printf("%v\n", MovingMedian([]int{3, 1, 3, 5, 10, 6, 4, 3, 1}))

	//fmt.Printf("%v\n", MovingMedian([]int{5, 2, 4, 6}))
	return

	fmt.Printf("%v\n", TimeDifference([]string{"10:00am", "11:45pm", "5:00am", "12:01am"}))
	fmt.Printf("%v\n", TimeDifference([]string{"1:10pm", "4:40am", "5:00pm"}))

	return

	fmt.Printf("%v\n", TriangleRow(1))
	fmt.Printf("%v\n", TriangleRow(2))
	fmt.Printf("%v\n", TriangleRow(6))

	return

	fmt.Printf("%v\n", PalindromeSwapper("anna"))
	fmt.Printf("%v\n", PalindromeSwapper("ahhhhhhhhh"))
	fmt.Printf("%v\n", PalindromeSwapper("notpossible"))
	fmt.Printf("%v\n", PalindromeSwapper("rcaecar"))

	return

	fmt.Printf("%v\n", CommandLine("empty="))

	return

	fmt.Printf("%v\n", StarRating("0.0"))
	fmt.Printf("%v\n", StarRating("1.00"))
	fmt.Printf("%v\n", StarRating("2.39"))
	fmt.Printf("%v\n", StarRating("1.99"))

	return

	fmt.Printf("%v\n", FizzBuzz(8))
	return

	fmt.Printf("%v\n", StringChanges("MNMNjMa"))
	fmt.Printf("%v\n", StringChanges("MrtyNNgMM"))
	return

	fmt.Printf("%v\n", SerialNumber("11.124.667"))
	return

	var m1 [][]float32 = [][]float32{
		[]float32{1.0, 2.0, 3.0},
		[]float32{4.0, 5.0, 6.0},
	}
	var m2 [][]float32 = [][]float32{
		[]float32{0.5, 0.5},
		[]float32{0.2, 0.8},
		[]float32{0.7, 0.3},
	}

	mRes, err := MatrixMultiply(m1, m2)

	if err != nil {
		fmt.Printf("Error : %v \n", err.Error())
		return
	}
	fmt.Printf("Result : %v \n", mRes)

	return

	fmt.Printf("%v\n", ElementMerger([]int{5, 7, 16, 1, 2}))
	fmt.Printf("%v\n", ElementMerger([]int{1, 1, 1, 2}))

	return

	fmt.Printf("%v\n", StringMerge("abc1*kyoo"))
	return
	fmt.Printf("%v\n", ASCIIConversion("hello world"))
	return

	fmt.Printf("%v\n", HDistance([]string{"10011", "10100"}))
	fmt.Printf("%v\n", HDistance([]string{"abcdef", "defabc"}))

	return

	fmt.Printf("%v\n", GroupTotals([]string{"P:1", "N:1", "Z:1", "P:0", "N:-2", "Z:-1"}))

	return

	fmt.Printf("%v\n", GroupTotals([]string{"X:-1", "Y:1", "X:-4", "B:3", "X:5"}))

	return
	fmt.Printf("%v\n", SnakeCase("cats AND*Dogs-are Awesome"))
	fmt.Printf("%v\n", SnakeCase("a b c d-e-f%g"))
	return

	fmt.Printf("%v\n", GCF([]int{106, 212}))
	fmt.Printf("%v\n", GCF([]int{45, 12}))
	fmt.Printf("%v\n", GCF([]int{1, 6}))
	fmt.Printf("%v\n", GCF([]int{12, 28}))

	return

	fmt.Printf("%v\n", NumberStream("6539923335"))
	return

	fmt.Printf("%v\n", ClosestEnemy([]int{0, 1, 0}))

	fmt.Printf("%v\n", ClosestEnemy([]int{2, 2, 2, 2, 0, 0, 1, 0}))

	fmt.Printf("%v\n", ClosestEnemy([]int{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 2}))
	fmt.Printf("%v\n", ClosestEnemy([]int{2, 0, 0, 0, 2, 2, 1, 0}))

	return

	fmt.Printf("%v\n", SumMultiplier([]int{4, 5, 5, 12}))

	fmt.Printf("%v\n", SumMultiplier([]int{2, 2, 2, 2, 4, 1}))
	fmt.Printf("%v\n", SumMultiplier([]int{1, 1, 2, 10, 3, 1, 12}))

	return

	fmt.Printf("%v\n", CamelCase("cats AND*Dogs-are Awesome"))
	fmt.Printf("%v\n", CamelCase("a b c d-e-f%g"))

	return
	fmt.Printf("%v\n", SimpleEvens(2222220222))

	fmt.Printf("%v\n", SimpleEvens(20864646452))

	return

	fmt.Printf("%v\n", OneDecremented(9876541110))

	return

	fmt.Printf("%v\n", AlphabetSearching("zacxyjbbkfgtbhdaielqrm45pnsowtuv"))

	return

	fmt.Printf("%v\n", StringPeriods("gg"))
	fmt.Printf("%v\n", StringPeriods("abcxabc"))
	fmt.Printf("%v\n", StringPeriods("affedaaffed"))
	fmt.Printf("%v\n", StringPeriods("abcababcababcab"))

	return

	fmt.Printf("%v\n", LargestFour([]int{1, 1, 1, -5}))
	fmt.Printf("%v\n", LargestFour([]int{0, 0, 2, 3, 7, 1}))

	return

	//)(()
	fmt.Printf("%v\n", RemoveBrackets(")(()"))
	fmt.Printf("%v\n", RemoveBrackets("))))(()"))
	fmt.Printf("%v\n", RemoveBrackets("(()))"))
	fmt.Printf("%v\n", RemoveBrackets("(())()((("))
	fmt.Printf("%v\n", RemoveBrackets("(()("))

	return

	fmt.Printf("%v\n", DistinctCharacters("abc123kkmmmm?"))
	fmt.Printf("%v\n", DistinctCharacters("12334bbmma:=6"))
	fmt.Printf("%v\n", DistinctCharacters("eeeemmmmmmmmm1000"))

	return

	fmt.Printf("%v\n", ThreeSum([]int{10, 2, 3, 1, 5, 3, 1, 4, -4, -3, -2}))

	fmt.Printf("%v\n", ThreeSum([]int{12, 3, 1, -5, -4, 7}))

	return

	//fmt.Printf("%v\n", BasicRomanNumerals("IV"))
	//fmt.Printf("%v\n", QuestionsMarks("aa6?9"))
	//fmt.Printf("%v\n",QuestionsMarks("aaaaaaarrrrr??????"))
	//fmt.Printf("%v\n",QuestionsMarks("9???1???9???1???9")) // true
	fmt.Printf("%v\n", QuestionsMarks("9???1???9??1???9"))                 // false
	fmt.Printf("%v\n", QuestionsMarks("5??aaaaaaaaaaaaaaaaaaa?5?5"))       // false
	fmt.Printf("%v\n", QuestionsMarks("mbbv???????????4??????ddsdsdcc9?")) // false
	return

	fmt.Printf("%v\n", QuestionsMarks("acc?7??sss?3rr1??????5"))
	fmt.Printf("%v\n", QuestionsMarks("arrb6???4xxbl5???eee5"))

	return

	fmt.Printf("%v\n", BasicRomanNumerals("XCV"))
	fmt.Printf("%v\n", BasicRomanNumerals("XI"))
	fmt.Printf("%v\n", BasicRomanNumerals("LXVII"))
	fmt.Printf("%v\n", BasicRomanNumerals("IV"))
	fmt.Printf("%v\n", BasicRomanNumerals("XIX"))
	fmt.Printf("%v\n", BasicRomanNumerals("XLVI"))
	fmt.Printf("%v\n", BasicRomanNumerals("XXIV"))

	return

	fmt.Printf("%v\n", PowerSetCount([]int{1, 2, 3, 4, 5, 6}))

	return

	tempy := []int{1, 2, 3}

	values := PowerSet(tempy)
	_ = values
	return

	xx := 1
	fmt.Printf("[] - %v \n", xx)
	xx++
	for u := 0; u < len(tempy); u++ {
		fmt.Printf("%v - %v \n", tempy[u], xx)
		xx++
	}
	for u := 0; u < len(tempy); u++ {
		for y := u + 1; y < len(tempy); y++ {
			if u != y {
				fmt.Printf("%v - %v - %v \n", tempy[u], tempy[y], xx)
				xx++
			}
		}
	}
	fmt.Printf("[1,2,3,4] - %v \n", xx)

	return

	fmt.Printf("%v\n", BitwiseTwo([]string{"100", "000"}))

	fmt.Printf("%v\n", BitwiseTwo([]string{"10100", "11100"}))

	return

	fmt.Printf("%v\n", LargestPair(453857))
	fmt.Printf("%v\n", LargestPair(363223311))
	fmt.Printf("%v\n", LargestPair(91))

	return
	fmt.Printf("%v\n", EvenPairs("3gy41d216"))

	fmt.Printf("%v\n", EvenPairs("5678dddd"))

	fmt.Printf("%v\n", EvenPairs("f09r28i8e67"))

	fmt.Printf("%v\n", EvenPairs("106a"))

	fmt.Printf("%v\n", EvenPairs("128fk9846mf78"))

	return

	fmt.Printf("%v\n", LongestIncreasingSequence([]int{10, 22, 9, 33, 21, 50, 41, 60, 22, 68, 90}))

	return

	fmt.Printf("%v\n", LongestIncreasingSequence([]int{2, 10, 3, 9, 11, 5}))

	return

	fmt.Printf("%v\n", LongestIncreasingSequence([]int{3, 4, -1, 0, 6, 2, 3}))

	return

	fmt.Printf("%v\n", WaveSorting([]int{0, 1, 2, 4, 1, 1, 1}))

	return

	fmt.Printf("%v\n", WaveSorting([]int{1, 1, 1, 1, 5, 2, 5, 1, 1, 3, 5, 6, 8, 3}))

	return

	fmt.Printf("%v\n", WaveSorting([]int{10, 90, 49, 2, 1, 5, 23, 45, 21, 22}))

	return

	fmt.Printf("%v\n", WaveSorting([]int{0, 331, 2, 333, 333, 333, 333333, 3, 8333333, 9999999}))

	return

	fmt.Printf("%v\n", WaveSorting([]int{0, 4, 22, 4, 14, 4, 2}))

	return

	var letters = []string{"A", "B", "C"}
	var numbers = []int{1, 2, 3}
	_ = letters

	sorted := Perm(numbers)
	sort.Slice(sorted, func(i, j int) bool {
		return fmt.Sprintf("%c%c%c", sorted[i][0], sorted[i][1], sorted[i][2]) < fmt.Sprintf("%c%c%c", sorted[j][0],
			sorted[j][1], sorted[j][2])
	})

	for _, value := range sorted {
		fmt.Printf("%v\n", value)
	}

	return

	CallMe(letters, letters, len(letters), 0, false, Fact(len(letters)), 0)

	return

	deleteItem := func(source []int, index int) []int {
		if len(source) == 0 {
			return source
		}
		if index > len(source)-1 || index < 0 {
			return source
		}

		var bef []int = source[:index]
		var aft []int = source[index+1:]

		for _, value := range aft {
			bef = append(bef, value)
		}

		return bef
	}

	itemsArr := []int{1, 3, 2, 3, 2234, 533, 43, 45}
	fmt.Printf("%v\n", itemsArr)

	itemsArr = deleteItem(itemsArr, 3)
	fmt.Printf("%v\n", itemsArr)

	itemsArr = deleteItem(itemsArr, 3)
	fmt.Printf("%v\n", itemsArr)

	itemsArr = deleteItem(itemsArr, 13)
	fmt.Printf("%v\n", itemsArr)

	return

	res := LoveStory(itemsArr, 10)

	fmt.Printf("Result : %v\n", res)

	return

	temp := "1234"

	div := func(text string, count int) (r []string) {
		if len(text) == 0 {
			return
		}
		for u := 0; u < len(text); u++ {
			r = append(r, text[u:u+count])
		}

		return
	}

	fmt.Printf("%v\n", div(temp, 1))

	return

	fmt.Printf("%v\n", ProductDigits(90))
	fmt.Printf("%v\n", ProductDigits(23))
	fmt.Printf("%v\n", ProductDigits(24))

	return
	fmt.Printf("%v\n", NextPalindrome(99))
	fmt.Printf("%v\n", NextPalindrome(9))

	return

	fmt.Printf("%v\n", NonrepeatingCharacter("abcdef"))
	fmt.Printf("%v\n", NonrepeatingCharacter("hello world hi hey"))

	return
	//fmt.Printf("%v\n", BinaryReversal("213"))
	fmt.Printf("%v\n", BinaryReversal("4567"))

	return

	fmt.Printf("%v\n", ArrayMatching([]string{"[5, 2, 3]", "[2, 2, 3, 10, 6]"}))
	fmt.Printf("%v\n", ArrayMatching([]string{"[1, 1, 2]", "[2, 5, 1, 2]"}))

	return
	arr := []int{1, 2, 3, 4}
	var results []int
	for _, perm := range permutations(arr) {
		val := ""
		for _, value := range perm {
			val += fmt.Sprintf("%v", value)
		}
		valx, _ := strconv.Atoi(val)
		results = append(results, valx)
	}
	/*
		sort.SliceStable(results, func(i int, j int) bool {
			return results[i] < results[j]
		})
	*/
	for _, perm := range results {
		fmt.Println(perm)
	}

	return

	fmt.Printf("%v\n", OtherProducts([]int{1, 4, 3}))
	fmt.Printf("%v\n", OtherProducts([]int{3, 1, 2, 6}))

	return

	fmt.Printf("%v\n", BitwiseOne([]string{"00011", "01010"}))

	return

	tempx := []rune("Eralp")

	newOne := ""

	for _, value := range tempx {
		if (rune(value) >= 'A' && rune(value) <= 'Z') {
			newOne += string(rune(value) + 32)
		} else if (rune(value) >= 'a' && rune(value) <= 'z') {
			newOne += string(rune(value) - 32)
		}
	}

	fmt.Printf("%v\n", newOne)

	return

	fmt.Printf("%v\n", FoodDistribution([]int{7, 5, 4, 3, 4, 5, 2, 3, 1, 4, 5}))

	return

	fmt.Printf("%v\n", FoodDistribution([]int{4, 5, 4, 5, 2, 3, 1, 2}))
	fmt.Printf("%v\n", FoodDistribution([]int{4, 2, 3, 2, 1}))
	fmt.Printf("%v\n", FoodDistribution([]int{1, 5, 4, 1}))
	fmt.Printf("%v\n", FoodDistribution([]int{3, 2, 1, 0, 4, 1, 0}))
	fmt.Printf("%v\n", FoodDistribution([]int{20, 5, 4, 1}))
	fmt.Printf("%v\n", FoodDistribution([]int{5, 2, 3, 4, 5}))
	fmt.Printf("%v\n", FoodDistribution([]int{5, 3, 1, 2, 1}))
	fmt.Printf("%v\n", FoodDistribution([]int{4, 5, 2, 3, 1, 0}))

	return

	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3}

	fmt.Printf("%v\n", a == b)

	a1 := []int{1, 2, 3}
	b1 := []int{1, 2, 3}

	fmt.Printf("%v - %v \n", EqualSlice(a1, b1), reflect.DeepEqual(a1, b1))

	return

	fmt.Printf("%v\n", TwoSum([]int{6, 2}))
	fmt.Printf("%v\n", TwoSum([]int{100, 90, 90, 90, 90, 11}))

	fmt.Printf("%v\n", TwoSum([]int{17, 4, 5, 6, 10, 11, 4, -3, -5, 3, 15, 2, 7}))
	fmt.Printf("%v\n", TwoSum([]int{7, 6, 4, 1, 7, -2, 3, 12}))

	return

	//1a23
	fmt.Printf("%v\n", ThreeNumbers("1a23"))
	fmt.Printf("%v\n", ThreeNumbers("2a3b5 w1o2rl3d g1gg92"))
	fmt.Printf("%v\n", ThreeNumbers("21aa3a ggg4g4g6ggg"))

	return

	fmt.Printf("%v\n", DifferentCases("cats AND*Dogs-are Awesome"))
	fmt.Printf("%v\n", DifferentCases("Daniel LikeS-coding"))
	fmt.Printf("%v\n", DifferentCases("a b c d-e-f%g"))

	return

	fmt.Printf("%v\n", HammingDistance([]string{"10011", "10100"}))
	fmt.Printf("%v\n", HammingDistance([]string{"helloworld", "worldhello"}))
	return

	fmt.Printf("%v\n", Superincreasing([]int{1, 2, 3, 4}))
	fmt.Printf("%v\n", Superincreasing([]int{1, 2, 5, 10}))
	fmt.Printf("%v\n", Superincreasing([]int{1, 3, 6, 13, 54}))

	return

	fmt.Printf("%v\n", OverlappingRanges([]int{5, 11, 1, 5, 1}))
	fmt.Printf("%v\n", OverlappingRanges([]int{1, 8, 2, 4, 4}))
	fmt.Printf("%v\n", OverlappingRanges([]int{4, 10, 2, 6, 3}))

	return

	fmt.Printf("%v\n", OffLineMinimum([]string{"5", "4", "6", "E", "1", "7", "E", "E", "3", "2"}))
	fmt.Printf("%v\n", OffLineMinimum([]string{"1", "2", "E", "E", "3"}))
	fmt.Printf("%v\n", OffLineMinimum([]string{"4", "E", "1", "E", "2", "E", "3", "E"}))

	return

	fmt.Printf("%v\n", MultiplicativePersistence(39))
	fmt.Printf("%v\n", MultiplicativePersistence(4))
	fmt.Printf("%v\n", MultiplicativePersistence(25))

	return
	fmt.Printf("%v\n", ChangingSequence([]int{-3, -2, -1, 0, 2}))

	return

	fmt.Printf("%v\n", ChangingSequence([]int{1, 2, 4, 6, 4, 3, 1}))
	fmt.Printf("%v\n", ChangingSequence([]int{-4, -2, 9, 10}))
	fmt.Printf("%v\n", ChangingSequence([]int{5, 4, 3, 2, 10, 11}))

	return

	fmt.Printf("%v\n", MeanMode([]int{10, 10, 10, 2}))
	fmt.Printf("%v\n", MeanMode([]int{7, 7, 12}))
	fmt.Printf("%v\n", MeanMode([]int{5, 3, 3, 3, 1}))
	fmt.Printf("%v\n", MeanMode([]int{1, 2, 3}))
	fmt.Printf("%v\n", MeanMode([]int{4, 4, 4, 6, 2}))
	return

	fmt.Printf("%v\n", MeanMode([]int{2, 5, 11, 8, 4}))
	fmt.Printf("%v\n", MeanMode([]int{2, 5, 11, 8, 4, 30}))
	return

	fmt.Printf("%v\n", DivisionStringified(2, 3))
	fmt.Printf("%v\n", DivisionStringified(6874, 67))
	fmt.Printf("%v\n", DivisionStringified(123456789, 10000))

	return

	fmt.Printf("%v\n", ArithGeo([]int{2, 4, 16, 24}))
	fmt.Printf("%v\n", ArithGeo([]int{2, 6, 18, 54}))

	return

	fmt.Printf("%v\n", NumberAddition("88Hello 3World!"))
	fmt.Printf("%v\n", NumberAddition("55Hello"))
	fmt.Printf("%v\n", NumberAddition("75Number9"))
	fmt.Printf("%v\n", NumberAddition("10 2One Number*1*"))

	return
	fmt.Printf("%v\n", CountingMinutesI("2:03pm-1:39pm"))
	fmt.Printf("%v\n", CountingMinutesI("1:23am-1:08am"))
	fmt.Printf("%v\n", CountingMinutesI("2:08pm-2:00am"))
	fmt.Printf("%v\n", CountingMinutesI("11:00am-2:10pm"))
	fmt.Printf("%v\n", CountingMinutesI("12:31pm-12:34pm"))
	fmt.Printf("%v\n", CountingMinutesI("5:00pm-5:11pm"))

	return

	fmt.Printf("%v\n", CountingMinutesI("1:00pm-11:00am"))
	fmt.Printf("%v\n", CountingMinutesI("1:23am-1:08am"))
	fmt.Printf("%v\n", CountingMinutesI("12:30pm-20:01pm"))
	fmt.Printf("%v\n", CountingMinutesI("12:30pm-12:00am"))

	return

	fmt.Printf("%v\n", ThirdGreatest([]string{"hello", "world", "after", "all"}))
	fmt.Printf("%v\n", ThirdGreatest([]string{"coder", "byte", "code"}))
	fmt.Printf("%v\n", ThirdGreatest([]string{"abc", "defg", "z", "hijk"}))

	return
	fmt.Printf("%v\n", RectangleArea([]string{"(0 0)", "(3 0)", "(0 2)", "(3 2)"}))
	fmt.Printf("%v\n", RectangleArea([]string{"(1 1)", "(1 3)", "(3 1)", "(3 3)"}))

	/*
		// ["(0 0)", "(3 0)", "(0 2)", "(3 2)"]  -> 6 ( width : 3 , height : 2)
		// ["(1 1)", "(1 3)", "(3 1)", "(3 3)"]  -> 4 ( widtg : ? , height : ?)


	*/

	//fmt.Printf("%v\n", ArrayAdditionI([]int{3, 5, -1, 8, 12}))

	// [7, 7, 12, 98, 106] the output should be 12 98
	// [1, 42, 42, 180] the output should be 42 42
	// [4, 90] the output should be 90 4

	//fmt.Printf("%v \n", AdditivePersistence(4))

	//fmt.Printf("%v\n", PowersofTwo(9))
	//fmt.Printf("%v\n", PowersofTwo(120))
	//fmt.Printf("%v\n", PowersofTwo(64))
	//fmt.Printf("%v\n", PowersofTwo(65536))
	//return

	items := []string{
		"Hello-LOL",
		"Sup DUDE!!?",
	}
	for _, value := range items {
		fmt.Printf("%v - %v\n", value, SwapCase(value))
	}

}

func BitwiseOne(strArr []string) string {
	if len(strArr) != 2 {
		return ""
	}
	if len(strArr[0]) != len(strArr[1]) {
		return ""
	}
	xor := func(str1, str2 string) (ret string) {
		ret = "0"
		if str1 == str2 {
			ret = "1"
		}
		return
	}
	_ = xor
	bor := func(str1, str2 string) (ret string) {
		ret = "1"
		if str1 == "0" && str2 == "0" {
			ret = "0"
		}
		return
	}

	result := ""
	runeItem1 := []rune(strArr[0])
	runeItem2 := []rune(strArr[1])
	for u := 0; u < len(runeItem1); u++ {
		result += bor(fmt.Sprintf("%c", runeItem1[u]), fmt.Sprintf("%c", runeItem2[u]))
	}
	return result
}

func ArithGeo(arr []int) string {
	if len(arr) < 3 {
		return "-1"
	}
	first := arr[0]
	for u := 0; u < len(arr); u++ {
		if arr[u] == 0 {
			return "-1"
		}
		if u > 0 {
			if arr[u] == first {
				return "-1"
			}
		}
	}
	exit := "-1"
	isArithmetic := true
	factor := arr[1] - arr[0]
	for u := 2; u < len(arr); u++ {
		if arr[u]-arr[u-1] == factor {
			continue
		} else {
			isArithmetic = false
			break
		}
	}
	if isArithmetic {
		return "Arithmetic"
	}
	isGeometric := true
	factor = arr[1] / arr[0]
	for u := 2; u < len(arr); u++ {
		if arr[u]/arr[u-1] == factor {
			continue
		} else {
			isGeometric = false
			break
		}
	}
	if isGeometric {
		return "Geometric"
	}
	return exit
}

func NumberAddition(str string) string {
	if len(str) < 2 {
		return "0"
	}
	isNumber := func(val rune) bool {
		if int(val) >= 48 && int(val) <= 57 {
			return true
		}
		return false
	}

	var numbers []int
	lastValue := ""
	isBefDetected := false
	for _, value := range str {
		if isNumber(rune(value)) {
			if isBefDetected {
				lastValue += string(value)
				continue
			} else {
				lastValue = string(value)
				isBefDetected = true
			}
		} else {
			if lastValue != "" {
				partialValue, _ := strconv.Atoi(lastValue)
				numbers = append(numbers, partialValue)
			}
			lastValue = ""
			isBefDetected = false
		}
	}
	if lastValue != "" {
		partialValue, _ := strconv.Atoi(lastValue)
		numbers = append(numbers, partialValue)
		lastValue = ""
	}
	total := 0
	for _, value := range numbers {
		total += value
	}
	return fmt.Sprintf("%d", total)
}

var punchItems = []string{
	" ",
	"-",
	"%",
	"’'",
	"()[]{}<>",
	":",
	",",
	"+",
	"‒–—―",
	"…",
	"!",
	".",
	"«»",
	"-‐",
	"?",
	"‘’“”",
	";",
	"/",
	"⁄",
	"␠",
	"·",
	"&",
	"@",
	"*",
	"\\",
	"•",
	"^",
	"¤¢$€£¥₩₪",
	"†‡",
	"°",
	"¡",
	"¿",
	"¬",
	"#",
	"№",
	"%‰‱",
	"¶",
	"′",
	"§",
	"~",
	"¨",
	"_",
	"|¦",
	"⁂",
	"☞",
	"∴",
	"‽",
	"※"}

var numbers = []string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
}

func DashInsert(str string) string {
	if len(str) == 0 {
		return ""
	}
	exit := ""
	checkNumber := func(char string) (int, bool) {
		val, err := strconv.Atoi(char)
		if err != nil {
			return 0, false
		}
		if val >= 0 && val <= 9 {
			return val, true
		}
		return 0, false
	}

	oddDetected := false
	firstValue := ""
	for index, value := range str {
		if val, isNumber := checkNumber(string(value)); isNumber {
			if !oddDetected && val%2 != 0 {
				firstValue = fmt.Sprintf("%d", val)
				oddDetected = true
				if index == len(str)-1 {
					exit += firstValue
				}
				continue
			}
			if oddDetected && val%2 != 0 {
				exit += firstValue + "-" + fmt.Sprintf("%d", val)
				firstValue = ""
				continue
			}
			if val%2 == 0 {
				if oddDetected {
					exit += firstValue + fmt.Sprintf("%d", val)
				} else {
					exit += fmt.Sprintf("%d", val)
				}
				oddDetected = false
				firstValue = ""
			}
		} else {
			oddDetected = false
		}
	}

	return exit
}

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

//    x y      x y      x y      x y
// ["(0 0)", "(3 0)", "(0 2)", "(3 2)"]  -> 6 ( width : 3 , height : 2)
// ["(1 1)", "(1 3)", "(3 1)", "(3 3)"]  -> 4 ( widtg : ? , height : ?)
func RectangleArea(strArr []string) string {
	if len(strArr) != 4 {
		return "0"
	}
	var points []*Point
	for _, value := range strArr {
		text := strings.Replace(value, "(", "", -1)
		text = strings.Replace(text, ")", "", -1)
		values := strings.Split(text, " ")
		if len(values) != 2 {
			continue
		}
		x, errX := strconv.Atoi(values[0])
		if errX != nil {
			continue
		}
		y, errY := strconv.Atoi(values[1])
		if errY != nil {
			continue
		}
		points = append(points, NewPoint(x, y))
	}
	if len(points) != 4 {
		return "0"
	}
	//sort Points
	for u := 0; u < len(points); u++ {
		for y := 0; y < len(points); y++ {
			if points[u].X > points[y].X {
				temp := points[u]
				points[u] = points[y]
				points[y] = temp
			}
		}
	}
	if len(points) != 0 {
		w := points[len(points)-1].X - points[0].X
		h := points[1].Y - points[0].Y
		return fmt.Sprintf("%v", math.Abs(float64(w*h)))
	} else {
		return "0"
	}
}

func ThirdGreatest(strArr []string) string {
	if len(strArr) < 3 {
		return ""
	}
	var items []string
	for _, value := range strArr {
		items = append(items, string(value))
	}
	/*
		sortByLength := func(arr []string) []string {
			for u := 0; u < len(arr); u++ {
				for y := 0; y < len(arr); y++ {
					if len(arr[u]) == len(arr[y]) {
						continue
					} else if len(arr[u]) > len(arr[y]) {
						temp := arr[u]
						arr[u] = arr[y]
						arr[y] = temp
					}
				}
			}
			return arr
		}
	*/

	sort.SliceStable(items, func(i, j int) bool {
		return len(items[i]) > len(items[j])
	})
	//items = sortByLength(items)
	return fmt.Sprintf("%v\n%v", items, items[2])

}

func CountingMinutesI(str string) string {
	if !strings.Contains(str, "am") && !strings.Contains(str, "pm") {
		return ""
	}
	result := ""
	values := strings.Split(str, "-")
	if len(values) != 2 {
		return ""
	}
	start := values[0]

	timeFixer := func(base string, val []string) string {
		for _, in := range val {
			if strings.Contains(base, in) {
				tempBase := strings.Replace(base, in, "", -1)
				tempSplit := strings.Split(tempBase, ":")
				if len(tempSplit) != 2 {
					return ""
				}
				return fmt.Sprintf("%v:%v%v", PadLeft(tempSplit[0], 2, "0"), PadLeft(tempSplit[1], 2, "0"), in)
			}
		}
		return ""
	}
	start = timeFixer(start, []string{"am", "pm"})
	end := values[1]
	end = timeFixer(end, []string{"am", "pm"})

	tS, _ := GetMinuteByTimeInfo(start)
	tE, _ := GetMinuteByTimeInfo(end)
	ReCalc := func() {
		end2 := "23:59pm"
		tE2, _ := GetMinuteByTimeInfo(end2)
		p1 := tE2 - tS + 1
		s1 := "12:00am"
		tS1, _ := GetMinuteByTimeInfo(s1)
		p2 := tE - tS1
		result = fmt.Sprintf("%d", p1+p2)
	}
	if strings.Contains(start, "pm") {
		if strings.Contains(end, "pm") {
			//pm - pm
			if tE-tS < 0 {
				ReCalc()
			} else {
				result = fmt.Sprintf("%d", tE-tS)
			}
		} else {
			//pm ->23:59pm - 00:00 am ->  am
			ReCalc()
		}
	} else {
		if strings.Contains(end, "am") {
			//am - am
			if tE-tS < 0 {
				ReCalc()
			} else {
				result = fmt.Sprintf("%d", tE-tS)
			}
		} else {
			//am ->11:59am - 12:00 pm ->  pm
			end2 := "11:59am"
			tE2, _ := GetMinuteByTimeInfo(end2)
			p1 := tE2 - tS + 1
			s1 := "12:00pm"
			tS1, _ := GetMinuteByTimeInfo(s1)
			p2 := tE - tS1
			result = fmt.Sprintf("%d", p1+p2)
		}
	}
	return result
}

type KeyValue struct {
	Index string
	Value int
}

func NewKeyValue(index string, value int) *KeyValue {
	return &KeyValue{
		Index: index,
		Value: value,
	}
}

func SortedTimeLine() []*KeyValue {
	fullTime := 60 * 12 * 2
	var items []*KeyValue
	whichPart := "am"
	for u := 0; u < fullTime; u++ {
		index := ""
		if u >= 60 {
			h := u / 60
			m := u - (h * 60)
			if u >= 720 {
				whichPart = "pm"
			}
			index = fmt.Sprintf("%v:%v%v",
				PadLeft(strconv.Itoa(h), 2, "0"),
				PadLeft(strconv.Itoa(m), 2, "0"), whichPart)
		} else {
			if u == 0 {
				index = fmt.Sprintf("12:00%v", whichPart)
			} else {
				index = fmt.Sprintf("00:%v%v", PadLeft(strconv.Itoa(u), 2, "0"), whichPart)
			}
		}
		item := NewKeyValue(index, u)
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Value < items[j].Value
	})
	return items
}

//noinspection ALL
func GetMinuteByTimeInfo(time string) (int, error) {
	if !strings.Contains(time, "am") && !strings.Contains(time, "pm") {
		return 0, errors.New("time value error.Please control it.!")
	}
	timeBreaker := 1200
	timeFixer := func(base string, val []string) (string, error) {
		for _, in := range val {
			if strings.Contains(base, in) {
				tempBase := strings.Replace(base, in, "", -1)
				tempBase = strings.Replace(tempBase, ":", "", -1)

				v, sErr := strconv.Atoi(tempBase)
				if sErr != nil {
					return "", errors.New(fmt.Sprintf("%v  - time value error.Please control it.!", sErr.Error()))
				}
				if in == "am" {
					if v > timeBreaker {
						v -= timeBreaker
					}
				} else {
					if v < timeBreaker {
						v += timeBreaker
					}
				}
				vStr := PadLeft(fmt.Sprintf("%v", v), 4, "0")
				return fmt.Sprintf("%c%c:%c%c%v", vStr[0], vStr[1], vStr[2], vStr[3], in), nil
			}
		}
		return "", errors.New("there is no compatible time suffix in the context that being passed!")
	}

	time, _ = timeFixer(time, []string{"am", "pm"})

	timeIndex := 0
	values := SortedTimeLine()
	for _, value := range values {
		if value.Index == time {
			timeIndex = value.Value
			break
		}
	}
	if timeIndex == 0 {
		return timeIndex, errors.New("didn't find time index data!")
	}
	return timeIndex, nil
}

func PadLeft(val string, length int, char string) string {
	if len(val) >= length {
		return val
	}
	add := ""
	for u := 0; u < length-len(val); u++ {
		add += char
	}
	return add + val
}

// [7, 7, 12, 98, 106] the output should be 12 98
// [1, 42, 42, 180] the output should be 42 42
// [4, 90] the output should be 90 4
func SecondGreatLow(arr []int) string {
	if len(arr) < 2 {
		return "0 0"
	} else if len(arr) == 2 {
		if arr[0] == arr[1] {
			return fmt.Sprintf("%d %d", arr[0], arr[1])
		}
	}
	//distinctEliminator
	distinct := func(bold []int) []int {
		var element []int
		isFound := func(items []int, value int) bool {
			for _, item := range items {
				if value == item {
					return true
				}
			}
			return false
		}
		for u := 0; u < len(bold); u++ {
			if !isFound(element, bold[u]) {
				element = append(element, bold[u])
			}
		}
		return element
	}
	tempArr := distinct(arr)
	//sorting
	for u := 0; u < len(tempArr); u++ {
		for y := 0; y < len(tempArr); y++ {
			if tempArr[u] < tempArr[y] {
				temp := tempArr[u]
				tempArr[u] = tempArr[y]
				tempArr[y] = temp
			}
		}
	}
	if len(tempArr) > 2 {
		return fmt.Sprintf("%v %v", tempArr[1], tempArr[len(tempArr)-2])
	} else if len(tempArr) == 2 {
		return fmt.Sprintf("%v %v", tempArr[1], tempArr[0])
	} else {
		return "0 0"
	}
}

//if num is 2718 then your program should return 2 because 2 + 7 + 1 + 8 = 18 and 1 + 8 = 9 and you stop at 9.

func AdditivePersistence(num int) int {
	if len(fmt.Sprintf("%d", num)) == 0 || len(fmt.Sprintf("%d", num)) == 1 {
		return 0
	}
	index := 0
start:
	total := 0
	val := strconv.Itoa(num)
	items := []rune(val)
	numberConvert := func(number int) int {
		return number - 48
	}
	for _, value := range items {
		total += numberConvert(int(value))
	}
	index++
	if len(fmt.Sprintf("%d", total)) != 1 {
		num = total
		goto start

	}
	return index
}

func SwapCase(str string) string {
	if len(str) == 0 {
		return ""
	}

	convertRuneToAscii := func(val rune) string {
		return fmt.Sprintf("%c", val)
	}

	elements := []rune(str)
	var exit []string
	for u := 0; u < len(elements); u++ {
		if elements[u] >= 97 && elements[u] <= 122 {
			elements[u] -= 32
		} else if elements[u] < 97 && elements[u] >= 65 {
			elements[u] += 32
		}
		exit = append(exit, convertRuneToAscii(elements[u]))
	}
	return strings.Join(exit, "")
}

func PowersofTwo(num int) string {
	exit := "true"
do:
	for num != 0 {
		if num%2 != 0 {
			exit = "false"
			break do
		}
		num /= 2
		if num == 1 {
			break do
		}
	}
	return exit
}

func TimeConvert(num int) string {
	if num < 0 {
		return "0:0"
	}
	exit := ""
	hour := num / 60
	if hour == 0 {
		exit = fmt.Sprintf("%d:%d", hour, num)
	} else {
		exit = fmt.Sprintf("%d:%d", hour, num-(hour*60))
	}
	return exit
}

/*
Have the function LetterCountI(str) take the str parameter being passed and return the first word with the greatest number of repeated letters. For example: "Today, is the greatest day ever!" should return greatest because it has 2 e's (and 2 t's) and it comes before ever which also has 2 e's. If there are no words with repeating letters return -1. Words will be separated by spaces.
*/

func LetterCountI(str string) string {
	if len(str) == 0 {
		return "-1"
	}
	words := strings.Split(str, " ")
	if len(words) == 0 {
		return "-1"
	}
	pool := make(map[int]map[string]int, len(words))
	keepSeachFrequency := func(text string) bool {
		for u := 0; u < len(text); u++ {
			for y := u + 1; y < len(text); y++ {
				if text[u] == text[y] {
					return true
				}
			}
		}
		return false
	}
	isFirst := false
	firstVal := ""
	firstIndex := 0
	distinct := func(text string) int {
		totalCount := 0
		for u := 0; u < len(text); u++ {
			for y := u + 1; y < len(text); y++ {
				if text[u] == text[y] {
					totalCount--
				}
			}
			totalCount++
		}
		return totalCount
	}
	convertAsciiInt := func(val int) string {
		return fmt.Sprintf("%c", val)
	}
	for index, value := range words {
		if keepSeachFrequency(value) {
			pool[index] = make(map[string]int, distinct(value))
			for _, char := range value {
				strChar := convertAsciiInt(int(char))
				count, isExist := (pool[index])[strChar]
				if isExist {
					(pool[index])[strChar] = count + 1
					if !isFirst {
						firstVal = strChar
						firstIndex = index
						isFirst = true
					}
				} else {
					(pool[index])[strChar] = 1
				}
			}
		}
	}
	findMaxFrequencyValue := func(result map[int]map[string]int, firstIndex int, firstValue string) int {
		//"Today, is the greatest day ever!"
		if len(result) == 0 || firstValue == "" {
			return -1
		}
		max := (result[firstIndex])[firstVal]
		activeIndex := firstIndex
		for index, item := range result {
			if activeIndex == index {
				continue
			}
			for _, word := range item {
				if word > max {
					max = word
					activeIndex = index
				}
			}
		}
		return activeIndex
	}
	getFirstMax := findMaxFrequencyValue(pool, firstIndex, firstVal)
	if getFirstMax == -1 {
		return "-1"
	}

	return words[getFirstMax]
}

/*
Have the function ArrayAdditionI(arr) take the array of numbers stored in arr and return the string true if any combination of numbers in the array (excluding the largest number) can be added up to equal the largest number in the array, otherwise return the string false. For example: if arr contains [4, 6, 23, 10, 1, 3] the output should return true because 4 + 6 + 10 + 3 = 23. The array will not be empty, will not contain all the same elements, and may contain negative numbers.
*/
func ArrayAdditionI(arr []int) string {
	if len(arr) == 0 {
		return "false"
	}
	index := 0
	for _, value := range arr {
		if value == arr[0] && index != 0 {
			return "false"
		}
		index++
	}
	findMax := func(arr []int) int {
		max := arr[0]
		for u := 1; u < len(arr); u++ {
			if arr[u] > max {
				max = arr[u]
			}
		}
		return max
	}
	maxNumber := findMax(arr)
	getSumWithOutMax := func(max int, arr []int) int {
		sum := 0
		for _, value := range arr {
			if value != max {
				sum += value
			}
		}
		return sum
	}
	return fmt.Sprintf("%v", getSumWithOutMax(maxNumber, arr) >= maxNumber)
}

func ABCheck(str string) string {
	if len(str) < 5 {
		return "false"
	}
	aFound := false
	bFound := false
	index := 0
	for _, value := range str {
		if string(value) == "a" && !bFound {
			aFound = true
			index = 0
			continue
		}
		if string(value) == "b" && !aFound {
			bFound = true
			index = 0
			continue
		}
		if index == 3 && string(value) == "b" && aFound {
			bFound = true
			break
		}
		if index == 3 && string(value) == "a" && bFound {
			aFound = true
			break
		}
		index++
	}
	return fmt.Sprintf("%v", aFound && bFound)

}

//sort strings
//Assume numbers and punctuation symbols will not be included in the string.
func AlphabetSoup(str string) string {
	if len(str) == 0 {
		return ""
	}
	var (
		punchItems = []string{
			" ",
			"’'",
			"()[]{}<>",
			":",
			",",
			"+",
			"‒–—―",
			"…",
			"!",
			".",
			"«»",
			"-‐",
			"?",
			"‘’“”",
			";",
			"/",
			"⁄",
			"␠",
			"·",
			"&",
			"@",
			"*",
			"\\",
			"•",
			"^",
			"¤¢$€£¥₩₪",
			"†‡",
			"°",
			"¡",
			"¿",
			"¬",
			"#",
			"№",
			"%‰‱",
			"¶",
			"′",
			"§",
			"~",
			"¨",
			"_",
			"|¦",
			"⁂",
			"☞",
			"∴",
			"‽",
			"※"}
	)
	var chars []string
	isNumber := func(val rune) bool {
		if int(val) >= 48 && int(val) <= 57 {
			return true
		}
		return false
	}
	isIllegalChar := func(val string) bool {
		for _, value := range punchItems {
			if string(value) == val {
				return true
			}
		}
		return false
	}

	for _, value := range str {
		if !isNumber(rune(value)) && !isIllegalChar(string(value)) {
			chars = append(chars, fmt.Sprintf("%d", value))
		}
	}
	if len(chars) == 0 {
		return ""
	}
	for u := 0; u < len(chars); u++ {
		for y := 0; y < len(chars); y++ {
			first, _ := strconv.Atoi(chars[u])
			next, _ := strconv.Atoi(chars[y])
			if first < next {
				temp := chars[u]
				chars[u] = chars[y]
				chars[y] = temp
			}
		}
	}
	convertAsciiInt := func(val int) string {
		return fmt.Sprintf("%c", val)
	}

	exit := ""
	for _, value := range chars {
		val, _ := strconv.Atoi(string(value))
		exit += convertAsciiInt(val)
	}
	return exit
}

func VowelCount(str string) string {
	if len(str) == 0 {
		return "0"
	}
	listOfVowels := []string{"a", "o", "u", "i", "e"}
	countOfVowels := 0
	for _, value := range str {
		for _, char := range listOfVowels {
			if string(char) == string(value) || strings.ToUpper(string(char)) == string(value) {
				countOfVowels++
			}
		}
	}
	return fmt.Sprintf("%d", countOfVowels)
}

func Palindrome(str string) string {
	if len(str) == 0 {
		return "false"
	}
	cleaner := func(item string) string {
		if len(item) == 0 {
			return item
		}
		val := item
		for _, value := range punchItems {
			val = strings.Replace(val, string(value), "", -1)
		}
		for _, value := range numbers {
			val = strings.Replace(val, string(value), "", -1)
		}
		return val
	}
	reverse := func(item string) string {
		if len(item) == 0 {
			return item
		}
		MAX := len(item) / 2
		var items []string
		for _, value := range item {
			items = append(items, string(value))
		}
		for u := 0; u < MAX; u++ {
			temp := items[u]
			items[u] = items[len(items)-u-1]
			items[len(items)-u-1] = temp
		}
		return strings.Join(items, "")
	}

	flatText := cleaner(str)
	reverseText := cleaner(reverse(str))
	return fmt.Sprintf("%v", flatText == reverseText)
}

func ExOh(str string) string {
	exit := "false"
	if len(str) == 0 {
		return exit
	}
	xCount, oCount := 0, 0
	for u := 0; u < len(str); u++ {
		if string(str[u]) == "x" {
			xCount++
		} else if string(str[u]) == "o" {
			oCount++
		}
	}
	exit = fmt.Sprintf("%v", xCount == oCount)
	return exit
}

func DivTemp(a, b float64) float64 {
	return a / b
}

func WordCount(str string) string {
	if len(str) == 0 {
		return "0"
	}
	var items = strings.Split(str, " ")
	return fmt.Sprintf("%d", len(items))
}

// +d+=3=+s+    --> TRUE
// f++d+        --> FALSE
var splitter = 3

//noinspection ALL
func SimpleSymbols(str string) string {
	if len(str) < 2 {
		return "false"
	}
	for _, value := range str {
		if value >= 97 && value <= 122 {
			if !strings.Contains(str, fmt.Sprintf("+%v+", string(value))) {
				return "false"
			}
		}
	}

	var items []string
	temp := ""
	mod := len(str)%splitter == 0

	isNumber := func(val rune) bool {
		if int(val) >= 48 && int(val) <= 57 {
			return true
		}
		return false
	}

	if mod {
		for u := 0; u <= len(str); u++ {
			if u != 0 && u%splitter == 0 {
				items = append(items, temp)
				temp = ""
			}
			if len(str) != u {
				temp += string(str[u])
			}
		}
	} else {
		max := len(str) - len(str)%splitter
		for u := 0; u <= max; u++ {
			if u != 0 && u%splitter == 0 {
				items = append(items, temp)
				temp = ""
			}
			if len(str) != u {
				temp += string(str[u])
			}
		}
		staging := ""
		for u := max; u < len(str); u++ {
			staging += string(str[u])
		}
		items = append(items, staging)

	}
	exit := "false"
	for _, value := range items {
		if len(string(value)) != splitter {
			return "false"
		} else {
			if isNumber(rune(string(value)[1])) {
				continue
			}
			if string(value)[0] == '+' && string(value)[2] == '+' {
				exit = "true"
			} else {
				exit = "false"
			}
		}
	}
	return exit

}

func SimpleAdding(num int) int {
	if num <= 0 {
		return 0
	}
	sum := 0
	for u := 1; u <= num; u++ {
		sum += u
	}
	return sum

}

func LetterCapitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	ConvertIntAscii := func(val rune) int {
		return int(val)
	}

	var isLower = func(val rune) bool {
		valForInt := ConvertIntAscii(rune(val))
		if valForInt >= 97 && valForInt <= 121 {
			return true
		}
		return false
	}
	var UppForMe = func(val string) string {
		if len(val) == 0 {
			return ""
		}
		exit := ""
		if isLower(rune(val[0])) {
			exit = strings.ToUpper(string(val[0]))
			for y := 1; y < len(val); y++ {
				exit += string(val[y])
			}
		} else {
			exit = val
		}
		return exit
	}
	var items []string
	slitted := strings.Split(str, " ")
	for _, value := range slitted {
		items = append(items, UppForMe(string(value)))
	}
	exit := ""
	for _, value := range items {
		exit += string(value) + " "
	}
	return strings.TrimRight(exit, " ")

}

func FirstReverse(str string) string {
	if len(str) == 0 {
		return ""
	}
	var slice []string
	for _, value := range str {
		slice = append(slice, string(value))
	}
	MAX := len(slice) / 2
	for n := 0; n < MAX; n++ {
		tmp := slice[n]
		slice[n] = slice[len(slice)-1-n]
		slice[len(slice)-1-n] = tmp
	}
	return strings.Join(slice, "")
}

func ConvertAsciiInt(val int) string {
	return fmt.Sprintf("%c", val)
}

func ConvertIntAscii(val rune) int {
	return int(val)
}

func isSpecificVowel(val string, items []rune) bool {
	for _, value := range items {
		if string(value) == val {
			return true
		}
	}
	return false
}

func isIllegalChar(val string) bool {
	for _, value := range punchItems {
		if string(value) == val {
			return true
		}
	}
	return false
}

func isNumber(val rune) bool {
	if int(val) >= 48 && int(val) <= 57 {
		return true
	}
	return false
}

func LetterChanges(str string) string {
	if len(str) == 0 {
		return ""
	}
	temp := ""
	for _, value := range str {
		if !isNumber(rune(value)) && !isIllegalChar(string(value)) {
			val := ConvertIntAscii(rune(value))
			result := ConvertAsciiInt(val + 1)
			if isSpecificVowel(result, vowels) {
				result = strings.ToUpper(result)
			}
			temp += result
		} else {
			temp += string(value)
		}
	}
	// replace each char with the next one!
	// replace the vowel chars with UPPERCASE one!

	return temp
}

func LongestWord(sen string) string {
	if sen == "" {
		return ""
	}

	for _, char := range punchItems {
		if strings.Contains(sen, char) {
			sen = strings.Replace(sen, char, "", -1)
		}
	}

	checkCharList := []string{" "}
	maxLengthWord := ""
	for _, punch := range checkCharList {
		items := strings.Split(sen, punch)
		if len(items) == 0 {
			continue
		}
		for _, word := range items {
			if len(word) > len(maxLengthWord) {
				maxLengthWord = word
			}
		}
	}

	return maxLengthWord
}
