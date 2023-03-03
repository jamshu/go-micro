package main

import (
	"auth/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webport = "80"
var count int64

type Config struct {
	DB *sql.DB
	Data data.Models
}

func main() {

	log.Printf("Starting Authentication Service...\n")

	

	//connect to database
	conn := connectToDB()
	if conn == nil {
		log.Panic("Cant connect to PostgreSQL database")
	}

	app := Config{
		DB: conn ,
		Data: data.New(conn),
	}


	srv := &http.Server{
		Addr: fmt.Sprintf(":%s",webport),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db,nil

}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection,err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not ready yet.....")
			count++
		} else {
			log.Println("Connection established to Postgres")
			return connection
		}
		if count >10 {
			log.Println(err)
			return nil
		}
		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}