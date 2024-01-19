# Starter Golang with GIN

- [x] Gin
- [x] Database
- [x] Goose
- [x] Dockerize
- [x] Migrations
- [x] Cobra

# How To Run

### Run Dev
```bash
make dev
```

# Directory folder

```
starter-go-gin
├── Makefile
├── README.md
├── bin
│   └── app.go
├── cmd
│   ├── migrate.go
│   ├── migrate_down.go
│   ├── root.go
│   └── serve.go
├── config
│   └── database_sql.go
├── constants
│   └── constants.go
├── database
│   └── migrations
│       └── 20240106051149_create_table_users.sql
├── dir.md
├── docker
│   ├── entrypoint.dev.sh
│   └── image.dockerfile
├── docker-compose.dev.yml
├── go.mod
├── go.sum
├── handlers
│   ├── user_handler.go
│   └── user_handler_test.go
├── libs
│   ├── libs.go
│   └── response.go
├── main.go
├── models
│   ├── users.go
│   └── users_test.go
└── routes
    ├── routes.go
    └── user.go

13 directories, 27 files
```