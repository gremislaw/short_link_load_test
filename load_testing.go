package main

import (
	"fmt"
	"math/rand"
	"time"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func generateRandomURL() string {
	// Инициализация генератора случайных чисел
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Генерируем случайную строку для URL
	randomStr := randStringBytes(10) // длина строки будет 10 символов

	// Возвращаем сгенерированный URL
	return fmt.Sprintf("https://example.com/%s", randomStr)
}

// Генерация случайной строки длиной n символов
func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func myTarget(t *vegeta.Target) error {
	// Генерация случайного URL
	randomStr := randStringBytes(10) // длина строки будет 10 символов
	randomURL := fmt.Sprintf("https://example.com/%s", randomStr)

	// Заполняем Target
	t.Method = "POST"
	t.URL = "http://localhost:8080/shorten" // REST API
	t.Body = []byte(fmt.Sprintf(`"%s"`, randomURL))

	// Возвращаем nil, если ошибок нет
	return nil
}

func main() {
	// Настройка интенсивности запросов
	rate := vegeta.Rate{Freq: 500, Per: time.Second}
	
	// Длительность теста
	duration := 1 * time.Hour

	// Создаём атакующего (attacker)
	attacker := vegeta.NewAttacker()

	// Запускаем нагрузочный тест
	var metrics vegeta.Metrics
	for res := range attacker.Attack(myTarget, rate, duration, "Load Test") {
		metrics.Add(res)
	}
	metrics.Close()

	// Выводим результаты
	fmt.Printf("Requests: %d\n", metrics.Requests)
	fmt.Printf("Success Rate: %.2f%%\n", metrics.Success*100)
	fmt.Printf("Latency (mean): %s\n", metrics.Latencies.Mean)
	fmt.Printf("Latency (95th percentile): %s\n", metrics.Latencies.P95)
	fmt.Printf("Latency (99th percentile): %s\n", metrics.Latencies.P99)
	fmt.Printf("Bytes In (mean): %.2f\n", metrics.BytesIn.Mean)
	fmt.Printf("Bytes Out (mean): %.2f\n", metrics.BytesOut.Mean)
}
