Create Migration 
migrate create -ext sql -dir database/migrations create_example_table

Migration Up 
migrate -path database/migrations -database "mysql:s!kid0@123IES@tcp(127.0.0.1:3306)/magic-hands?charset=utf8mb4&parseTime=True&loc=Local" up

Check version Migration 
migrate -path database/migrations -database "mysql:s!kid0@123IES@tcp(127.0.0.1:3306)/magic-hands?charset=utf8mb4&parseTime=True&loc=Local" version