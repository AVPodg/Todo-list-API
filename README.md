# REST Notes API 📝

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=flat&logo=docker)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-336791?style=flat&logo=postgresql)

RESTful API сервис для управления заметками, написанный на **Go (Golang)** с соблюдением принципов **Чистой Архитектуры (Clean Architecture)**.
Этот проект демонстрирует использование современных практик разработки: внедрение зависимостей (DI), структурированное логирование, работу с базой данных через драйверы и написание тестов.

## 🚀 Особенности

- **Чистая Архитектура:** Четкое разделение слоев (Transport -> Service -> Repository).
- **RESTful API:** Полный набор CRUD операций (Создание, Чтение, Обновление, Удаление).
- **PostgreSQL:** Надежное хранение данных с использованием драйвера `pgx` и библиотеки `sqlx`.
- **Dockerized:** Полная контейнеризация приложения и базы данных через Docker Compose.
- **Dependency Injection:** Модульная структура кода для удобства тестирования и поддержки.
- **Тестирование:**
  - **Unit Tests:** Тестирование бизнес-логики с использованием **Mocks** (`stretchr/testify`).
  - **Integration Tests:** Проверка реального взаимодействия с базой данных.
- **Graceful Shutdown:** Корректное завершение работы при получении сигналов ОС (SIGTERM, SIGINT).
- **Go 1.22 Routing:** Использование нового стандартного роутера `http.ServeMux`.

## 🛠 Технологический стек

- **Язык:** Go (Golang) 1.22+
- **База данных:** PostgreSQL 15
- **Контейнеризация:** Docker, Docker Compose
- **Библиотеки:**
  - `jmoiron/sqlx` - Расширения для работы с SQL
  - `jackc/pgx` - Драйвер PostgreSQL
  - `stretchr/testify` - Ассерты и моки для тестов

## 📂 Структура проекта

```text
.
├── cmd/
│   └── api/            # Точка входа в приложение (main.go)
├── internal/
│   ├── config/         # Логика конфигурации
│   ├── domain/         # Основные сущности (Models)
│   ├── httpapi/        # HTTP Хендлеры (Транспортный уровень)
│   ├── repo/           # Репозиторий базы данных (Уровень данных)
│   └── service/        # Бизнес-логика (Use Cases)
├── docker-compose.yml  # Оркестрация контейнеров
├── Dockerfile          # Описание сборки образа
└── go.mod              # Зависимости
