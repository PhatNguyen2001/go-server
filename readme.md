# magichands-backend

- Gin needs Go >= 1.18

## AdonisJS command

- Documentation: <https://gin-gonic.com/>

### Available commands

```bash
  build                 Compile go code to file exe.

  migration
    migrate create -ext sql -dir database/migrations create_example_table   
    migrate -path database/migrations -database "mysql:s!kid0@123IES@tcp(127.0.0.1:3306)/magic-hands?charset=utf8mb4&parseTime=True&loc=Local" up                                   
    migrate -path database/migrations -database "mysql:s!kid0@123IES@tcp(127.0.0.1:3306)/magic-hands?charset=utf8mb4&parseTime=True&loc=Local" version                         

```

## Git-flow

- Pull latest source code from main
- Create new branch with prefix: feat/fix/refactor: pr-purpose
- Create PR on github and merge

## Setup your env with local databse

```bash
mv env_sample .env
```

Install dependencies

```bash
go mod download
```

Migration and DB seeder

```bash
migrate -path database/migrations -database "mysql:s!kid0@123IES@tcp(127.0.0.1:3306)/magic-hands?charset=utf8mb4&parseTime=True&loc=Local" up  

```

start local server

```bash
go build 
.\go-server.exe
```
