# Базовый образ Go
FROM golang:1.23.3

# Создаем рабочую директорию
WORKDIR /app

# Копируем исходный код и зависимости
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем приложение
RUN go build -o todo-api .

# Открываем порт 3000
EXPOSE 3000

# Запускаем приложение
CMD ["./todo-api"]