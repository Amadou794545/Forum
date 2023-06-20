package Database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddUser(email string, pseudo string, password string, imgPath string) {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	if imgPath == "" {
		imgPath = "default.jpg"
	}

	HashedPassword, error := GenerateFromPassword(password)

	if error != nil {
		fmt.Println(error)
	}

	_, err = db.Exec(`
		INSERT INTO users (email, pseudo, password, imgPath) VALUES ($1, $2, $3, $4);
	`, email, pseudo, HashedPassword, imgPath)

	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

func AddPost(title string, imagePath string, description string, userID int, hobbieID int) {
	var db *sql.DB
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		INSERT INTO Posts (title, imgPath, description, id_user, id_hobbie) VALUES ($1, $2, $3, $4, $5);
	`, title, imagePath, description, userID, hobbieID)

	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

func AddComment(description string, userID int, postID string) {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		INSERT INTO Comments (description, id_user, id_post) VALUES ($1, $2, $3);
	`, description, userID, postID)

	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

func AddLike(userID int, postID int, commentID int, isLike int) {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	likeValue := 0
	if isLike == 1 {
		likeValue = 1
	} else if isLike == -1 {
		likeValue = -1
	} else {
		log.Fatal("likeValue invalide")
	}

	_, err = db.Exec(`
		INSERT INTO Likes (id_post, id_comment, is_like) VALUES ($1, $2, $3);
	`, postID, commentID, likeValue)

	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

func AddHobbie(imgPath string, description string) {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		INSERT INTO Hobbies (img_path, description) VALUES ($1, $2);
	`, imgPath, description)

	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}
