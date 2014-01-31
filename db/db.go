package db

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/jjhageman/dev-status/dev"
	_ "github.com/lib/pq"
	"log"
	//"os"
)

func InitDb() *gorp.DbMap {
	//url := os.Getenv("HEROKU_POSTGRESQL_COPPER_URL")
	//conn, _ := pq.ParseURL(url)
	conn := "postgres://jeremyhageman@localhost:5432/devstatus"
	conn += "?sslmode=disable"
	// connect to db using standard Go database/sql API
	db, err := sql.Open("postgres", conn)
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(dev.Dev{}, "devs").SetKeys(true, "ID")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
