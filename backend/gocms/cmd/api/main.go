package main

import (
	"fmt"
	"os"
)

func main() {
	dbHost := os.Getenv("GOAPI_DB_HOST")
	dbUser := os.Getenv("GOAPI_DB_USER")
	dbPass := os.Getenv("GOAPI_DB_PASSWORD")
	dbName := os.Getenv("GOAPI_DB_NAME")
	fmt.Println(dbHost)
	fmt.Println(dbName)
	fmt.Println(dbPass)
	fmt.Println(dbUser)
}
