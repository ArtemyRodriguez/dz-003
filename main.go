package main

import "fmt"

func main() {
	for {
		transactions := scanTransaction()
		if len(transactions) == 0 {
			fmt.Println("Ошибка: пустой массив. Заполните массив.")
			continue
		}
		for {
			var result string
			result = getResult(transactions)
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
	var tr float64
	transactions := []float64{}
	fmt.Print("Введите транзакции(0 - завершить ввод): ")
	for {
		_, err := fmt.Scan(&tr)

		if err != nil || tr == 0 {
			break
		}

		transactions = append(transactions, tr)
	}

	return transactions
}

func chooseOperationType() int {
	for {
		var operationCode int
		fmt.Print("Выберите тип операции:\n >>> 1 - посчитать сумму эл. массива\n >>> 2 - посчитать сред. арифм эл. массива\n >>> 3 - посчитать медиану эл. массива\nВвод: ")
		_, err := fmt.Scan(&operationCode)

		var discard string
		fmt.Scanln(&discard)

		if err != nil {
			fmt.Println("Ошибка. Введите целое число от 1 до 3")
			continue
		} else if operationCode != 1 && operationCode != 2 && operationCode != 3 {
			fmt.Println("Ошибка. Введите целое число от 1 до 3")
			continue
		}
		return operationCode
	}
}
func CalcuelateSum(transactions []float64) float64 {
	transactionsNew := transactions
	var sum float64
	for _, value := range transactionsNew {
		sum += value
	}
	return sum
}
func CalcuelateAvg(transactions []float64) float64 {
	transactionsNew := transactions
	var sum float64
	for _, value := range transactionsNew {
		sum += value
	}
	len := len(transactionsNew)
	avg := 0.0
	avg = sum / float64(len)
	return avg
}
func CalcuelateMed(transactions []float64) float64 {
	sorted := make([]float64, len(transactions))
	copy(sorted, transactions)

	n := len(sorted)
	for i := 1; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	median := 0.0
	if n%2 == 0 {
		//если кол-во элементов массива четное
		median = (sorted[(n/2)-1] + sorted[n/2]) / 2
	} else {
		//если кол-во элементов массива нечетное
		median = sorted[n/2]
	}
	return median
}
func getResult(transactions []float64) string {
	var result string
	fmt.Printf("Массив транзакций: %.f\n", transactions)
	operationCode := chooseOperationType()
	if operationCode == 1 {
		sum := CalcuelateSum(transactions)
		result = fmt.Sprintf("Сумма элементов массива: %.2f\n", sum)
	} else if operationCode == 2 {
		avg := CalcuelateAvg(transactions)
		result = fmt.Sprintf("Ср. арифм. элементов массива: %.2f\n", avg)
	} else if operationCode == 3 {
		med := CalcuelateMed(transactions)
		result = fmt.Sprintf("Медиана массива: %.2f\n", med)
	}
	return result
}
func doUseraWantToRepeat() bool {
	var doUseraWantToRepeat int
	fmt.Print("Выбрать другую операцию? 1 - Да/люб. др. - НЕТ. Ввод: ")
	fmt.Scan(&doUseraWantToRepeat)
	if doUseraWantToRepeat == 1 {
		return true
	} else {
		return false
	}
}
func doUserWantToExit() bool {
	var doUserWantToExit int
	fmt.Print(">>> Вы хотите выйти? <<<\n 1) 1 - Да\n 2) люб. др. - НЕТ. \nВвод: ")
	fmt.Scan(&doUserWantToExit)
	if doUserWantToExit == 1 {
		return false
	} else {
		return true
	}
}
