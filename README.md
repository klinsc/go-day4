รันทีละบรรทัด
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/caarlos0/env/v11
go install github.com/pressly/goose/v3/cmd/goose@latest
goose down
goose up
go run cmd/student/main.go
