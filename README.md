<a name="readme-top"></a>
# todos v1.0.0
Тестовое задания на стажировку в ecom.tech.
---

##  О проекте



Проект демонстрирует:
- работу с Go Modules
- организацию кода по `cmd / internal`
- написание HTTP-серверов
- обработку запросов и ответов
- логирование через log/slog
- unit-test'ы
- запуск через DockerHub

---

##  Стек технологий

- **Go**
- **net/http**
- **сontext**
- **log/slog**
- **Docker**
- **Git**
- **Postman**

---

##  Структура проекта

```text
ecom-tt/
├── cmd/                # Точки входа в приложение
│   └── main.go         # main package
├── internal/           # Внутренняя логика приложения
│   ├── handlers/       # HTTP-хендлеры
│   └── storage/        # Хранилище данных
│── Dockerfile          # Docker-конфигурация
├── go.mod              # Go modules
├── .gitignore
└── README.md
```

## Установка
Клонировать репозиторий:
```text
git clone https://github.com/zamurabims/ecom-tt.git
```
Перейти в папку проекта:
```text
cd ecomt-tt
```
Загрузить зависимости:
```text
go mod tidy
```
## Локальный запуск
```text
go run ./cmd
```
## Docker
```text
docker build -t ecom-tt .
docker run -p 8080:8080 ecom-tt
```
## Запуск unit-test'ов
```text
go test ./...
```
## Пример POST-запроса ("http://localhost:8080/todos")
```text
{
  "title": "ecom",
  "description": "tech",
  "completed": false
}
```
Ответ:
```
{
    "id": 1,
    "title": "ecom",
    "description": "tech",
    "completed": false
}
```

## Контакты
Nikita Zamura — zamyranikita@gmail.com  
Telegram — [@mysurnameiszamura](http://t.me/mysurnameiszamura)


<p align="right">(<a href="#readme-top">Вернуться к началу</a>)</p>
