# создаем отдельный билд с именем builder
FROM golang:1.21.1-alpine AS builder
WORKDIR /app

# "из" (.) и "куда" (/app)
COPY . .

RUN go build -o main main.go

RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz


# создаем новый билд уже на основе alpine
FROM alpine:3.13
WORKDIR /app
# из ранее созданного билда берем сбилженный main файл и копируем его
COPY --from=builder /app/main .
# далее так же копируем полученную программу миграции
COPY --from=builder /app/migrate ./migrate
# копируем env файлы
COPY app.env .
# копируем файл
COPY start.sh .
# простая проверка перез запуском контейнеров
COPY wait-for.sh .

# копируем файлы с миграциями
COPY db/migration ./migration

EXPOSE 8080

# эти команды будут переданы как аргументы команде ниже
CMD ["/app/main"]
# передаем в start.sh параметр /app/main, которые и будет подставлен вместо $@
ENTRYPOINT ["/app/start.sh"]