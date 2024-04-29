package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "hello world"
	sentence = strings.ReplaceAll(sentence, " ", "") // Удаляем пробелы из предложения

	counts := make(map[rune]int)
	total := 0

	// Подсчитываем количество вхождений каждой буквы в предложении
	for _, char := range sentence {
		counts[char]++
		total++
	}

	// Выводим результаты
	for char, count := range counts {
		percentage := float64(count) / float64(total)
		fmt.Printf("%c - %d %.2f\n", char, count, percentage)
	}
}
