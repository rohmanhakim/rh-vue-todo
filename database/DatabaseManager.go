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

	return true
}

func CreateTasksTable() bool{
	// create Tasks table
	_, err = db.Exec("create table if not exists " + TABLE_TASKS + "(id SERIAL PRIMARY KEY, title varchar, notes varchar, done  boolean)")
	if err != nil {
		panic(err)
		return false
	}

	return true
}
