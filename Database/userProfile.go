package Database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func UpdateImgProfile(imgPath string, userID int) {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("UPDATE Users SET imgPath = ? WHERE id_user = ?", imgPath, userID)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}
