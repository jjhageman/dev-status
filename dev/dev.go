package dev

import (
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/jjhageman/dev-status/db"
	"log"
	//"os"
)

type Dev struct {
	ID        int64
	FirstName string
	LastName  string
	GithubID  string
	Status    string
}

var statuses = [3]string{"available", "looking", "unavailable"}
var dbmap *gorp.DbMap

func init() {
	//url := os.Getenv("HEROKU_POSTGRESQL_COPPER_URL")
	//url += " sslmode=require"
	url := "postgres://jjhageman@localhost:5432/devstatus?sslmode=disable"
	dbmap = db.InitDb(url)
	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Dev{}, "devs").SetKeys(true, "ID")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err := dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
}

func NewDev(first_name string, last_name string, github_id string, status string) (*Dev, error) {
	if !validStatus(status) {
		return nil, fmt.Errorf("invalid status")
	}
	return &Dev{0, first_name, last_name, github_id, status}, nil
}

func (d *Dev) Save() error {
	err := dbmap.Insert(d)
	checkErr(err, "Insert failed")
	return err
}

func All() []*Dev {
	var devs []*Dev
	_, err := dbmap.Select(&devs, "select * from devs order by id")
	checkErr(err, "select all failed")
	return devs
}

func Find(id int64) (*Dev, error) {
	var dev Dev
	err := dbmap.SelectOne(&dev, "select * from devs where id=$1", id)
	//checkErr(err, "select by id failed")
	return &dev, err
}

func validStatus(status string) bool {
	for _, s := range statuses {
		if s == status {
			return true
		}
	}
	return false
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
