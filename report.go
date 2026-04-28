package main

import (
	"fmt"
	"time"
)

// generateReport — функция для генерации и вывода отчёта по результатам
func generateReport(results []Result) {
	fmt.Println("\n=== ОТЧЁТ ПО ОБРАБОТКЕ URL ===")
	fmt.Printf("Всего обработано URL: %d\n", len(results))

	totalDuration := time.Duration(0)
	successful := 0

	fmt.Println("\nДетализация по URL:")
	fmt.Println("------------------")
	for _, result := range results {
		fmt.Printf("URL: %s | Статус: %s | Время: %v\n",
			result.Job.URL, result.Status, result.Duration)

		if result.Status == "обработан" {
			successful++
			totalDuration += result.Duration
		}
	}

	if successful > 0 {
		average := totalDuration / time.Duration(successful)
		fmt.Printf("\nСТАТИСТИКА:\n")
		fmt.Printf("Успешных операций: %d из %d\n", successful, len(results))
		fmt.Printf("Общее время выполнения: %v\n", totalDuration)
		fmt.Printf("Среднее время на запрос: %v\n", average)
	} else {
		fmt.Println("\nНе было успешно обработанных запросов.")
	}
}
