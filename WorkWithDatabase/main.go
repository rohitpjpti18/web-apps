package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	id         int
	username   string
	surname    string
	age        int
	university string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}

	// catch to error.

}

func addUser(db *sql.DB, username string, surname string, age int, university string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into testTable (username, surname, age, university) values (?, ?, ?, ?)")
	_, err := stmt.Exec(username, surname, age, university)
	checkError(err)
	tx.Commit()
}

func getUsers(db *sql.DB, id2 int) User {
	rows, err := db.Query("select * from testTable")
	checkError(err)
	for rows.Next() {
		var tempUser User
		err = rows.Scan(&tempUser.id, &tempUser.username, &tempUser.surname, &tempUser.age, &tempUser.university)
		checkError(err)
		if tempUser.id == id2 {
			return tempUser
		}
	}
	return User{}
}

func main() {
	db, _ := sql.Open("sqlite3", "database/godb.db")
	db.Exec("create table if not exists testTable (id integer, username text, surname text, age Integer, university text)")
	addUser(db, "harry", "potter", 19, "Hogwarts")
	addUser(db, "ronald", "weasley", 19, "Hogwarts")
}
