package Database

import (
	"database/sql"
	"log"
)

func DeleteUser(userID int) error {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
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
	db, err = sql.Open("sqlite3", "./database.db")
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
	db, err = sql.Open("sqlite3", "./database.db")
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

func RemoveDislikePost(userID int, postID int) error {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM PostsDislikes WHERE id_user = $1 AND id_post = $2", userID, postID)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}

func RemoveDislikeComment(userID int, postID int) error {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM CommentsDislikes WHERE id_user = $1 AND id_comment = $2", userID, postID)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}

func RemoveLikeComment(userID int, postID int) error {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM CommentsLikes WHERE id_user =$1 AND id_comment = $2", userID, postID)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}

func RemoveLikePost(userID int, postID int) error {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM PostsLikes WHERE id_user =$1 AND id_post = $2", userID, postID)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}
