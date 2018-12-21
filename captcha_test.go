package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCaptcha1111To1115Say1PlusOneToFive(t *testing.T) {

	opeRand := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eigth", "nine"}
	opeRator := [3]string{"+", "-", "*"}
	opd1 := ""
	opt := ""
	opd2 := ""

	for pt := 1; pt <= 2; pt++ {
		for iopd1, vopd1 := range opeRand {
			for ioperator, voperator := range opeRator {
				for iopd2, vopd2 := range opeRand {
					if pt == 1 {
						opd1 = strconv.Itoa(iopd1 + 1)
						opd2 = vopd2
					} else {
						opd1 = vopd1
						opd2 = strconv.Itoa(iopd2 + 1)
					}
					opt = voperator
					result := Captcha(pt, iopd1+1, ioperator+1, iopd2+1)
					expected := opd1 + opt + opd2
					fmt.Println(result)
					fmt.Println(expected)
					if result != expected {
						t.Errorf(result)
					}
				}
			}
		}
	}
}
