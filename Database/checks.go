package Database

import (
	"database/sql"
	"log"
)

func CheckUsername(pseudo string) bool {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT pseudo FROM Users WHERE pseudo = $1", pseudo)
	if err != nil {
		return false
	}
	defer rows.Close()
	if rows.Next() {
		return true // exist
	}
	db.Close()
	return false // does not exist
}

func CheckCommentDislikes(userID int, postID int) bool {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return false
	}
	defer db.Close()
	var counter int
	err = db.QueryRow("SELECT COUNT(*) FROM CommentsDislikes WHERE id_user = $1 AND id_comment = $2", userID, postID).Scan(&counter)
	if counter == 0 {
		return true
	}
	return false
}

func CheckCommentLikes(userID int, postID int) bool {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return false
	}
	defer db.Close()
	var counter int
	err = db.QueryRow("SELECT COUNT(*) FROM CommentsLikes WHERE id_user = $1 AND id_comment = $2", userID, postID).Scan(&counter)
	if counter == 0 {
		return true
	}
	return false
}

func CheckDislikes(userID int, postID int) bool {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return false
	}
	defer db.Close()
	var counter int
	err = db.QueryRow("SELECT COUNT(*) FROM PostsDislikes WHERE id_user = $1 AND id_Post = $2", userID, postID).Scan(&counter)
	if counter == 0 {
		return true
	}
	return false
}

func CheckLikes(userID int, postID int) bool {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return false
	}
	defer db.Close()
	var counter int
	err = db.QueryRow("SELECT COUNT(*) FROM PostsLikes WHERE id_user = $1 AND id_Post = $2", userID, postID).Scan(&counter)
	if counter == 0 {
		return true
	}
	return false
}

func CheckEmail(email string) bool {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT email FROM Users WHERE email = $1", email)
	if err != nil {
		return false
	}
	defer rows.Close()
	if rows.Next() {
		return true // exist
	}
	db.Close()
	return false // does not exist
}

func CheckLogin(identifier string, password string) bool {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	var encodedHash string
	err = db.QueryRow("SELECT password FROM Users WHERE email = ? OR pseudo = ?", identifier, identifier).Scan(&encodedHash)
	if err != nil {
		log.Fatal(err)
	}

	match, err := ComparePasswordAndHash(password, encodedHash)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
	return match
}
