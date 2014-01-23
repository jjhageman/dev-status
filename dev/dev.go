package dev

import (
	"fmt"
	//"github.com/coopernurse/gorp"
)

type Dev struct {
	ID        int64
	FirstName string
	LastName  string
	GithubID  string
	Status    string
}

var statuses = [3]string{"available", "looking", "unavailable"}

func NewDev(first_name string, last_name string, github_id string, status string) (*Dev, error) {
	if !validStatus(status) {
		return nil, fmt.Errorf("invalid status")
	}
	return &Dev{0, first_name, last_name, github_id, status}, nil
}

func (d *Dev) save() error {
	return nil
}

func All() []*Dev {
	var devs []*Dev
	_, err := Dbmap.Select(&devs, "select * from devs order by id")
	return devs
}

func validStatus(status string) bool {
	for _, s := range statuses {
		if s == status {
			return true
		}
	}
	return false
}
