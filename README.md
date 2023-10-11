# wb_l0 task

Разработан демонстрационный сервис с использованием go, nats-streaming, postgreSQL.

Работа с бд была произведена через sqlc

## Для запуска миграций

Используется https://github.com/golang-migrate/migrate

```
migrate -path db/migrations -database "databaseUrl" -verbose up 
migrate -path db/migrations -database "databaseUrl" -verbose down
```