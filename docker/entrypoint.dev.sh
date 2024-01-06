go install
go mod tidy
go build
cp .env.dev .env
./starter-go-gin migrate
./starter-go-gin serve