# Лабораторная работа №10: Веб-разработка FastAPI(Python) vs Gin(Go)

## Информация о студенте
- ФИО: Евстигнеев Фёдор Алексеевич
- Группа: 220032-11
- Вариант: №5
- Лабораторная: №10
- Репозиторий: lab10-variant5

## Выполненные задания (Вариант 5)

### Средней сложности
- Задание 1: Создано простое API на Go(Gin) с эндпоинтами `/ping`, `GET /api/users`, `POST /api/users`
- Задание 5: Реализована передача сложных структур данных JSON между сервисами (структура User с вложенным Address)
- Задание 7: Реализован graceful shutdown в обоих сервисах (Go: обработка сигналов + context.WithTimeout, Python: lifespan менеджер)

### Повышенной сложности
- Задание 1: Реализована архитектура gRPC (proto/data.proto + Python клиент-стаб).
- Задание 5: Сервисы развёрнуты в Docker Compose с общей сетью lab10-network, healthcheck и зависимостями запуска


##  Быстрый старт

### Запуск через Docker Compose
```bash
docker compose up --build -d

После запуска сервисы доступны по адресам:
Go Gin API: http://localhost:8080
Python FastAPI: http://localhost:8000
Swagger UI: http://localhost:8000/docs
Примечание: если порт 8080 занят, измените маппинг в docker-compose.yml на 8081:8080 и обращайтесь к http://localhost:8081