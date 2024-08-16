package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetDB() (*sql.DB, error) {
	//load .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	var (
		dbHost     = os.Getenv("HOST")
		dbPort     = os.Getenv("PORT")
		dbUser     = os.Getenv("USER")
		dbPassword = os.Getenv("PASSWORD")
		dbName     = os.Getenv("DBNAME")
	)

	//connect to db
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", psqlconn)
	//defer db.Close()

	// err = db.Ping()
	// errs.CheckErr("Database isn't connected", err)
	return db, nil
}
