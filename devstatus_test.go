package main

import (
	//"fmt"
	"github.com/coopernurse/gorp"
	"github.com/jjhageman/dev-status/db"
	"github.com/jjhageman/dev-status/dev"
	"github.com/rcrowley/go-tigertonic/mocking"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
)

var dbmap *gorp.DbMap

func newDevOrFatal(t *testing.T, first_name string, last_name string, github_id string, status string) *dev.Dev {
	dev, err := dev.NewDev(first_name, last_name, github_id, status)
	if err != nil {
		t.Fatalf("new dev: %v", err)
	}
	return dev
}

func initDbMap() *gorp.DbMap {
	dbmap = db.InitDb("postgres://jjhageman@localhost:5432/devstatus_test?sslmode=disable")
	dbmap.TraceOn("", log.New(os.Stdout, "gorptest: ", log.Lmicroseconds))
	dbmap.AddTableWithName(dev.Dev{}, "devs").SetKeys(true, "ID")
	err := dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	return dbmap
}

func dropAndClose(dbmap *gorp.DbMap) {
	dbmap.DropTablesIfExists()
	dbmap.Db.Close()
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func TestGetUser(t *testing.T) {
	dbmap := initDbMap()
	defer dropAndClose(dbmap)

	dev := newDevOrFatal(t, "Bob", "Jones", "killer_bob", "unavailable")
	dev.Save()

	code, _, response, err := getUser(
		mocking.URL(mux, "GET", "/user/"+strconv.FormatInt(dev.ID, 10)),
		mocking.Header(nil),
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
	if code != http.StatusOK {
		t.Fatal(code)
	}
	if response.FirstName != dev.FirstName {
		t.Errorf("expected Bob, got %v", response.FirstName)
		t.Fatal(response)
	}
	if response.LastName != dev.LastName {
		t.Errorf("expected %v, got %v", dev.LastName, response.LastName)
		t.Fatal(response)
	}
	if response.GithubID != dev.GithubID {
		t.Errorf("expected %v, got %v", dev.GithubID, response.GithubID)
		t.Fatal(response)
	}
	if response.Status != dev.Status {
		t.Errorf("expected %v, got %v", dev.Status, response.Status)
		t.Fatal(response)
	}
}

func TestGetInvalidUser(t *testing.T) {
	dbmap := initDbMap()
	defer dropAndClose(dbmap)

	code, _, response, err := getUser(
		mocking.URL(mux, "GET", "/user/99"),
		mocking.Header(nil),
		nil,
	)
	log.Println(err)
	if err.Error() != "User not found" {
		t.Errorf("expected User not found error, got %v", err.Error())
		t.Fatal(err)
	}

	if code != http.StatusNotFound {
		t.Errorf("expected StatusNotFound, got %v", code)
		t.Fatal(code)
	}

	if response != nil {
		t.Errorf("expected 'Invalid user' status, got %v", response.Status)
		t.Fatal(response.Status)
	}
}
