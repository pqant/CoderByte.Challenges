package main

import (
	"fmt"
	"strconv"
	"strings"
)

var vowels = []rune{'a', 'e', 'i', 'o', 'u'}

func main() {

	items := []string{
		"coderbyte",
		"hooplah",
	}
	for _, value := range items {
		fmt.Printf("%v-%v\n", value, AlphabetSoup(value))
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
