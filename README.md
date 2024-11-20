## REST API for Creating TODO Lists in Go

### read .env file
    go get github.com/joho/godotenv

### Read yaml file 
    go get github.com/spf13/viper

### Postgres :
    docker pull postgres
### Run docker postres image:
    docker run --name=todo-db -e POSTGRES_PASSWORD="12345" -p 5436:5432 -d --rm postgres
### Migrate table:
    migrate create -ext sql -dir ./schema -seq init
### Database table migration:
    migrate -path ./schema -database 'postgres://postgres:12345@localhost:5436/postgres?sslmode=disable' up

### Database connection command:
    docker exec -it conatiner_id /bin/bash
### Postgres command line command  
    psql -U postgres

### Database connection package
    go get github.com/jmoiron/sqlx

### Logging package 
    go get github.com/sirupsen/logrus

### JWT token package 
    github.com/dgrijalva/jwt-go