Веб-приложение для хранения заметок с авторизацией. Backend на Go, фронтенд на чистом HTML/CSS/JS.

## Стек

| Слой | Технология |
|---|---|
| Backend | Go 1.23 |
| База данных | PostgreSQL 16 |
| Роутер | go-chi/chi |
| БД драйвер | jackc/pgx |
| Миграции | golang-migrate |
| Авторизация | JWT (golang-jwt) |
| DI | uber/fx |
| Контейнеризация | Docker + Docker Compose |

## Архитектура

Проект построен на слоистой архитектуре (Layered Architecture):

```
domain/       — модели и интерфейсы сервисов
infrastructure/ — реализация репозиториев (SQL запросы)
application/  — бизнес логика (сервисы)
api/          — HTTP хендлеры, роутер, middleware
di/           — сборка зависимостей через fx
config/       — конфигурация приложения
```

## Запуск через Docker

Склонируй репозиторий:
```bash
git clone https://github.com/SilverName608/go-notes.git
cd go-notes
```

Запусти:
```bash
docker-compose up --build
```

Приложение запустится на `http://localhost:8080`. Миграции применяются автоматически.

## Запуск локально

Требования: Go 1.23+, PostgreSQL 16

Создай `.env` файл:
```
DB_DSN=postgres://postgres:YOUR_PASSWORD@localhost:5432/go_notes?sslmode=disable
HTTP_PORT=8080
JWT_SECRET=your_secret_key
```

Установи golang-migrate:
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Накати миграции:
```bash
make migrate-up
```

Запусти сервер:
```bash
make run
```

## API

### Публичные эндпоинты

| Метод | URL | Описание |
|---|---|---|
| GET | `/api/v1/notes` | Получить все заметки |
| POST | `/api/v1/auth/register` | Регистрация |
| POST | `/api/v1/auth/login` | Вход |

### Защищённые эндпоинты (требуют JWT токен)

| Метод | URL | Описание |
|---|---|---|
| POST | `/api/v1/notes` | Создать заметку |
| GET | `/api/v1/notes/{id}` | Получить заметку по ID |
| PUT | `/api/v1/notes/{id}` | Обновить заметку |
| DELETE | `/api/v1/notes/{id}` | Удалить заметку |

Токен передаётся в заголовке:
```
Authorization: Bearer <token>
```

### Примеры запросов

**Регистрация:**
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"john","email":"john@example.com","password":"password123"}'
```

**Создать заметку:**
```bash
curl -X POST http://localhost:8080/api/v1/notes \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{"title":"Моя заметка","body":"Текст заметки"}'
```

## Makefile команды

```bash
make run          # запустить сервер
make build        # собрать бинарник
make migrate-up   # накатить миграции
make migrate-down # откатить миграции
```
