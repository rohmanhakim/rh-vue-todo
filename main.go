package main

import (
	"github.com/rohmanhakim/rh-vue-todo/database"
	_ "github.com/lib/pq"
	"github.com/unrolled/render"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/rohmanhakim/rh-vue-todo/controller"
	"log"
)

func main(){

	if database.ConnectToDb() == false {
		return
	}

	if database.InitTables() == false {
		return
	}

	ren := render.New()
	router := mux.NewRouter().StrictSlash(false)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.HandleFunc("/task/all",func(w http.ResponseWriter, r *http.Request) {controller.GetAllTaskHandler(w,r,ren)}).Methods("GET")
	router.HandleFunc("/task",func(w http.ResponseWriter, r *http.Request) {controller.PostAddNewTaskHandler(w,r,ren)}).Methods("POST")
	router.HandleFunc("/task/{id}",func(w http.ResponseWriter, r *http.Request) {controller.DeleteTaskHandler(w,r,ren)}).Methods("DELETE")
	router.HandleFunc("/task/{id}",func(w http.ResponseWriter, r *http.Request) {controller.PostGetTaskDetails(w,r,ren)}).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()

}
