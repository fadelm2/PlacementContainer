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
- Redis : https://github.com/redis/go-redis/v9

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

### Run Web Server
``` bash
go run cmd/web/main.go
```


### TABLE YARD
![Table Yard](TableYard.png)

### TABLE Blocks
![Table Blocks](blocks.png)

### TABLE Yard PLAN
![Table Yard PLAN](TableYardPlan.png)



### POST SUGGESTION
POST http://localhost:8080/api/suggestion
Content-Type: application/json

{
"yard": "YRD1",
"container_number": "ALFI000001",
"container_size": 20,
"container_height": 8.6,
"container_type": "DRY"
}
### POST Placement
POST http://localhost:8080/api/placement
Content-Type: application/json

{
"yard": "YRD1",
"container_number": "ALFI000001",
"block": "3F",
"slot": 1,
"row": 1,
"tier": 2,
"size": 20,
"Height": 8.6,
"Type":   "DRY"
}
### POST PICKUP
POST http://localhost:8080/api/pickup
Content-Type: application/json

{
"yard": "YRD1",
"container_number": "ALFI000001"
}
