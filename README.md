# gologin
Simple Go Login using Fiber , Gorm-Postgres

# how to run 

// ------ skip this step if you already have postgresql ----//

- Install docker
- run docker command : docker run --name mac-postgres -e POSTGRES_PASSWORD='Postgres123!@#' -d -p 5432:5432 postgres 

// --------------------------------------------------------------------//
- go mod init github.com/creamyshit/gologin
- go mod tidy
- go run main.go
