package main

import "strconv"

func main() {
	// Captcha(1, 1, 1, 1)
}

func Captcha(pattern, operand1, operator, operand2 int) string {

	opeRator := [3]string{"+", "-", "*"}
	opeRand := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eigth", "nine"}
	opd1 := ""
	opt := ""
	opd2 := ""
	result := ""

	if pattern <= 2 && pattern >= 1 && operand1 <= 9 && operand1 >= 1 && operator <= 3 && operator >= 1 && operand2 <= 9 && operand2 >= 1 {
		for i, v := range opeRator {
			if i+1 == operator {
				opt = v
			}
		}
		for i, v := range opeRand {
			if i+1 == operand1 {
				if pattern == 1 {
					opd1 = strconv.Itoa(i + 1)
				} else {
					opd1 = v
				}
			}
			if i+1 == operand2 {
				if pattern == 1 {
					opd2 = v
				} else {
					opd2 = strconv.Itoa(i + 1)
				}
			}
		}
	}
	result = opd1 + opt + opd2
	return result

}
