package main

import (
	"fmt"
	"strings"
)

var vowels = []rune{'a', 'e', 'i', 'o', 'u'}

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

func main() {


	items := []string{
		"hello world",
		"i ran there",
	}
	for _, value := range items {
		fmt.Printf("%v-%v\n", value, LetterCapitalize(value))
	}

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
