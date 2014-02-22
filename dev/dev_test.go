package dev

import (
	"github.com/coopernurse/gorp"
	"github.com/jjhageman/dev-status/db"
	"log"
	"os"
	"testing"
)

func newDevOrFatal(t *testing.T, first_name string, last_name string, github_id string, status string) *Dev {
	dev, err := NewDev(first_name, last_name, github_id, status)
	if err != nil {
		t.Fatalf("new dev: %v", err)
	}
	return dev
}

func initDbMap() *gorp.DbMap {
	dbmap = db.InitDb("postgres://jjhageman@localhost:5432/devstatus_test?sslmode=disable")
	dbmap.TraceOn("", log.New(os.Stdout, "gorptest: ", log.Lmicroseconds))
	dbmap.AddTableWithName(Dev{}, "devs").SetKeys(true, "ID")
	err := dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	return dbmap
}

func dropAndClose(dbmap *gorp.DbMap) {
	dbmap.DropTablesIfExists()
	dbmap.Db.Close()
}

func TestAll(t *testing.T) {
	dbmap := initDbMap()
	defer dropAndClose(dbmap)

	//var ds []Dev
	//_, err := dbmap.Select(&ds, "select * from devs order by id")
	//checkErr(err, "Select failed")
	//log.Println("All rows:")
	//for x, p := range ds {
	//	log.Printf("    %d: %v\n", x, p)
	//}

	dev1 := newDevOrFatal(t, "Bob", "Jones", "killer_bob", "unavailable")
	dev2 := newDevOrFatal(t, "Jim", "Jones", "killer_bob", "unavailable")

	dev1.Save()
	dev2.Save()

	devs := All()
	if len(devs) != 2 {
		t.Errorf("expected 2 devs, got %v", len(devs))
	}
	if *devs[0] != *dev1 && *devs[1] != *dev1 {
		t.Errorf("missing dev: %v", dev1)
	}
}

func TestFind(t *testing.T) {
	dbmap := initDbMap()
	defer dropAndClose(dbmap)

	dev := newDevOrFatal(t, "Tom", "Jones", "killer_bob", "unavailable")
	dev.Save()

	found := Find(dev.ID)

	if *found != *dev {
		t.Errorf("missing dev: %v", dev)
	}
}

func TestNewDev(t *testing.T) {
	first_name := "John"
	last_name := "Doe"
	github_id := "lambda_joe"
	status := "available"
	dev := newDevOrFatal(t, first_name, last_name, github_id, status)
	if dev.FirstName != first_name {
		t.Errorf("expected first name %q, got %q", first_name, dev.FirstName)
	}
	if dev.LastName != last_name {
		t.Errorf("expected last name %q, got %q", last_name, dev.LastName)
	}
	if dev.GithubID != github_id {
		t.Errorf("expected github id %q, got %q", github_id, dev.GithubID)
	}
	if dev.Status != status {
		t.Errorf("expected status %q, got %q", status, dev.Status)
	}
}

func TestNewDevInvalidStatus(t *testing.T) {
	_, err := NewDev("John", "Doe", "lambda_joe", "bogus")
	if err == nil {
		t.Errorf("expected 'invalid status' error, got nil")
	}
}

func TestSaveDev(t *testing.T) {
	dbmap := initDbMap()
	defer dropAndClose(dbmap)

	dev := newDevOrFatal(t, "Bill", "Jones", "killer_bob", "unavailable")
	dev.Save()

	count, err := dbmap.SelectInt("select count(*) from devs")
	checkErr(err, "select all failed")
	if count != 1 {
		t.Errorf("expected 1 saved dev, got %q", count)
	}
}
