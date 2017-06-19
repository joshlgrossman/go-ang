package main

import (
	"database/sql"
	"net/http"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joshlgrossman/go-ang/server/db"
	"github.com/joshlgrossman/go-ang/server/routes"
)

func main() {
	var err error
	db.Conn, err = sql.Open("mysql", "root:@/test")
	defer db.Conn.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/ws/tasks/{id}", routes.TaskRoute)
	router.HandleFunc("/ws/tasks", routes.TaskRoute)
	router.HandleFunc("/ws/test", routes.Test)

	http.Handle("/ws/", router)
	http.HandleFunc("/", routes.Static("client/build/"))
	http.ListenAndServe(":8080", nil)

}
