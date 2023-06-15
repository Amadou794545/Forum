package Database

import (
	"database/sql"
	"fmt"
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
	fmt.Println(imgPath)
	_, err = db.Exec("UPDATE Users SET imgPath = $1 WHERE id_user = $2", imgPath, userID)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}
