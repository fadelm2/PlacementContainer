# Placement-Container




## Description
APP Ticket Helpdesk IT
## Tech Stack

- Golang : https://github.com/golang/go
- PostgreSql (Database) : https://www.postgresql.org

## Framework & Library

- GoFiber (HTTP Framework) : https://github.com/gofiber/fiber
- GORM (ORM) : https://github.com/go-gorm/gorm
- ENV (Configuration) : https://github.com/joho/godotenv
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator


## Configuration

All configuration is in `.env` file.

## API Spec

All API Spec is in `api` folder.

## Database Migration

All database migration is in `db/migrations` folder.

### Create Migration

```Shell
migrate create -ext sql -dir db/migrations table_xxxxx
```

### Run Migration

```shell
migrate -database "postgres://admin:123456@localhost:5432/pelabuhan?sslmode=disable" -path db/migrations up
```

### Run Appication


### Run unit test


```bash
go test -v ./test/
```


### Run Web Server
``` bash
go run cmd/web/main.go
```

### migrate 


### TABLE YARD
![Table Yard](TableYard.PNG)


### TABLE Blocks
![Table Blocks](TableBlocks.PNG)

### TABLE Yard PLAN
![Table Yard PLAN](TableYardPlan.PNG)
