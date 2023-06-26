package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	expression, _ := reader.ReadString('\n')

	result, err := evaluateExpression(expression)
	if err != nil {
		fmt.Println("Output:", err)
		return
	}

	fmt.Println("Output:", result)
}

func evaluateExpression(expression string) (string, error) {
	expression = strings.TrimSpace(expression)

	// Проверяем, содержит ли выражение арабские и римские числа одновременно
	if containsArabic(expression) && containsRoman(expression) {
		return "", fmt.Errorf("Ошибка: использование одновременно разных систем счисления")
	}

	// Проверяем, содержит ли выражение допустимый оператор
	if !strings.ContainsAny(expression, "+-*/") {
		return "", fmt.Errorf("Ошибка: неверный формат математической операции")
	}

	// Разделяем выражение на два числа и оператор
	split := strings.Fields(expression)
	if len(split) != 3 {
		return "", fmt.Errorf("Ошибка: неверный формат математической операции")
	}

	// Парсим числа
	var a, b int
	var err error
	if containsArabic(expression) {
		a, err = strconv.Atoi(split[0])
		if err != nil || a < 1 || a > 10 {
			return "", fmt.Errorf("Ошибка: неверное число: %s", split[0])
		}

		b, err = strconv.Atoi(split[2])
		if err != nil || b < 1 || b > 10 {
			return "", fmt.Errorf("Ошибка: неверное число: %s", split[2])
		}
	} else {
		a, err = romanToArabic(split[0])
		if err != nil {
			return "", fmt.Errorf("Ошибка: неверное число: %s", split[0])
		}

		b, err = romanToArabic(split[2])
		if err != nil {
			return "", fmt.Errorf("Ошибка: неверное число: %s", split[2])
		}
	}

	// Выполняем арифметическую операцию
	var result int
	operator := split[1]

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return "", fmt.Errorf("Ошибка: деление на ноль")
		}
		result = a / b
	}

	// Конвертируем результат в римскую систему счисления, если входные данные были в римской системе
	if containsRoman(expression) {
		roman, err := arabicToRoman(result)
		if err != nil {
			return "", err
		}
		return roman, nil
	}

	return strconv.Itoa(result), nil
}

func containsArabic(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func containsRoman(s string) bool {
	romans := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, r := range romans {
		if strings.Contains(s, r) {
			return true
		}
	}
	return false
}

func romanToArabic(roman string) (int, error) {
	romans := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	arabic, ok := romans[roman]
	if !ok {
		return 0, fmt.Errorf("Ошибка: неверное римское число: %s", roman)
	}

	return arabic, nil
}

func arabicToRoman(arabic int) (string, error) {
	if arabic <= 0 {
		return "", fmt.Errorf("Ошибка: римская система счисления не поддерживает отрицательные и нулевые числа")
	}

	romans := []struct {
		Value  int
		Symbol string
	}{
		{100, "C"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, r := range romans {
		for arabic >= r.Value {
			roman.WriteString(r.Symbol)
			arabic -= r.Value
		}
	}

	return roman.String(), nil
}
