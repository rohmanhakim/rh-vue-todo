package main

import (
	"github.com/rohmanhakim/rh-vue-todo/database"
	_ "github.com/lib/pq"
)

func main(){

	database.ConnectToDb()
	database.InitTables()
}
