package main

import (
	_ "github.com/lib/pq"
	"net/http"
	"encoding/base64"
	"strings"
	"github.com/gorilla/mux"
	"TodoApp/main/handler"
	"TodoApp/main/database"
	"TodoApp/main/model"
	"log"
	"strconv"
)

func main() {
	database.Connect()
	defer database.SQL.Close()

	model.CreateUserTable()
	model.CreateTodoTable()
	r := mux.NewRouter()
	r.HandleFunc("/app/todos", use(handler.AllTodosHandler, basicAuth)).Methods("GET")
	r.HandleFunc("/app/todos", use(handler.AddAllTodoHandler, basicAuth)).Methods("PUT")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))
	http.Handle("/", r)
	log.Println("App started")
	http.ListenAndServe(":8080", nil)

}

func basicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {


		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		log.Println("username = ",pair[0]," password = ", pair[1])
		user, findErr := model.FindUserByUsername(pair[0])


		if findErr != nil || pair[0] != user.Username || pair[1] != user.Password {
			http.Error(w, "Not authorized", 401)
			return
		}
		r.Header.Set("user_id", strconv.Itoa(user.Id))
		h.ServeHTTP(w, r)
	}
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}
