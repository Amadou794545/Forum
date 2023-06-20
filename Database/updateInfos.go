package Database

import (
	"database/sql"
	"log"
)

func UpdateImgProfile(imgPath string, userID int) {
	var db *sql.DB
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("UPDATE Users SET imgPath = $1 WHERE id_user = $2", imgPath, userID)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}
func UpdateUsername(username string, id_user int) {
	var db *sql.DB
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(
		"UPDATE users SET pseudo = $1 WHERE id_user = $2", username, id_user)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}
