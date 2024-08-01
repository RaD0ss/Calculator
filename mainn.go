package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	roman := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	var input1, input2, symbol, input3 string
	fmt.Println("Введите выражение (например, 2 + 3):")
	fmt.Scanln(&input1, &symbol, &input2, &input3)
	for i := 0; i < 10; i++ {
		if input1 == roman[i] {
			input1 = fmt.Sprint(i + 1)
			for j := 0; j < 10; j++ {
				if input2 == roman[j] {
					input2 = fmt.Sprint(j + 1)
					inputr := input1 + " " + symbol + " " + input2 + " " + input3 + " "
					resultint := calculateR(inputr)
					if resultint != 0 {
						fmt.Println("Результат:", integerToRoman(resultint))
					} else {
						panic("Некорректно !")
					}
					return
				}
			}
		}
	}
	input := input1 + " " + symbol + " " + input2 + " " + input3 + " "
	result := calculateINT(input)
	fmt.Println(result)
}

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func calculateINT(input string) int {

	operands := strings.Fields(input)
	if len(operands) > 3 {
		panic("Неверный формат ввода !")
	}

	a, err1 := strconv.Atoi(operands[0])
	b, err2 := strconv.Atoi(operands[2])
	if err1 != nil || err2 != nil {
		panic("Введены некорректные значения")
	}

	if ((a >= 1) && (a <= 10)) && ((b >= 1) && (b <= 10)) {

		switch operands[1] {
		case "+":
			return a + b
		case "-":
			return a - b
		case "*":
			return a * b
		case "/":
			if b == 0 {
				panic("Деление на 0")
			}
			return a / b
		default:
			panic("Неверная операция !")
		}

	} else {
		panic("Введены некорректные значения")
	}
}

func calculateR(inputr string) int {

	operands := strings.Fields(inputr)
	if len(operands) > 3 {
		panic("Неверный формат !")
	}

	a, err1 := strconv.Atoi(operands[0])
	b, err2 := strconv.Atoi(operands[2])
	if err1 != nil || err2 != nil {
		panic("Введены некорректные значения")
	}

	if ((a >= 1) && (a <= 10)) && ((b >= 1) && (b <= 10)) {

		switch operands[1] {
		case "+":
			return a + b
		case "-":
			return a - b
		case "*":
			return a * b
		case "/":
			if b == 0 {
				panic("Деление на 0")
			}
			return a / b
		default:
			panic("Неверная операция !")
		}

	} else {
		panic("Введены некорректные значения")
	}
}
