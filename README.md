# Нагрузочное тестирование REST API с использованием Vegeta

Этот проект реализует нагрузочное тестирование для сервиса, который создает короткие ссылки через REST API. Используется [Vegeta](https://github.com/tsenart/vegeta) для проведения нагрузочных тестов, чтобы проверить, как сервис справляется с большим количеством запросов. Тестирование проводим с помощью Go.

## Описание

В проекте настроен нагрузочный тест, который отправляет POST-запросы на эндпоинт сокращения URL. Тест имитирует реальную нагрузку с 500 запросами в секунду на протяжении 60 секунд. Мы проверяем, как сервис обрабатывает множество запросов и отслеживаем такие метрики, как скорость отклика, количество успешных запросов и задержку.

## Как запустить

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/gremislaw/short_link_load_test.git
   ```
2. Скачайте зависимости:
   ```bash
   go mod tidy
   ```
4. Запустите нагрузочный тест:
   ```bash
   go run load_testing.go
   ```

## Нормальная нагрузка: тестирование при ожидаемом уровне использования. 

      RPS: 300
      Duration: 60s
      Requests: 18000
      Success Rate: 100.00%
      Latency (mean): 2.360071ms
      Latency (95th percentile): 4.786207ms
      Latency (99th percentile): 9.024889ms
      Bytes In (mean): 29.00
      Bytes Out (mean): 32.00

## Пиковая нагрузка: тестирование при максимальном уровне использования.

      RPS: 1000
      Duration: 60s
      Requests: 60000
      Success Rate: 100.00%
      Latency (mean): 4.692973ms
      Latency (95th percentile): 11.80991ms
      Latency (99th percentile): 36.041264ms
      Bytes In (mean): 29.00
      Bytes Out (mean): 32.00

## Долгосрочная нагрузка: имитация длительной работы системы под высокой нагрузкой для проверки устойчивости.

      RPS: 500
      Duration: 1 hour
      Requests: 1800000
      Success Rate: 100.00%
      Latency (mean): 4.463109ms
      Latency (95th percentile): 11.110041ms
      Latency (99th percentile): 20.323977ms
      Bytes In (mean): 29.00
      Bytes Out (mean): 32.00

## Дополнительные метрики

Вы можете настроить разные параметры для проверки:

Частота запросов — через параметр Freq в vegeta.Rate.
Длительность теста — через параметр duration.
Тип запроса — настроить Method и Body в vegeta.Target.

---


> **"В мире больших сервисов, как Ozon, важно не только выдержать нагрузку, но и обеспечивать молниеносную реакцию. Тесты — это не просто цифры, это уверенность, что ваш сервис будет готов к любому количеству запросов!"**

