# Link Storage Service

Сервис позволяет сохранять ссылки, получать их по короткому идентификатору и вести статистику обращений.

## Instalation

Запуск сервиса происходит через `Docker`, поэтому необходимо его наличие на машине

### Prerequisites

Предварительно необходимо подготовить `.env` файл и положить его в корень проекта

Пример `.env` файла:

```
ENABLE_ALPINE_PRIVATE_NETWORKING=true
TZ=Europe/Moscow

DB_CONNECTION=postgresql://test:qwerty@postgres_db:5432/test_db?sslmode=disable
POSTGRES_USER=test
POSTGRES_PASSWORD=qwerty
POSTGRES_DB=test_db
```

В корне проекта есть шаблон `.env` файла: `.env.dist`, который можно взять за основу

### Запуск в Docker
Для запуска сервиса в `Docker` необходимо запустить команду:
```
docker compose build && docker compose up
```

## Usage
После запуска сервиса в `Docker` он станет доступен по адресу http://localhost:8080/

Примеры всех энд-поинтов есть в файле `test.http`

## Technologies

* [Go](https://go.dev/) - Go 1.26
* [Postgresql](https://www.postgresql.org/) - Database Postgres 18