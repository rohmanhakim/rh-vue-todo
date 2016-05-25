package controller

import (
	"net/http"
	"github.com/unrolled/render"
	"github.com/rohmanhakim/rh-vue-todo/model"
	"github.com/rohmanhakim/rh-vue-todo/database"
	"encoding/json"
	"strconv"
	"fmt"
)

func GetAllTaskHandler(rw http.ResponseWriter, req *http.Request, ren *render.Render){

	var getAllTasksResponse model.GetAllTasksResponse
	var tasks 				[]model.Task
	var err 				error

	tasks,err = SelectAllTaskFromDb()
	if err != nil {
		panic(err)
	}

	getAllTasksResponse.Tasks = tasks

	ren.JSON(rw,http.StatusOK,getAllTasksResponse)
}

func PostAddNewTaskHandler(rw http.ResponseWriter, req *http.Request, ren *render.Render){

	var err error
	var task model.Task
	var commonResponse model.CommonResponse

	// Decode the incoming Go-Kilat Bid
	err = json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		panic(err)
	}

	err = InsertTaskToDb(task)
	if err != nil {
		commonResponse.Status = 502
		commonResponse.Success = false
		ren.JSON(rw,http.StatusInternalServerError,commonResponse)
	}

	ren.JSON(rw,http.StatusOK,commonResponse)
}

func SelectAllTaskFromDb() ([]model.Task, error) {

	var tasks []model.Task

	if database.IsConnectedToDb() {

		row, err := database.GetDb().Query(`SELECT * FROM task`)
		if err != nil {
			panic(err)
		}

		for (row.Next()) {

			var _task 	model.Task
			var _id 	int
			var _title 	string
			var _notes 	string

			err := row.Scan(&_id, &_title, &_notes)
			if err != nil {
				return tasks, err
			}

			_task.Id 	= strconv.Itoa(_id)
			_task.Title = _title
			_task.Notes = _notes

			tasks = append(tasks,_task)
		}
	}
	return tasks, nil
}

func InsertTaskToDb(task model.Task) error {
	if database.IsConnectedToDb() {

		var err error

		query := "INSERT INTO task(title,notes) VALUES ("
		query += "'" + task.Title + "',"
		query += "'" + task.Notes + "')"

		stmt, err := database.GetDb().Prepare(query)
		if err != nil {
			panic(err)
			return err
		}

		_, err = stmt.Exec()
		if err != nil {
			panic(err)
			return err
		}

		fmt.Printf("Success inserting new task")
	}
	return nil
}