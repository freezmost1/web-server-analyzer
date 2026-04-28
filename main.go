package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Список URL для обработки
	urls := []string{
		"https://example.com/api/1",
		"https://example.com/api/2",
		"https://example.com/api/3",
		"https://example.com/api/4",
		"https://example.com/api/5",
		"https://example.com/api/6",
		"https://example.com/api/7",
		"https://example.com/api/8",
		"https://example.com/api/9",
		"https://example.com/api/10",
	}

	const workerCount = 3 // Количество одновременно работающих воркеров

	fmt.Printf("Запуск обработки %d URL с %d воркерами...\n", len(urls), workerCount)

	// Обработка URL и получение результатов
	results := processURLs(urls, workerCount)

	// Генерация и вывод отчёта
	generateReport(results)
}
