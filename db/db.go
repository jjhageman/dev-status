package db

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	"log"
	//"os"
)

func InitDb(conn string) *gorp.DbMap {
	//url := os.Getenv("HEROKU_POSTGRESQL_COPPER_URL")
	//conn, _ := pq.ParseURL(url)
	//conn := "postgres://jjhageman@localhost:5432/devstatus"
	//conn += "?sslmode=disable"
	// connect to db using standard Go database/sql API
	db, err := sql.Open("postgres", conn)
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
