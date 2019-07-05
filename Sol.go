package main

import (
	"fmt"
	"strconv"
	"strings"
)

var vowels = []rune{'a', 'e', 'i', 'o', 'u'}

func main() {
	//fmt.Printf("%v\n", ArrayAdditionI([]int{5, 7, 16, 1, 2}))
	//fmt.Printf("%v\n", ArrayAdditionI([]int{3, 5, -1, 8, 12}))

	fmt.Printf("%v\n", TimeConvert(126))
	fmt.Printf("%v\n", TimeConvert(45))

	return

	items := []string{
		"567",
		"2129",
		"399047",
		"56730",
		"99946",
	}
	for _, value := range items {
		fmt.Printf("%v - %v\n", value, DashInsert(value))
	}

}

var punchItems = []string{
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
	return "true"

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
