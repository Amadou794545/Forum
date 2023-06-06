package Database

import (
	"database/sql"
	"log"
)

func DeleteUser(userID int) error {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM Users WHERE id_user = ?", userID)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}

func DeletePost(postID int) error {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM Posts WHERE id_post = ?", postID)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}

func DeleteComment(commentID int) error {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM Comments WHERE id_comment = ?", commentID)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}
