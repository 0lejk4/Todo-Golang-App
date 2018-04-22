package model

import (
	"TodoApp/main/database"
)

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserId    int    `db:"user_id"`
}

var todoSchema = `
	create table todo  (
	id serial primary key ,
	completed boolean,
    title text,
	user_id INTEGER     NOT NULL
    CONSTRAINT todo_user_fk
    REFERENCES user_entity
);`

func CreateTodoTable() {
	database.SQL.MustExec(todoSchema)
	populateTodo()
}

func populateTodo() {
	SaveTodo(Todo{UserId: 1, Title: "Kek", Completed: false})
	SaveTodo(Todo{UserId: 1, Title: "Kek", Completed: false})
	SaveTodo(Todo{UserId: 1, Title: "Kek", Completed: false})
}

func SaveTodo(todo Todo) {
	database.SQL.MustExec("INSERT INTO todo (title, completed, user_id) VALUES ($1, $2, $3)", todo.Title, todo.Completed, todo.UserId)
}

func DeleteAllTodoOfUser(user_id int) {
	database.SQL.MustExec("DELETE from todo where user_id = $1", user_id)
}

func ReinsertAll(todos []Todo, user_id int) {
	DeleteAllTodoOfUser(user_id)
	for _, todo := range todos {
		SaveTodo(todo)
	}
}

func FindAllTodoByUserId(userId int) ([]Todo) {
	todos := []Todo{}
	database.SQL.Select(&todos, "SELECT * from todo where user_id = $1", userId)
	return todos
}
