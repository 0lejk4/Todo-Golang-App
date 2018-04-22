package model

import (
	"TodoApp/main/database"
)

type User struct {
	Id       int
	Username string
	Password string
}

var userSchema = `
	DROP TABLE IF EXISTS todo;
	DROP TABLE IF EXISTS user_entity;
	create table user_entity  (
	id serial primary key,
    username text,
    password text
);`

func CreateUserTable() {
	database.SQL.MustExec(userSchema)
	populateUserTable()
}

func populateUserTable() {
	SaveUser(User{0, "0lejk4", "1234"})
	SaveUser(User{0, "Test", "1234"})
	SaveUser(User{0, "Dcp", "1234"})
}

func SaveUser(user User) {
	database.SQL.MustExec("INSERT INTO user_entity (username, password) VALUES ($1, $2)", user.Username, user.Password)
}

func FindUserById(userId int) (User) {
	user := User{}
	database.SQL.Get(&user, "SELECT * from user_entity where id = $1", userId)
	return user
}

func FindUserByUsername(username string) (User, error) {
	user := User{}
	error := database.SQL.Get(&user, "SELECT * from user_entity where username = $1", username)
	return user, error
}
