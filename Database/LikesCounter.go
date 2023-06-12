package Database

import (
	"database/sql"
	"log"
)

func CountLikesPost(postID int) (int, error) {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Likes WHERE id_post = ? AND is_like = 1", postID).Scan(&count)
	if err != nil {
		return 0, err
	}

	db.Close()

	return count, nil
}

func CountDislikesPost(postID int) (int, error) {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Likes WHERE id_post = ? AND is_like = -1", postID).Scan(&count)
	if err != nil {
		return 0, err
	}

	db.Close()

	return count, nil
}

func CountLikesComment(commentID int) (int, error) {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Likes WHERE id_comment = ? AND is_like = 1", commentID).Scan(&count)
	if err != nil {
		return 0, err
	}

	db.Close()

	return count, nil
}

func CountDislikesComment(commentID int) (int, error) {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Likes WHERE id_comment = ? AND is_like = -1", commentID).Scan(&count)
	if err != nil {
		return 0, err
	}

	db.Close()

	return count, nil
}
