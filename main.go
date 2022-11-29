package main

import (
	"fmt"
	"strings"
)

func main() {
	summa(10, "+", 10)
	calc(10, "+", 10)
}

func calc(s int, operation string, f int) {
	switch operation {
	case "+":
		fmt.Println(s + f)
	case "-":
		fmt.Println(s - f)
	case "*":
		fmt.Println(s * f)
	case "/":
		fmt.Println(s / f)

	}
	if s > 10 {
		panic(s)
	}
	if f > 10 {
		panic(f)
	}
}
func summa(one int, operation string, two int) {
	var result int

	switch operation {
	case "+":
		result = one + two
	case "-":
		result = one - two
	case "*":
		result = one * two
	case "/":
		result = one / two

	}
	var resultText string
	for {
		switch {
		case result > 9:
			n := result / 10
			resultText += strings.Repeat("X", n)
			result = result - 10*n
		}
		if result < 10 {
			break
		}
	}
	switch result {
	case 1:
		resultText += "I"
	case 2:
		resultText += "II"
	case 3:
		resultText += "III"
	case 4:
		resultText += "IV"
	case 5:
		resultText += "V"
	case 6:
		resultText += "VI"
	case 7:
		resultText += "VII"
	case 8:
		resultText += "VIII"
	case 9:
		resultText += "IX"
	case 10:
		resultText += "X"
	}
	if one > 10 {
		panic(one)
	}
	if two > 10 {
		panic(two)
	}
	fmt.Println(resultText)
}
