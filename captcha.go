package main

import "strconv"

func main() {
	// Captcha(1, 1, 1, 1)
}

func Captcha(pattern, operand1, operator, operand2 int) string {

	// patTern := [2]string{"1", "2"}
	// opeRand1 := [5]int{1, 2, 3, 4, 5}
	opeRator := [3]string{"+", "-", "*"}
	opeRand := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eigth", "nine"}

	op1 := ""
	oper := ""
	op2 := ""
	result := ""

	// for i, v := range opeRand1 {
	// 	if i+1 == operand1 {
	// 		op1 = strconv.Itoa(v)
	// 	}
	// }
	if pattern <= 2 && pattern >= 1 && operand1 <= 9 && operand1 >= 1 && operator <= 3 && operator >= 1 && operand2 <= 9 && operand2 >= 1 {

		for i, v := range opeRator {
			if i+1 == operator {
				oper = v
			}
		}

		for i, v := range opeRand {
			if i+1 == operand1 {
				if pattern == 1 {
					op1 = strconv.Itoa(i + 1)
				} else {
					op1 = v
				}
			}
			if i+1 == operand2 {
				if pattern == 1 {
					op2 = v
				} else {
					op2 = strconv.Itoa(i + 1)
				}
			}
		}
	}

	result = op1 + oper + op2

	return result

	// if operand2 == 1 {
	// 	return "1+one"
	// } else if operand2 == 2 {
	// 	return "1+two"
	// } else if operand2 == 3 {
	// 	return "1+three"
	// } else if operand2 == 4 {
	// 	return "1+four"
	// } else if operand2 == 5 {
	// 	return "1+five"
	// } else {
	// 	result = strconv.Itoa(operand2)
	// 	return "Error " + result
	// }

}
