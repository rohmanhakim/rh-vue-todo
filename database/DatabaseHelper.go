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
	if CreateTaskTable() == false {
		return false
	}

	return true
}

func CreateTaskTable() bool{
	// create Task table
	_, err = db.Exec("create table if not exists task(id varchar PRIMARY KEY, title varchar, notes varchar)")
	if err != nil {
		panic(err)
		return false
	}

	return true
}
