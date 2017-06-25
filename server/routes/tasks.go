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

const tableName = "tasks"

var errorJSON = struct {
	err bool
}{
	true,
}
var successJSON = struct {
	success bool
}{
	true,
}

func get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projection := "id, title, description, status"

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
			err := taskRowsQuery.Scan(&(taskModel.ID), &(taskModel.Title), &(taskModel.Description), &(taskModel.Status))
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

func post(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var taskModel models.Task

	err := decoder.Decode(&taskModel)
	taskModel.Status = "pending"

	if err != nil {
		result, _ := json.Marshal(errorJSON)
		w.Write(result)
	} else {
		defer r.Body.Close()

		_, err := squirrel.
			Insert(tableName).
			Columns("title", "description", "status").
			Values(taskModel.Title, taskModel.Description, taskModel.Status).
			RunWith(db.Conn).
			Query()

		if err != nil {
			result, _ := json.Marshal(errorJSON)
			w.Write(result)
		} else {

			idQuery := squirrel.
				Select("id").
				From(tableName).
				OrderBy("id DESC").
				Limit(1).
				RunWith(db.Conn)

			idQuery.QueryRow().Scan(&(taskModel.ID))

			result, _ := json.Marshal(taskModel)
			w.Write(result)
		}
	}

}

func put(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var taskModel models.Task

	err := decoder.Decode(&taskModel)

	if err != nil {
		result, _ := json.Marshal(errorJSON)
		w.Write(result)
	} else {
		defer r.Body.Close()

		_, err := squirrel.
			Update(tableName).
			SetMap(map[string]interface{}{
				"title":       taskModel.Title,
				"description": taskModel.Description,
				"status":      taskModel.Status,
			}).
			Where(squirrel.Eq{"id": taskModel.ID}).
			RunWith(db.Conn).
			Query()

		if err != nil {
			result, _ := json.Marshal(errorJSON)
			w.Write(result)
		} else {
			result, _ := json.Marshal(successJSON)
			w.Write(result)
		}
	}

}

// TaskRoute endpoint for task related CRUD operations
func TaskRoute(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		get(w, r)
	case http.MethodPost:
		post(w, r)
	case http.MethodPut:
		put(w, r)
	default:
		w.Write([]byte("Unsupported method"))
	}

}
