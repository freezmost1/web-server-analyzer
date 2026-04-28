package main

import (
	"math/rand"
	"sync"
	"time"
)

// Job — структура задания для обработки
type Job struct {
	ID  int
	URL string
}

// Result — структура результата обработки
type Result struct {
	Job      Job
	Status   string
	Duration time.Duration
}

// worker — функция-воркер для обработки заданий
func worker(jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		// Имитация HTTP-запроса со случайной задержкой от 100 мс до 1 с
		delay := time.Duration(100+rand.Intn(900)) * time.Millisecond
		time.Sleep(delay)

		result := Result{
			Job:      job,
			Status:   "обработан",
			Duration: delay,
		}

		results <- result
	}
}

// processURLs — основная функция для обработки URL с использованием Worker Pool
func processURLs(urls []string, workerCount int) []Result {
	var wg sync.WaitGroup

	// Каналы для заданий и результатов (буферизованные)
	jobs := make(chan Job, len(urls))
	results := make(chan Result, len(urls))

	// Запуск воркеров
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go worker(jobs, results, &wg)
	}

	// Отправка заданий в канал (Fan-out)
	go func() {
		for i, url := range urls {
			jobs <- Job{ID: i, URL: url}
		}
		close(jobs) // Закрываем канал заданий после отправки всех URL
	}()

	// Сбор результатов и ожидание завершения всех воркеров (Fan-in)
	var allResults []Result
	go func() {
		wg.Wait()      // Ждём завершения всех воркеров
		close(results) // Закрываем канал результатов
	}()

	// Чтение результатов из канала
	for result := range results {
		allResults = append(allResults, result)
	}

	return allResults
}
