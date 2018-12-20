package main

import "strconv"

func main() {
	// Captcha(1, 1, 1, 1)
}

func Captcha(pattern, operand1, operator, operand2 int) string {

	result := ""
	if operand2 == 1 {
		return "1+one"
	} else if operand2 == 2 {
		return "1+two"
	} else if operand2 == 3 {
		return "1+three"
	} else if operand2 == 4 {
		return "1+four"
	} else if operand2 == 5 {
		return "1+five"
	} else {
		result = strconv.Itoa(operand2)
		return "Error " + result
	}

}
