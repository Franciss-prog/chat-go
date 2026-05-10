package main

type Auth struct {
	Email    string `json:"email"`
	Password int    `json:"password"`
}

type Register struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password int    `json:"password"`
}
