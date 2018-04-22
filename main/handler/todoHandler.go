package handler

import (
	"net/http"
	"TodoApp/main/model"
	"strconv"
	"encoding/json"
)


func AllTodosHandler(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(r.Header.Get("user_id"))

	respondWithJSON(w, http.StatusOK, model.FindAllTodoByUserId(userId))
}

func AddAllTodoHandler(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(r.Header.Get("user_id"))
	var todos []model.Todo
	json.NewDecoder(r.Body).Decode(&todos)
	model.ReinsertAll(todos, userId)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
