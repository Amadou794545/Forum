package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	CreateUsersTable()
	CreateCommentsTable()
	CreateLikesTable()
	CreatePostsTable()
	CreateHobbiesTable()

	AddHobbie("../pictures/cinema.png", "Cinema")
	AddHobbie("../pictures/cuisine.png", "Cuisine")
	AddHobbie("../pictures/informatique.png", "Informatique")
	AddHobbie("../pictures/jeux.png", "Jeux")
	AddHobbie("../pictures/lecture.png", "Lecture")
	AddHobbie("../pictures/musique.png", "Musique")
	AddHobbie("../pictures/sport.png", "Sport")
}

func checkLogin(identifier string, password string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Users WHERE (email = ? OR pseudo = ?) AND password = ?", identifier, identifier, password).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}
