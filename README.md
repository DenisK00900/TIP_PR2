# Практическое занятие №2
по дисциплине «Технологии индустриального программирования».

# Требования
- Go 1.22 или выше
- Git

# Структура проекта
```text
F:.
│   go.mod
│   README.md
│
├───.vscode
│       tasks.json
│
├───bin
│       myapp.exe
│
├───cmd
│   └───myapp
│           main.go
│
├───internal
│   └───app
│       │   app.go
│       │
│       └───handlers
│               ping.go
│
└───utils
        httpjson.go
        logger.go
```

# Скачивание и запуск
```text
git clone https://github.com/DenisK00900/TIP_PR2.git
cd TIP_PR2
go run ./cmd/server
```

# Запросы
```text
curl http://localhost:8080/
curl http://localhost:8080/ping
curl http://localhost:8080/fail
```
(или другой порт, если он был изменён)
