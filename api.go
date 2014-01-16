package main

type UserRequest struct {
	ID string `json:"id"`
}

type UserResponse struct {
	FirstName string
	LastName  string
	GithubID  string
	Status    string
}
