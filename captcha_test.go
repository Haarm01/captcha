package main

import "testing"

func TestCaptcha1111To1115Say1PlusOneToFive(t *testing.T) {

	for i := 1; i <= 5; i++ {
		value := ""
		if i == 1 {
			value = "1+one"
		} else if i == 2 {
			value = "1+two"
		} else if i == 3 {
			value = "1+three"
		} else if i == 4 {
			value = "1+four"
		} else if i == 5 {
			value = "1+five"
		}
		result := Captcha(1, 1, 1, i)
		expected := value
		if result != expected {
			t.Errorf(result)
		}
	}
}

// func TestCaptcha1112Say1PlusTwo(t *testing.T) {
// 	result := Captcha(1, 1, 1, 2)
// 	expected := "1+two"
// 	if result != expected {
// 		t.Errorf(result)
// 	}
// }

// func TestCaptcha1113Say1PlusThree(t *testing.T) {
// 	result := Captcha(1, 1, 1, 3)
// 	expected := "1+three"
// 	if result != expected {
// 		t.Errorf(result)
// 	}
// }

// func TestCaptcha1114Say1PlusFour(t *testing.T) {
// 	result := Captcha(1, 1, 1, 4)
// 	expected := "1+four"
// 	if result != expected {
// 		t.Errorf(result)
// 	}
// }

// func TestCaptcha1115Say1PlusFive(t *testing.T) {
// 	result := Captcha(1, 1, 1, 5)
// 	expected := "1+five"
// 	if result != expected {
// 		t.Errorf(result)
// 	}
// }
