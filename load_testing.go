package main

import (
	"fmt"
	"time"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func main() {
	// Настроим параметры нагрузки
	target := "http://localhost:8080/shorten" // REST API

	// Инициализация вектора нагрузочного тестирования
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    target,
		Body: []byte(`"https://example.com"`),
	})

	// Настройка интенсивности запросов (например, 100 запросов в секунду)
	rate := vegeta.Rate{Freq: 500, Per: time.Second}
	
	// Длительность теста (например, 10 секунд)
	duration := 10 * time.Second

	// Создаём атакующего (attacker)
	attacker := vegeta.NewAttacker()

	// Запускаем нагрузочный тест
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Load Test") {
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
