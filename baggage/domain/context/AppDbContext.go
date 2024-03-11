package model

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"sync"
)

type AppDbContext struct {
	Client *sql.DB
}

var context *AppDbContext
var once sync.Once

func GetDbContext() *AppDbContext {
	once.Do(func() {
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbAddress := os.Getenv("DB_ADDRESS")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)

		client, err := sql.Open("mysql", dataSource)

		if err != nil {
			panic(err)
		}

		context = &AppDbContext{Client: client}
	})

	return context
}
