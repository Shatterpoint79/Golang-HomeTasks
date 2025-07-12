package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func checkType(input string) bool {
	validTypes := map[string]bool{
		"AVG": true,
		"SUM": true,
		"MED": true,
	}
	return validTypes[input]
}

func inputType() string {
	fmt.Println("Введите тип операции:AVG/SUM/MED")
	fmt.Println("AVG - для расчёта среднего значения чисел")
	fmt.Println("SUM - для получения суммы введённых чисел")
	fmt.Println("MED - для получения медианного значения введённых чисел")
	fmt.Print("> ")
	var input string

	for {
		fmt.Scanln(&input)
		input = strings.ToUpper(input)

		if input == "" {
			fmt.Println("Пустой ввод, повторите")
			continue
		}

		if checkType(input) {
			return input
		}

		fmt.Println("Некорректная операция")
	}
}

func getValue() []int {
	var str string
	fmt.Println("Введите через запятую целые числа")
	fmt.Scanln(&str)
	str = strings.ReplaceAll(strings.TrimSpace(str), " ", "") //Убрать пробелы в начале, середине и конце
	parts := strings.Split(str, ",")                          // Весь ввод будет через запятую

	var cleanedStr []int //Убрать пустой ввод между запятыми типа 1,,3

	for _, p := range parts {
		if p == "" {
			continue
		}

		num, err := strconv.Atoi(p)
		if err != nil {
			fmt.Println("Ввод не является целым числом")
			continue
		}
		cleanedStr = append(cleanedStr, num)

	}
	return cleanedStr

}

func main() {

	result := switchOperation()
	fmt.Printf("result: %.2f\n", result)
}

func switchOperation() float64 {

	choice := inputType()
	value := getValue()

	switch choice {
	case "AVG":
		return calcAVG(value)
	case "SUM":
		return calcSUM(value)
	case "MED":
		return calcMED(value)
	}

	return 0
}

func calcAVG(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}

	total := 0
	for _, value := range numbers {
		total += value
	}

	return float64(total) / float64(len(numbers))

}

func calcSUM(numbers []int) float64 {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return float64(sum)

}

func calcMED(numbers []int) float64 {

	if len(numbers) == 0 {
		return 0
	}

	sorted := make([]int, 0, len(numbers))

	copy(sorted, numbers)

	sort.Ints(sorted)

	l := len(sorted)
	if l%2 == 0 {
		return float64(sorted[l/2-1]+sorted[l/2]) / 2
	}
	return float64(sorted[l/2])

}
