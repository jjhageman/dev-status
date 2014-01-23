package dev

import "testing"

func newDevOrFatal(t *testing.T, first_name string, last_name string, github_id string, status string) *Dev {
	dev, err := NewDev(first_name, last_name, github_id, status)
	if err != nil {
		t.Fatalf("new dev: %v", err)
	}
	return dev
}

func TestAll(t *testing.T) {
	dev1 := newDevOrFatal(t, "Bob", "Jones", "killer_bob", "unavailable")
	dev2 := newDevOrFatal(t, "Bob", "Jones", "killer_bob", "unavailable")

	dev1.save()
	dev2.save()

	devs := All()
	if len(devs) != 2 {
		t.Errorf("expected 2 devs, got %v", len(devs))
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
	dev := newDevOrFatal(t, "Bob", "Jones", "killer_bob", "unavailable")
	dev.save()
}
