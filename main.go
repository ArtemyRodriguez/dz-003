package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		transactions := scanTransaction()
		if len(transactions) == 0 {
			fmt.Println("Ошибка: пустой массив. Заполните массив.")
			continue
		}
		for {

			result := getResult(transactions)
			fmt.Print(result)
			doUseraWantToRepeat := doUseraWantToRepeat()
			if !doUseraWantToRepeat {
				break
			}
		}
		doUserWantToExit := doUserWantToExit()
		if !doUserWantToExit {
			fmt.Println("Программа завершена")
			break
		}
	}
}

func scanTransaction() []float64 {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите транзакции черзе запятую: ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	transactionsStr := strings.Split(input, ",")

	transactions := []float64{}

	for _, numStr := range transactionsStr {
		numStr = strings.TrimSpace(numStr)
		if num, err := strconv.ParseFloat(numStr, 64); err == nil {
			transactions = append(transactions, num)
		}
	}
	return transactions
}

func chooseOperationType() string {
	for {
		var operation string
		fmt.Print("Выберите тип операции:\n >>> SUM - посчитать сумму\n >>> AVG - посчитать среднее\n >>> MED - посчитать медиану\nВвод: ")
		fmt.Scan(&operation)

		operation = strings.ToUpper(operation)

		if operation == "SUM" || operation == "AVG" || operation == "MED" {
			return operation
		} else {
			fmt.Println("Ошибка. Введите SUM, AVG или MED")
		}
	}
}

func CalcuelateSum(transactions []float64) float64 {
	var sum float64
	for _, value := range transactions {
		sum += value
	}
	return sum
}

func CalcuelateAvg(transactions []float64) float64 {
	var sum float64
	for _, value := range transactions {
		sum += value
	}
	avg := sum / float64(len(transactions))
	return avg
}

func CalcuelateMed(transactions []float64) float64 {
	sorted := make([]float64, len(transactions))
	copy(sorted, transactions)

	n := len(sorted)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	median := 0.0
	if n%2 == 0 {
		median = (sorted[n/2-1] + sorted[n/2]) / 2
	} else {
		median = sorted[n/2]
	}
	return median
}

func getResult(transactions []float64) string {
	var result string
	fmt.Printf("Массив транзакций: %.f\n", transactions)
	operation := chooseOperationType()

	if operation == "SUM" {
		sum := CalcuelateSum(transactions)
		result = fmt.Sprintf("Сумма элементов массива: %.2f\n", sum)
	} else if operation == "AVG" {
		avg := CalcuelateAvg(transactions)
		result = fmt.Sprintf("Ср. арифм. элементов массива: %.2f\n", avg)
	} else if operation == "MED" {
		med := CalcuelateMed(transactions)
		result = fmt.Sprintf("Медиана массива: %.2f\n", med)
	}

	return result
}

func doUseraWantToRepeat() bool {
	var doUseraWantToRepeat int
	fmt.Print("Выбрать другую операцию? 1 - Да/люб. др. - НЕТ. Ввод: ")
	fmt.Scan(&doUseraWantToRepeat)
	return doUseraWantToRepeat == 1
}

func doUserWantToExit() bool {
	var doUserWantToExit int
	fmt.Print(">>> Вы хотите выйти? <<<\n 1) 1 - Да\n 2) люб. др. - НЕТ. \nВвод: ")
	fmt.Scan(&doUserWantToExit)
	return doUserWantToExit != 1
}
