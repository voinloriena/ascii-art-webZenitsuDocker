# ---------- Этап сборки ----------
FROM golang:1.25-alpine AS builder

# Устанавливаем git (нужно для go mod tidy)
RUN apk add --no-cache git

# Устанавливаем рабочую директорию
WORKDIR /src

# Копируем go.mod (чтобы зависимости кэшировались)
COPY go.mod ./
RUN go mod tidy || true

# Копируем все файлы проекта
COPY . .

# Собираем бинарник (главный файл — main.go)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /server .

# ---------- Этап запуска ----------
FROM alpine:3.18

# Устанавливаем сертификаты (для HTTPS)
RUN apk add --no-cache ca-certificates

# Создаём рабочую директорию
WORKDIR /app

# Копируем бинарь и необходимые папки из сборочного контейнера
COPY --from=builder /server /app/server
COPY templates /app/templates
COPY static /app/static
COPY back /app/back
COPY banners /app/banners
COPY asciigo /app/asciigo

# Указываем порт, который слушает приложение
EXPOSE 8080

# Запуск приложения
CMD ["./server"]
