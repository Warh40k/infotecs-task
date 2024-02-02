# Тестовое задания на позицию «Go-разработчик»

Ссылка на репозиторий Github: https://github.com/Warh40k/infotecs-task

## Установка

В корневой директории проекта:
1. Собрать образ приложения в Docker:
```bash
docker build --tag=ewallet:latest .
```
2. Запустить контейнера приложения и базы данных в docker-compose:
```bash
docker-compose up -d
```
3. Установить go-migrate (если отсутствует)
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
3. Применить миграции базы данных:
```bash
migrate -path ./schema/postgresql/ -database 'postgres://dev:dev@localhost:5432/ewallet?sslmode=disable' up
```
PS .env 