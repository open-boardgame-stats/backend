// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Player struct {
	ID    string  `json:"id"`
	Owner *User   `json:"owner"`
	Name  *string `json:"name"`
}

type PlayerInput struct {
	Name *string `json:"name"`
}

type User struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"createdAt"`
	Player    *Player `json:"player"`
}

type UserInput struct {
	Name string `json:"name"`
}