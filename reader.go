package go_test_sort_funcs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readIntLines читает строки со значениями типа int из файла и записывает их в возращаемый массив
func readIntLines(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	// чтение количества строк с тестами
	// (структура тестовых данных подразумевает число
	// строк с тестовыми массивами в первой строке файла)
	var n int
	_, err = fmt.Fscan(file, &n)
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	result := make([][]int, n)
	count := 0

	// максимальный размер числа 1e5
	// максимальное количество чисел 1e4
	scanner := bufio.NewScanner(file)
	const maxCapacity int = 1e9
	line := make([]byte, maxCapacity)
	scanner.Buffer(line, maxCapacity)

	// построчное чтение тестовых данных
	for scanner.Scan() {
		tmpLine := scanner.Text()
		strArray := strings.Fields(tmpLine)
		intArray := make([]int, len(strArray))
		// перевод чисел из строчного представления в int
		for i := 0; i < len(strArray); i++ {
			intArray[i], err = strconv.Atoi(strArray[i])
			if err != nil {
				printError(err)
				os.Exit(1)
			}
		}
		result[count] = intArray
		count++
	}

	if err = scanner.Err(); err != nil {
		printError(err)
		os.Exit(1)
	}

	err = file.Close()
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	return result
}

// readTests читает строки со тестовыми массивами с элементами типа int из файла
func readTests(n int) [][]int {
	var result [][]int
	for i := 1; i < n+1; i++ {
		result = append(result, readIntLines("./tests/test"+strconv.Itoa(i)+".txt")...)
	}
	return result
}

// readTests читает строки с массивами с ответами на тесты с элементами типа int из файла
func readAnswers(n int) [][]int {
	var result [][]int
	for i := 1; i < n+1; i++ {
		result = append(result, readIntLines("./tests/answer"+strconv.Itoa(i)+".txt")...)
	}
	return result
}
