go mod tidy
cp .env.dev .env
go install github.com/cosmtrek/air@latest
go run . migrate
air serve