package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// URL представляет структуру для хранения информации о URL.
type URL struct {
	Description string
	Tags        string
	Date        time.Time
}

func main() {
	fmt.Println("Программа для добавления URL в список")
	fmt.Println("Для выхода из приложения нажмите q")

	// Создаем карту для хранения URL, где ключом будет имя URL
	urls := make(map[string]URL)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Выберите действие:")
		fmt.Println("a - Добавить URL")
		fmt.Println("l - Вывести список URL")
		fmt.Println("r - Удалить URL")
		fmt.Println("q - Выйти из программы")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		switch input {
		case "a":
			fmt.Println("Введите новую запись в формате <имя_ссылки описание теги>")
			if !scanner.Scan() {
				break
			}
			args := strings.Fields(scanner.Text())
			if len(args) < 3 {
				fmt.Println("Введите правильные аргументы в формате имя_ссылки описание теги")
				continue
			}

			urlName := args[0]
			urlDescription := args[1]
			urlTags := args[2]

			// Создаем новую структуру URL
			newURL := URL{
				Description: urlDescription,
				Tags:        urlTags,
				Date:        time.Now(),
			}

			// Добавляем новую структуру URL в карту
			urls[urlName] = newURL

			fmt.Println("URL успешно добавлен")

		case "l":
			// Вывод списка добавленных URL
			fmt.Println("Список добавленных URL:")
			for name, url := range urls {
				fmt.Printf("Имя: %s\n", name)
				fmt.Printf("Описание: %s\n", url.Description)
				fmt.Printf("Теги: %s\n", url.Tags)
				fmt.Printf("Дата: %s\n\n", url.Date.Format("2006-01-02 15:04:05"))
			}

		case "r":
			fmt.Println("Введите имя ссылки, которую нужно удалить")
			if !scanner.Scan() {
				break
			}
			urlName := scanner.Text()

			// Проверяем, существует ли такой URL
			if _, ok := urls[urlName]; ok {
				delete(urls, urlName)
				fmt.Println("URL успешно удален")
			} else {
				fmt.Println("URL с таким именем не найден")
			}

		case "q":
			fmt.Println("Программа завершена.")
			return

		default:
			fmt.Println("Некорректный ввод.")
		}
	}
}