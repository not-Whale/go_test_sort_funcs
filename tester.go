package go_test_sort_funcs

import (
	"fmt"
	"os"
	"time"
)

// printError выводит в стандартный поток вывода форматированное сообщение об ошибке
func printError(err error) {
	fmt.Printf("%s%v%s\n", redColor, err.Error(), resetColor)
}

// isArraysEqual проверяет на равенство данные массивы
func isArraysEqual(first, second []int) bool {
	// проверка на совпадение длины
	n1 := len(first)
	n2 := len(second)
	if n1 != n2 {
		return false
	}
	// поэлементное сравнение
	for i := 0; i < n1; i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

// printTestErr выводит сообщение о проваленном из-за неправильного ответа тесте
func printTestErr(inputData, outputData, sortData []int, n int) {
	fmt.Printf("%sTEST №%v - WRONG ANSWER%s\n", redColor, n+1, resetColor)
	fmt.Printf("Input array: %v\n", inputData)
	fmt.Printf("Your result: %v\n", outputData)
	fmt.Printf("Expected: %v\n", sortData)
}

// printTestPass выводит сообщение о пройденном тесте с временем выполнения
func printTestPass(n int, time time.Duration) {
	fmt.Printf("%sTEST №%v - PASSED%v\nTime: %v\n", greenColor, n, resetColor, time)
	fmt.Printf("-----------------------\n")
}

// TestSortFunc запускает тесты данной сортировки и выводит сообщение об ошибке с завершением программы
// в случае неправильного ответа или информацию о времени выполнения каждого теста в противном случае
func TestSortFunc(sortFunc func([]int) []int) {
	// загрузка тестов и ответов на них
	tests := readTests(testFilesNumber)
	answers := readAnswers(testFilesNumber)

	// запуск тестов на входных данных
	n := len(tests)
	for i := 0; i < n; i++ {
		// подсчет времени
		startTime := time.Now()
		outputData := sortFunc(tests[i])
		sortDuration := time.Since(startTime)
		// проверка на правильность результата
		if !isArraysEqual(outputData, answers[i]) {
			printTestErr(tests[i], outputData, answers[i], i)
			os.Exit(1)
		} else {
			printTestPass(i, sortDuration)
		}
	}
	fmt.Printf("Result: OK\n")
}
