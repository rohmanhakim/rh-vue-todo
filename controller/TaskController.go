package controller

import (
	"net/http"
	"github.com/unrolled/render"
	"github.com/rohmanhakim/rh-vue-todo/model"
	"github.com/rohmanhakim/rh-vue-todo/database"
)

func GetAllTaskHandler(rw http.ResponseWriter, req *http.Request, ren *render.Render){

	var getAllTasksResponse model.GetAllTasksResponse
	var tasks 				[]model.Task
	var err 				error

	tasks,err = GetAllTaskFromDb()
	if err != nil {
		panic(err)
	}

	getAllTasksResponse.Tasks = tasks

	ren.JSON(rw,http.StatusOK,getAllTasksResponse)
}

func GetAllTaskFromDb() ([]model.Task, error) {

	var tasks []model.Task

	if database.IsConnectedToDb() {

		row, err := database.GetDb().Query(`SELECT * FROM task`)
		if err != nil {
			panic(err)
		}

		for (row.Next()) {

			var _task 	model.Task
			var _id 	string
			var _title 	string
			var _notes 	string

			err := row.Scan(&_id, &_title, &_notes)
			if err != nil {
				return tasks, err
			}

			_task.Id 	= _id
			_task.Title = _title
			_task.Notes = _notes

			tasks = append(tasks,_task)
		}
	}
	return tasks, nil
}
