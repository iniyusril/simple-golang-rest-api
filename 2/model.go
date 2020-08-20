package main

type Users struct {
	Id        string `form:"id" json:"id"`
	FirstName string `form:"firstname" json:"first_name"`
	LastName  string `form:"lastname" json:"last_name"`
}

type Response struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Users `json:"data"`
}
