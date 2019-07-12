package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
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

func IntSliceToStrSlice(items []int) (str []string) {
	for _, value := range items {
		str = append(str, strconv.Itoa(value))
	}
	return
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

	calc := diffCalc(newArr)
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

/*
func WaveSorting(arr []int) string {
	if len(arr) < 2 {
		return "false"
	}
	result := "true"

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	removeItem := func(slice []int, i int) []int {
		if len(slice) == 0 {
			return slice
		}
		slice[i] = slice[len(slice)-1]
		// We do not need to put s[i] at the end, as it will be discarded anyway
		return slice[:len(slice)-1]
	}

	permGenerator := func(list []int) (result [][]int) {
		if len(list) == 0 {
			return
		}
		for u := 0; u < len(list); u++ {
			var temp []int
			for y := 0; y < len(list); y++ {
				temp = append(temp, list[u])
				if u != y {
					temp = append(temp, list[y])
				}
			}
			result = append(result, temp)
		}
		return
	}

	checkStatus := func(arrNew []int) (result bool) {
		result = true
		for u := 1; u < len(arrNew)-1; u++ {
			if arrNew[u-1] > arrNew[u] && arrNew[u] < arrNew[u+1] {
				fmt.Printf("OK->%v-%v-%v\n", arrNew[u-1], arrNew[u], arrNew[u+1])
				continue
			} else {
				fmt.Printf("WRONG->%v-%v-%v\n", arrNew[u-1], arrNew[u], arrNew[u+1])
				result = false
				break
			}
		}
		return
	}
	_ = checkStatus(arr)

	var list []int
	for {
		if len(arr) >= 3 {
			if arr[1] != arr[0] && arr[2] != arr[1] && arr[2] != arr[0] {
				list = append(list, arr[1], arr[0], arr[2])
				arr = removeItem(arr, 0)
				arr = removeItem(arr, 0)
				arr = removeItem(arr, 0)
			} else {
				result = "false"
				break
			}
		} else {
			if len(arr) == 0 {
				break
			} else {
				//compare the LAST arr item with the last list[n] item over even / odd mode
			}
		}
	}

	return result
}

*/

func Gen(list []int) (results [][]int) {
	if len(list) == 0 {
		return
	}
	var fact func(int) int
	fact = func(a int) int {
		if a == 0 {
			return 1
		}
		return a * fact(a-1)
	}
	var remove func([]int, int) []int
	var find func(int, []int) (int, error)

	remove = func(slice []int, i int) []int {
		if len(slice) == 0 {
			return slice
		}
		slice[i] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}

	find = func(index int, unused []int) (r int, e error) {
		if index < len(unused) {
			r = unused[index]
			e = nil
		} else {
			e = errors.New("no item")
		}
		return
	}

	total := fact(len(list))
	for u := 0; u < total; u++ {
		results = append(results, []int{})
	}
	fix := fact(len(list)) / len(list)
	m := 0
	_ = remove
	for u := 0; u < len(list); u++ {
		var v []int

		for x := 0; x < len(list); x++ {
			if u != x {
				v = append(v, list[x])
			}
		}

		for n := 0; n < fix; n++ {
			results[m] = append(results[m], list[u])
			c := len(v)
			for c != 0 {
				item, err := find(n, v)
				if err == nil {
					results[m] = append(results[m], item)
					v = remove(v, n)
				}
				c--
			}
			m++

			for x := 0; x < len(list); x++ {
				if u != x {
					v = append(v, list[x])
				}
			}

		}

		//m++
	}

	return
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




//noinspection ALL
func main() {
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

	xz := Gen([]int{1, 2, 3, 4})
	for _, value := range xz {
		fmt.Printf("%v\n", value)
	}
	return

	//fmt.Printf("%v\n", WaveSorting([]int{0, 4, 22, 4, 14, 4, 2}))
	return

	//fmt.Printf("%v\n", WaveSorting([]int{0, 1, 2, 4, 1, 1, 1}))

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
