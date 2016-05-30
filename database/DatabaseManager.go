package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var dbUsername 		string = "arifrohman"
var dbName 			string = "rh-vue-todo"
var db 				*sql.DB
var err 			error
var isConnectedToDb bool

//=================== TABLE NAME =====================
var TABLE_TASKS = "tasks"
var TABLE_USERS = "users"

func ConnectToDb() bool{
	db, err = sql.Open("postgres", "user=" + dbUsername + " dbname=" + dbName + " sslmode=disable")
	if err != nil {
		panic(err)
		isConnectedToDb = false
		return false
	}

	isConnectedToDb = true
	return true
}

func IsConnectedToDb() bool{
	return isConnectedToDb
}

func DisconnectFromDb(){
	db.Close()
}

func GetDb() *sql.DB{
	return db
}

func InitTables() bool{
	if CreateTasksTable() == false {
		return false
	}

	if CreateUsersTable() == false {
		return false
	}

	return true
}

func CreateTasksTable() bool{
	// create Tasks table
	_, err = db.Exec("create table if not exists " + TABLE_TASKS + "(id SERIAL PRIMARY KEY, title varchar(255), notes varchar(255), done  boolean)")
	if err != nil {
		panic(err)
		return false
	}

	return true
}

func CreateUsersTable() bool{
	//create Users table
	_, err = db.Exec("create table if not exists " + TABLE_USERS + "(id SERIAL PRIMARY KEY, email varchar(50), password varchar(255))")
	if err != nil {
		panic(err)
		return false
	}

	return true
}
