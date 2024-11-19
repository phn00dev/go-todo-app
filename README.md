# Go todo app

## read .env file
      go get github.com/joho/godotenv

## Read yaml file 
        go get github.com/spf13/viper

## Postgres :
        docker pull postgres
## Run docker postres image:
        docker run --name=todo-db -e POSTGRES_PASSWORD="12345" -p 5436:5432 -d --rm postgres
## Migrate table:
        migrate create -ext sql -dir ./schema -seq init
## Database table migration:
         migrate -path ./schema -database 'postgres://postgres:12345@localhost:5436/postgres?sslmode=disable' up

## database view:
        docker exec -it conatiner_id /bin/bash
## postgres command line command  
        psql -U postgres

## database connection package
        go get github.com/jmoiron/sqlx

