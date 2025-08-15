package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("'ОБМЕННИК'")
	for {
		moneyInWallet, exchangeOperationType, targetCurrencyType, exchangeCurrencyType := whatDoUserWant()
		result, err := getResult(moneyInWallet, exchangeOperationType)
		if err != nil {
			panic("Невалидные параметры для обмена.")
		}

		printExchangeResult(moneyInWallet, result, exchangeCurrencyType, targetCurrencyType)

		userChoice := doUserWantToRepeat()
		if !userChoice {
			fmt.Print("Программа завершена. Рад был помочь!")
			break
		}
	}
}

func getValidCurrencyInput(prompt string, validOptions ...int) int {
	var input int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Ошибка: введите целое число")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		valid := false
		for _, option := range validOptions {
			if input == option {
				valid = true
				break
			}
		}

		if !valid {
			fmt.Printf("Ошибка: введите одно из следующих значений: %v\n", validOptions)
			continue
		}

		return input
	}
}

func whatDoUserWant() (float64, float64, int, int) {
	//рубли и доллары
	usdToRub := 79.65
	rubToUsd := 1 / 79.65
	//рубли и евро
	eurToRub := 92.74
	rubToEur := 1 / 92.74
	//доллары и евро
	eurToUsd := 1.17
	usdToEur := 1 / 1.17
	var exchangeCurrencyType int
	var targetCurrencyType int
	var exchangeOperationType float64
	var moneyInWallet float64

	//выбор исходной валюты с проверкой
	exchangeCurrencyType = getValidCurrencyInput(
		"Выберите валюту для обмена: 1 - Рубли, 2 - Доллары, 3 - Евро. Ввод: ",
		1, 2, 3,
	)

	//ввод кол-во исходной валюты с проверкой
	moneyInWallet = getValidMoneyInput()

	//если пользователь выбрал исходной валютой рубли
	if exchangeCurrencyType == 1 {
		targetCurrencyType = getValidCurrencyInput(
			"Выберите целевую валюту: 2 - Доллары, 3 - Евро. Ввод: ",
			2, 3,
		)
		switch {
		case targetCurrencyType == 2:
			exchangeOperationType = rubToUsd
		case targetCurrencyType == 3:
			exchangeOperationType = rubToEur
		}

		//если пользователь выбрал исходной валютой доллары
	} else if exchangeCurrencyType == 2 {
		targetCurrencyType = getValidCurrencyInput(
			"Выберите целевую валюту: 1 - рубли, 3 - Евро. Ввод: ",
			1, 3,
		)
		switch {
		case targetCurrencyType == 1:
			exchangeOperationType = usdToRub
		case targetCurrencyType == 3:
			exchangeOperationType = usdToEur
		}

		//если пользователь выбрал исходной валютой евро
	} else if exchangeCurrencyType == 3 {
		targetCurrencyType = getValidCurrencyInput(
			"Выберите целевую валюту: 1 - рубли, 2 - Доллары. Ввод: ",
			1, 2,
		)
		switch {
		case targetCurrencyType == 1:
			exchangeOperationType = eurToRub
		case targetCurrencyType == 2:
			exchangeOperationType = eurToUsd
		}
	}
	return moneyInWallet, exchangeOperationType, targetCurrencyType, exchangeCurrencyType
}

func getValidMoneyInput() float64 {
	var money float64
	for {
		fmt.Print("Введите кол-во валюты: ")
		_, err := fmt.Scan(&money)
		if err != nil || money <= 0 {
			fmt.Println("Ошибка: введите положительное число")
			// Очищаем буфер ввода перед повторной попыткой
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		return money
	}
}

func getResult(moneyInWallet float64, exchangeOperationType float64) (float64, error) {
	if moneyInWallet <= 0 {
		return 0, errors.New("error! No_valid_params_to_exchange")
	}
	result := moneyInWallet * exchangeOperationType
	return result, nil
}

func doUserWantToRepeat() bool {
	var repeatOrNot int
	for {
		fmt.Print("Повторить расчет? 1 - Да/2 - Нет. Ввод: ")
		_, err := fmt.Scan(&repeatOrNot)
		if err != nil || (repeatOrNot != 1 && repeatOrNot != 2) {
			fmt.Println("Ошибка: введите 1 или 2")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		break
	}
	return repeatOrNot == 1
}
func printExchangeResult(moneyInWallet, result float64, exchangeCurrencyType, targetCurrencyType int) {
	//вывод для пользователя
	if exchangeCurrencyType == 1 {
		switch {
		case targetCurrencyType == 2:
			mainResult := fmt.Sprintf("%.2f Руб. = %.2f Долл. США", moneyInWallet, result)
			fmt.Println(mainResult)
		case targetCurrencyType == 3:
			mainResult := fmt.Sprintf("%.2f Руб. = %.2f Евро", moneyInWallet, result)
			fmt.Println(mainResult)
		}
	} else if exchangeCurrencyType == 2 {
		switch {
		case targetCurrencyType == 1:
			mainResult := fmt.Sprintf("%.2f Долл. США = %.2f Руб.", moneyInWallet, result)
			fmt.Println(mainResult)
		case targetCurrencyType == 3:
			mainResult := fmt.Sprintf("%.2f Долл. США = %.2f Евро.", moneyInWallet, result)
			fmt.Println(mainResult)
		}
	} else if exchangeCurrencyType == 3 {
		switch {
		case targetCurrencyType == 1:
			mainResult := fmt.Sprintf("%.2f Евро = %.2f Руб.", moneyInWallet, result)
			fmt.Println(mainResult)
		case targetCurrencyType == 2:
			mainResult := fmt.Sprintf("%.2f Евро = %.2f Долл. США", moneyInWallet, result)
			fmt.Println(mainResult)
		}
	}
}
