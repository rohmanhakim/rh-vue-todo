package model

type Task struct {
	Id		string 	`json:"id"`
	Title 	string 	`json:"title"`
	Notes	string	`json:"notes"`
	Done	bool	`json:"done"`
}
