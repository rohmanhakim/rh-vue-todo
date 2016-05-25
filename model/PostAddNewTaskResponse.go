package model

type PostAddNewTaskResponse struct {
	Status 	int 	`json:"status"`
	Success bool 	`json:"success"`
	Id 		string  `json:"id"`
}
