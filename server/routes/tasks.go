package routes

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/joshlgrossman/go-ang/server/db"
	"github.com/joshlgrossman/go-ang/server/models"
)

func get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projection := "title, description, status"
	tableName := "tasks"

	taskRows := squirrel.
		Select(projection).
		From(tableName).
		RunWith(db.Conn)

	if val, ok := vars["id"]; ok {
		taskRows = taskRows.Where(squirrel.Eq{"id": val})
	}

	taskRowsQuery, err := taskRows.Query()

	if err == nil {
		var taskModels []models.Task

		for taskRowsQuery.Next() {
			var taskModel models.Task
			err := taskRowsQuery.Scan(&(taskModel.Title), &(taskModel.Description), &(taskModel.Status))
			if err == nil {
				taskModels = append(taskModels, taskModel)
			}
		}

		result, _ := json.Marshal(taskModels)
		w.Write(result)
	} else {
		log.Fatal(err)
	}
}

// TaskRoute endpoint for task related CRUD operations
func TaskRoute(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		get(w, r)
	default:
		w.Write([]byte("Unsupported method"))
	}

}
