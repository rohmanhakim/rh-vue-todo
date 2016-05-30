package controller

import (
	"fmt"
	"github.com/rohmanhakim/rh-vue-todo/database"
	"strconv"
	"github.com/rohmanhakim/rh-vue-todo/model"
	"net/http"
	"github.com/unrolled/render"
	"encoding/json"
"github.com/rohmanhakim/rh-vue-todo/helper"
)


func PostRegisterNewUser(rw http.ResponseWriter, req *http.Request, ren *render.Render){
	var id 						int
	var err 					error
	var user 					model.User
	var response 	model.CommonWithIdResponse

	err = json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		panic(err)
		helper.RenderErrorResponse(500,rw,ren)
	}

	id, err = InsertUserToDb(user)
	if err != nil {
		panic(err)
		helper.RenderErrorResponse(500,rw,ren)
	}

	response.Status = 200
	response.Success = true
	response.Id = strconv.Itoa(id)
	ren.JSON(rw,http.StatusOK,response)
}

func SelectUserFromDb(id int) (model.Task, error){

	var user model.User

	query := "SELECT * FROM " + database.TABLE_USERS + " WHERE id = " + strconv.Itoa(id)

	if database.IsConnectedToDb() {

		row, err := database.GetDb().Query(query)
		if err != nil {
			panic(err)
			return user, err
		}

		for (row.Next()) {

			var _id 		int
			var _email 		string
			var _password 	string

			err := row.Scan(&_id, &_email, &_password)
			if err != nil {
				panic(err)
				return task, err
			}

			user.Id 		= strconv.Itoa(_id)
			user.Email  	= _email
			user.Password 	= _password
		}

	}

	return user, nil
}

func InsertUserToDb(user model.User) (int,error) {

	var id 	int

	if database.IsConnectedToDb() {

		query := "INSERT INTO " + database.TABLE_USERS + "(email,password) VALUES ("
		query += "'" + user.Email + "',"
		query += "'" + user.Password + "') RETURNING id"

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

			fmt.Printf("Success inserting new user with id %d\n",id)
		}
	}
	return id, nil
}
