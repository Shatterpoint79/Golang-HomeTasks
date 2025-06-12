package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const USDtoEUR float64 = 0.88  //действующие
const USDtoRUB float64 = 78.66 //соотношения

var currencyRates = map[string]map[string]float64{
	"USD": {
		"EUR": USDtoEUR,
		"RUB": USDtoRUB,
	},
	"EUR": {
		"USD": 1 / USDtoRUB,
		"RUB": USDtoRUB / USDtoEUR,
	},
	"RUB": {
		"USD": 1 / USDtoRUB,
		"EUR": USDtoEUR / USDtoRUB,
	},
}

func containsCurrency(dim []string, item string) bool {
	for _, v := range dim {
		if v == item {
			return true
		}
	}
	return false
}

func scanInputCurrency(currencySlice []string) (string, error) {
	var typeCurrency string
	s, err := fmt.Scan(&typeCurrency)
	if err != nil || s != 1 {
		return "", errors.New("Произошла ошибка ввода или введено больше одной валюты")
	}
	if !containsCurrency(currencySlice, typeCurrency) {
		return "", errors.New("Вы ввели значение не из списка")
	}
	return typeCurrency, nil
}

func Exit() bool {
	var input string
	fmt.Print("Выйти? [y/n]: ")
	fmt.Scan(&input)
	return input == "Y" || input == "y"
}

func InputAndCheckAmount() (float64, error) {
	var input string
	fmt.Scan(&input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil || num < 0 {
		return 0.00, errors.New("Это не число или введённое число - отрицательное.")
	}
	return num, nil
}

func removeSelectedCurrency(slice []string, value string) []string {
	result := make([]string, 0, len(slice)-1)
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

func CalculateCurrencyRate(typeStart, typeEnd string) (float64, error) {
	if rateFrom, ok := currencyRates[typeStart]; ok {
		if rateTo, ok := rateFrom[typeEnd]; ok {
			return rateTo, nil
		}
	}
	return 0, errors.New("курс для указанных валют не найден")
}

func printConversionResult(currencyFrom, currencyTo string, amount, result float64) {
	fmt.Println("Исходная валюта:", currencyFrom)
	fmt.Printf("Количество: %.2f \n", amount)
	fmt.Println("Целевая валюта:", currencyTo)
	fmt.Printf("Результат: %.2f \n", result)
}

func main() {

	var currencyEnd = []string{}

	for {

		var currencyStart = []string{"USD", "EUR", "RUB"}

		fmt.Println("Вы можете выбрать исходную валюту ( " + strings.Join(currencyStart[:], ", ") + " )")
		typeCurrencyStart, err := scanInputCurrency(currencyStart)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Println("Вы выбрали валюту: ", typeCurrencyStart)
		}

		currencyEnd = removeSelectedCurrency(currencyStart, typeCurrencyStart)

		fmt.Println("Введите количество валюты: ")
		amountCurrencyStart, err := InputAndCheckAmount()
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Println("Вы ввели: ", amountCurrencyStart)
		}

		fmt.Println("Вы можете выбрать целевую валюту ( " + strings.Join(currencyEnd[:], ", ") + " )")
		typeCurrencyEnd, err := scanInputCurrency(currencyEnd)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Println("Вы выбрали валюту: ", typeCurrencyEnd)
		}

		rate, err := CalculateCurrencyRate(typeCurrencyStart, typeCurrencyEnd)
		if err != nil {
			fmt.Println("Ошибка конвертации:", err)
			continue
		}

		result := amountCurrencyStart * rate

		printConversionResult(typeCurrencyStart, typeCurrencyEnd, amountCurrencyStart, result)

		exit := Exit()
		if exit {
			break
		}
	}
}
