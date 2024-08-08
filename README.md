#### Запуск приложения из папки 15-minutes
- go run main.go -port 4562

#### Запуск приложения из папки aka-microservice
1) cp .env.example .env
2) go run cmd/main.go


#### Для обоих приложений структура URL для получения результата выглядит так:
- /get-items?id=872&id=5234
