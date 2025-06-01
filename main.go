package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Фильтрация среза

func main() {
	numSlice := EnterNum("Введите числа через запятую: ") // Получаем ввод чисел с консоли
	evenOrNot := Even_or_not_even(numSlice)               // Получаем анонимную функцию для фильтрации
	evenNum, nonEvenNum := evenOrNot()                    // Вызываем анонимную функцию

	evenNumStr := NumToString(evenNum)
	nonEvenNumStr := NumToString(nonEvenNum)

	fmt.Printf("Чётные числа: %s\n", evenNumStr)
	fmt.Printf("Нечётные числа: %s\n", nonEvenNumStr)
}

// Функция конвертирует срез в строку с запятыми
func NumToString(nums []int) string {
	var strSlice []string

	for _, num := range nums {
		strSlice = append(strSlice, strconv.Itoa(num)) // Конвентируем число в строку
	}
	return strings.Join(strSlice, ", ")
}

func Even_or_not_even(numSlice []int) func() ([]int, []int) {
	var evenNum []int    // Срез с чётными числами
	var nonEvenNum []int // Срез с нечётными

	// Анонимная функция
	return func() ([]int, []int) {
		for _, value := range numSlice {
			if value%2 == 0 {
				evenNum = append(evenNum, value)
			} else {
				nonEvenNum = append(nonEvenNum, value)
			}
		}
		return evenNum, nonEvenNum
	}
}

// Ввод числа
func EnterNum(prompt string) []int {
	var text string

	fmt.Print((prompt))
	reader := bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text) // Удаление пробелов сначала и в конце

	if text == "" {
		fmt.Println("Вы ничего не ввели")
	}

	trimText := strings.ReplaceAll(text, " ", "") //Текст без лишних пробелов
	textSlice := strings.Split(trimText, ",")

	var numSlice []int // Срез для хранения чисел

	for _, s := range textSlice {
		num, err := strconv.Atoi(s) // Преобразуем строку в int
		if err != nil {
			fmt.Println("Ошибка приобразования")
			log.Fatal(err)
		}
		numSlice = append(numSlice, num) // Добавляем число в срез
	}

	return numSlice // Возвращаем срез
}
