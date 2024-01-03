package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type RomanNum struct {
	Value  int
	Symbol string
}

var romanNums = []RomanNum{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{9, "IX"},
	{10, "X"},
	{40, "XL"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
}

type ParsedInputData struct {
	n1, n2, operator string
}

func main() {
	startApp()
}

func startApp() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите: ")
	inputString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Ошибка при считывании: ", err)
		return
	}
	inputString = strings.TrimSuffix(inputString, "\n")
	inputString = strings.ReplaceAll(inputString, " ", "")
	parsedData, err := parseInputData(inputString)

	if err != nil {
		log.Fatal(err)
	}

	n1, _ := convertNumber(parsedData.n1)
	n2, _ := convertNumber(parsedData.n2)

	if !isRoman(n1) && !isRoman(n2) {
		result, err := arithmeticOperation(n1, n2, parsedData.operator)
		if err != nil {
			log.Fatal(err)
		}
		converted, _ := convertNumber(strconv.Itoa(result))
		if result < 1 {
			log.Fatal("ошибка: так как в римской системе нет отрицательных чисел")
		} else {
			fmt.Printf("Результат: %s=%s\n", inputString, converted)
		}

	} else {
		result, err := arithmeticOperation(parsedData.n1, parsedData.n2, parsedData.operator)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Результат: %s=%d\n", inputString, result)
	}
}

func parseInputData(s string) (ParsedInputData, error) {
	operators := [4]string{"+", "-", "*", "/"}
	var operator string
	fmt.Println(s)
	for _, op := range operators {
		if strings.Contains(s, op) {
			operator = op
			break
		}
	}

	operands := strings.Split(s, operator)
	if len(operands) != 2 {
		return ParsedInputData{}, fmt.Errorf("ошибка: Неверный формат данных. пример ввода: a+b")
	}

	n1 := operands[0]
	n2 := operands[1]
	op := operator

	n1IsRoman, n2IsRoman := isRoman(n1), isRoman(n2)
	if (n1IsRoman && !n2IsRoman) || (!n1IsRoman && n2IsRoman) {
		return ParsedInputData{}, fmt.Errorf("ошибка: цифры должны быть одного формата, римские или арабские")
	}

	return ParsedInputData{
		n1,
		n2,
		op,
	}, nil
}

func arithmeticOperation(ns1, ns2, operator string) (int, error) {
	n1, err := strconv.Atoi(ns1)
	if err != nil {
		return 0, fmt.Errorf("ошибка: цифры должны быть в диапазоне от 0 до 10")
	}

	n2, err := strconv.Atoi(ns2)
	if err != nil {
		return 0, fmt.Errorf("ошибка: цифры должны быть в диапазоне от 0 до 10")
	}

	if n1 < 0 || n1 > 10 || n2 < 0 || n2 > 10 {
		return 0, fmt.Errorf("ошибка: цифры должны быть в диапазоне от 0 до 10")
	}

	var result int
	switch operator {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "*":
		result = n1 * n2
	case "/":
		if n2 != 0 {
			result = n1 / n2
		} else {
			return 0, fmt.Errorf("ошибка: Деление на ноль")
		}
	default:
		return 0, fmt.Errorf("ошибка: Неподдерживаемый оператор")
	}
	return result, nil
}

func isRoman(s string) bool {
	validChars := "IVXLCDM"
	for _, char := range s {
		if !strings.ContainsRune(validChars, char) {
			return false
		}
	}
	return true
}

func convertNumber(input string) (string, error) {
	isRoman := isRoman(input)

	if isRoman {
		arabicNumber, err := RomanToArabic(input)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%d", arabicNumber), nil
	}

	arabicNumber, err := strconv.Atoi(input)
	if err != nil {
		return "", fmt.Errorf("неверный формат числа: %v", err)
	}

	romanNumber, err := ArabicToRoman(arabicNumber)
	if err != nil {
		return "", err
	}
	return romanNumber, nil
}

func ArabicToRoman(arabic int) (string, error) {
	if arabic <= 0 || arabic > 3999 {
		return "", fmt.Errorf("недопустимое арабское число для преобразования в римское")
	}

	result := ""
	for i := len(romanNums) - 1; i >= 0; i-- {
		for arabic >= romanNums[i].Value {
			result += romanNums[i].Symbol
			arabic -= romanNums[i].Value
		}
	}
	return result, nil
}

func RomanToArabic(roman string) (int, error) {
	result := 0
	for _, num := range romanNums {
		for strings.HasPrefix(roman, num.Symbol) {
			result += num.Value
			roman = roman[len(num.Symbol):]
		}
	}
	if roman != "" {
		return 0, fmt.Errorf("недопустимая римская цифра: %s", roman)
	}
	return result, nil
}
