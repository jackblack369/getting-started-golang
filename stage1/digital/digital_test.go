package digital

import (
	"fmt"
	"math/bits"
	"testing"
)

func TestFunc(t *testing.T) {
	res := CountBits2(1234)
	fmt.Println("res is ", res)
}

func Persistence(n int) int {
	//convert int to string
	s := fmt.Sprintf("%d", n)
	//multiply each digit in the string
	//if the result is a single digit, return the result, else repeat the process
	//return the number of times the process was repeated
	count := 0
	for len(s) > 1 {
		count++
		product := 1
		for _, v := range s {
			product *= int(v - '0')
		}
		s = fmt.Sprintf("%d", product)
	}
	return count
}

func CountBits(n uint) int {
	count := 0
	for n > 0 {
		/*
			This operation results in a binary number where each bit is set to 1
			if the corresponding bits of both operands are 1, and 0 otherwise.
			int(n & 1) evaluates to 1 if n is odd (least significant bit is 1) and 0 if n is even (least significant bit is 0).
		*/
		count += int(n & 1)
		fmt.Println("n is ", n, " unicode is ", n&1, " count is ", count)
		n >>= 1
	}
	return count
}

func CountBits2(n uint) int {
	return bits.OnesCount(n)
}

func TestDiceScore(t *testing.T) {
	dice := [5]int{4, 4, 5, 4, 4}
	res := DiceScore(dice)
	fmt.Println("res is ", res)
}

func DiceScore(dice [5]int) int {
	/*
			Three 1's => 1000 points
		    Three 6's =>  600 points
		    Three 5's =>  500 points
		    Three 4's =>  400 points
		    Three 3's =>  300 points
		    Three 2's =>  200 points
		    One   1   =>  100 points
		    One   5   =>   50 point

	*/
	/*
		init digital_map_score[int, map[int, score]],
		e.g.
		digital_map_score{1 : {1:100, 2:200, 3:1000, 4:1100, 5:1200, 6:2000}}
	*/
	digital_score_map := map[int]map[int]int{
		1: {1: 100, 2: 200, 3: 1000, 4: 1100, 5: 1200, 6: 2000},
		2: {1: 0, 2: 0, 3: 200, 4: 0, 5: 0, 6: 400},
		3: {1: 0, 2: 0, 3: 300, 4: 0, 5: 0, 6: 600},
		4: {1: 0, 2: 0, 3: 400, 4: 0, 5: 0, 6: 800},
		5: {1: 50, 2: 100, 3: 500, 4: 550, 5: 600, 6: 1000},
		6: {1: 0, 2: 0, 3: 600, 4: 0, 5: 0, 6: 1200},
	}
	/*
		iterate through the dice array and count the number of times each number appears,
		then extract the same number of points from the score array
	*/
	digital_res_map := map[int]int{}
	// iterate through from 1 to 6
	for i := 1; i <= 6; i++ {
		count := 0
		for _, v := range dice {
			if v == i {
				count++
			}
		}
		digital_res_map[i] = digital_score_map[i][count]
	}
	// sum the points
	res := 0
	for _, v := range digital_res_map {
		res += v
	}
	return res
}

func TestPerimeter(t *testing.T) {
	res := Perimeter(5)
	fmt.Println("res is ", res)
}

func Perimeter(n int) int {
	cal_array := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			cal_array[i] = 1
		} else {
			cal_array[i] = cal_array[i-1] + cal_array[i-2]
		}
	}
	fmt.Println("cal_array is ", cal_array)
	sum_res := 0
	for _, v := range cal_array {
		sum_res += v
	}

	perimeter := sum_res * 4
	return perimeter
}

func TestRGB(t *testing.T) {
	res := RGB(254, 253, 252)
	fmt.Println("res is ", res)
}

func RGB(r, g, b int) string {
	//convert the decimal number to hexadecimal
	res := ConvertHex(r) + ConvertHex(g) + ConvertHex(b)
	return res
}

func TestConvertHex(t *testing.T) {
	res := ConvertHex(254)
	fmt.Println("res is ", res)
}

func ConvertHex(n int) string {
	// convert the decimal number to hexadecimal
	// if the number is less than 0, return 0
	// if the number is greater than 255, return 255
	// if the number is less than 16, return 0 + hexadecimal
	// else return hexadecimal
	if n < 0 {
		return "00"
	} else if n > 255 {
		return "FF"
	} else if n < 16 {
		return fmt.Sprintf("0%X", n)
	} else {
		// convert the decimal number to hexadecimal
		return fmt.Sprintf("%X", n)
	}
}

func Str2Array(temp string) []string {
	numArray := []rune(temp)
	res := make([]string, len(numArray))
	for i, v := range numArray {
		res[i] = string(v)
	}
	return res
}

func TestValidBraces(t *testing.T) {
	res := ValidBraces("{}{}()[{}]{{}}(){()}") // {([])}
	fmt.Println("res is ", res)
}

func ValidBraces(str string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range str {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func TestMultiplicationTable(t *testing.T) {
	res := MultiplicationTable(4)
	fmt.Println("res is ", res)
}

func MultiplicationTable(size int) [][]int {
	res := make([][]int, size)
	// locate row num which will used to add the value
	for i := 0; i < size; i++ {
		res[i] = make([]int, size)
		for j := 0; j < size; j++ {
			res[i][j] = (j + 1) * (i + 1)
		}
		fmt.Printf("res[%d] is %v", i, res[i])
	}
	return res
}
