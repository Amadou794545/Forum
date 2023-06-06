package Database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
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
