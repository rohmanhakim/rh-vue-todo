package controller

import (
	"net/http"
	"github.com/unrolled/render"
	"github.com/rohmanhakim/rh-vue-todo/model"
	"github.com/rohmanhakim/rh-vue-todo/database"
	"encoding/json"
	"strconv"
	"fmt"
	"github.com/rohmanhakim/rh-vue-todo/helper"
	"github.com/gorilla/mux"
)

func GetAllTaskHandler(rw http.ResponseWriter, req *http.Request, ren *render.Render){

	var getAllTasksResponse model.GetAllTasksResponse
	var tasks 				[]model.Task
	var err 				error

	tasks,err = SelectAllTaskFromDb()
	if err != nil {
		panic(err)
		helper.RenderErrorResponse(500,rw,ren)
	}

	getAllTasksResponse.Tasks = tasks

	ren.JSON(rw,http.StatusOK,getAllTasksResponse)
}

func PostAddNewTaskHandler(rw http.ResponseWriter, req *http.Request, ren *render.Render){

	var id 						int
	var err 					error
	var task 					model.Task
	var postAddNewTaskResponse 	model.PostAddNewTaskResponse

	err = json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		panic(err)
		helper.RenderErrorResponse(500,rw,ren)
	}

	id, err = InsertTaskToDb(task)
	if err != nil {
		panic(err)
		helper.RenderErrorResponse(500,rw,ren)
	}

	postAddNewTaskResponse.Status = 200
	postAddNewTaskResponse.Success = true
	postAddNewTaskResponse.Id = strconv.Itoa(id)
	ren.JSON(rw,http.StatusOK,postAddNewTaskResponse)
}

func PostGetTaskDetails(rw http.ResponseWriter, req *http.Request, ren *render.Render){

	var err 	error
	var task 	model.Task
	var id 		int

	vars := mux.Vars(req)
	id, err = strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
		helper.RenderErrorResponse(500,rw,ren)
	}

	task,err = SelectTaskFromDb(id)
	if err != nil {
		panic(err)
		helper.RenderErrorResponse(500,rw,ren)
	}

	if task.Id == "" {
		http.NotFound(rw,req)
	} else {
		ren.JSON(rw,http.StatusOK,task)
	}
}

func DeleteTaskHandler(rw http.ResponseWriter, req *http.Request, ren *render.Render){

	var err 			error
	var id 				string

	vars := mux.Vars(req)
	id = vars["id"]

	err = DeleteTaskFromDb(id)
	if err != nil {
		panic(err)
		helper.RenderErrorResponse(500,rw,ren)
	}

	helper.RenderOKResponse(rw,ren)
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

func SelectTaskFromDb(id int) (model.Task, error){
	
	var task model.Task

	query := "SELECT * FROM task WHERE id = " + strconv.Itoa(id)

	if database.IsConnectedToDb() {

		row, err := database.GetDb().Query(query)
		if err != nil {
			panic(err)
			return task, err
		}

		for (row.Next()) {

			var _id 	int
			var _title 	string
			var _notes 	string

			err := row.Scan(&_id, &_title, &_notes)
			if err != nil {
				panic(err)
				return task, err
			}

			task.Id 	= strconv.Itoa(_id)
			task.Title = _title
			task.Notes = _notes

		}

	}

	return task, nil
}

func InsertTaskToDb(task model.Task) (int,error) {

	var id 	int

	if database.IsConnectedToDb() {

		query := "INSERT INTO task(title,notes) VALUES ("
		query += "'" + task.Title + "',"
		query += "'" + task.Notes + "') RETURNING id"

		row, err := database.GetDb().Query(query)
		if err != nil {
			panic(err)
			return 0, err
		}

		for (row.Next()) {

			err := row.Scan(&id)
			if err != nil {
				return 0, err
			}

			fmt.Printf("Success inserting new task with id %d\n",id)
		}
	}
	return id, nil
}

func DeleteTaskFromDb(id string) error {

	query := "DELETE FROM task where id = "
	query += id

	if database.IsConnectedToDb() {

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

		fmt.Printf("Success inserting new task with id %d\n", id)
	}

	return nil
}
