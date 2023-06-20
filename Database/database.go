package Database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	CreateUsersTable()
	CreateCommentsTable()
	CreateLikesTable()
	CreatePostsTable()
	CreateHobbiesTable()

	AddHobbie("../pictures/Dadas/cinema.png", "Cinema")
	AddHobbie("../pictures/Dadas/cuisine.png", "Cuisine")
	AddHobbie("../pictures/Dadas/informatique.png", "Informatique")
	AddHobbie("../pictures/Dadas/jeux.png", "Jeux")
	AddHobbie("../pictures/Dadas/lecture.png", "Lecture")
	AddHobbie("../pictures/Dadas/musique.png", "Musique")
	AddHobbie("../pictures/Dadas/sport.png", "Sport")
}
