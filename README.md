# YoFio - Backend Golang - Prueba técnica

## Server Setup
1. Init Golang Project
```bash
go mod init github.com/fereicod/test-golang-yofio
```
3. Install GoDotEnv for **.env** files
```bash
    go get github.com/joho/godotenv
```
> We use this library to establish the configuration to the database, it is necessary to create the .env file and define the necessary variables for connection.
4. Install Go-SQL-Driver
```bash
    go get github.com/go-sql-driver/mysql
```
> It's important to know that you will have to import this library to connect your golang server with your mysql database. Normally using like this example.
```golang
    import (
        _ "github.com/go-sql-driver/mysql"
    )
```
5. Run the server
```bash
    go run main.go
```

---

## API Calls
### Investment
These endpoints works for create a new investment, and also for getting the statistics in our database.
#### POST REQUEST
To create a new investment, you have to call to this endpoint **'/credit-assignment'** with a POST REQUEST and provide this data in the body of the request.
```json
    {
        "investment": 3000
    }
```
Using curl will be something like this.
```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{
        "investment": 3000
    }' \
    http://localhost:3690/credit-assignment
```
#### GET REQUEST
To get the statistics that are stored in this database, you will have to call to this endpoint **'/statistics'** with a GET REQUEST., this shows the total investments, successful and unsuccessful, and their average respectively.
Using curl will be something like this.
```bash
    curl http://localhost:3690/statistics
```
## Example
![investment](https://github.com/fereicod/test-golang-yofio/assets/6632994/a57b45a7-6754-4b6f-a1ef-45e037d81db4)
![statistics](https://github.com/fereicod/test-golang-yofio/assets/6632994/02b5114f-f0c7-4a5e-93d0-7e6226d25d97)
