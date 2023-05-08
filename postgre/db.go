
package postgre


import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var currentDB *sql.DB
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}

	currentDB = connect()
	presetTables(currentDB)
}

type Wrapper struct {
	db *sql.DB
}

func GetWrapper() Wrapper {
	return Wrapper{currentDB}
}

func GetEnv(key string) string {
	val, found := os.LookupEnv(key)
	if !found {
		log.Fatalln("An env var is missing: ", key)
	}
	return val
}

func getpsqlconn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		GetEnv("dbhost"), GetEnv("dbport"), GetEnv("dbuser"), GetEnv("dbpassword"), GetEnv("dbname"),
	)
}
func connect() *sql.DB {
	db, err := sql.Open("postgres", getpsqlconn())
	if err != nil {
		log.Panicln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("Database connected")
	return db
}

func presetTables(db *sql.DB) {
	initcommands := [...]string{
		`
			CREATE SCHEMA IF NOT EXISTS whatwasit
		;`,

		`
			CREATE TABLE IF NOT EXISTS whatwasit.credentials (
				credential_id SERIAL PRIMARY KEY NOT NULL,
				access_hash CHAR(84) NOT NULL,
				password VARCHAR(256) NOT NULL,
				login VARCHAR(256) NOT NULL
			)
		;`,

	}

	for _, comm := range initcommands {
		_, err := db.Exec(comm)
		if err != nil {
			log.Panicln(err.Error())
		}
	}
}
