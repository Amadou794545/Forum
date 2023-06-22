package Database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	CreateUsersTable()
	CreateCommentsTable()
	CreateCommentsDislikesTable()
	CreateCommentsLikesTable()
	CreatePostDislikesTable()
	CreatePostLikesTable()
	CreateLikesTable()
	CreatePostsTable()
	CreateHobbiesTable()

	AddHobbie("../Pictures/Dadas/cinema.png", "Cinema")
	AddHobbie("../Pictures/Dadas/cuisine.png", "Cuisine")
	AddHobbie("../Pictures/Dadas/informatique.png", "Informatique")
	AddHobbie("../Pictures/Dadas/jeux.png", "Jeux")
	AddHobbie("../Pictures/Dadas/lecture.png", "Lecture")
	AddHobbie("../Pictures/Dadas/musique.png", "Musique")
	AddHobbie("../Pictures/Dadas/sport.png", "Sport")
}
