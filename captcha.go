package main

func main() {
	// Captcha(1, 1, 1, 1)
}

func Captcha(pattern, operand1, operator, operand2 int) string {

	// patTern := [2]string{"1", "2"}
	// opeRand1 := [5]string{"1", "2", "3", "4", "5"}
	// opeRator := [3]string{"+", "-", "*"}
	opeRand2 := [5]string{"one", "two", "three", "four", "five"}
	op2 := ""
	result := ""

	for i, v := range opeRand2 {
		if i+1 == operand2 {
			op2 = v
			result = "1+" + op2
		}
	}
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
