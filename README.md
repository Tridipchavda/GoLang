## Task 7 : Authentication Middleware and JWT

This project implements a RESTful API for managing books. It is built using GORM (a Go ORM) for database operations and Mux for routing HTTP requests. The project follows a clean folder structure to maintain code organization.

## EndPoints

| Method    | Route     | Description       | 
| ----------| --------- | ----------------- | 
| POST       | /login    | Login by userId & Password (client will get token as cookie for further authorization purpose)    |
| POST       | /generate    | Generate the JWT Token By getting Username and Password   |

## Example Of Required .env Variables 

```
DATABASE_PORT="5432"
DATABASE_NAME="bacancy"
DATABASE_HOST="localhost"
DATABASE_USER="bacancy"
DATABASE_PASSWORD="password"
```

## Running The Project

1. Make sure you have Go and Postgress installed on your system.
2. Clone the repository.
3. Create a .env file with necessary environment variables.
4. Run `go run main.go` to start the server.
