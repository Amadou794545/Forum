package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createUsersTable()
	createCommentsTable()
	createLikesTable()
	createPostsTable()
}

func createUsersTable() {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS Users (
            id_user INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT,
            pseudo TEXT,
            password TEXT
        )
    `)
	if err != nil {
		log.Fatal(err)
	}
}

func createPostsTable() {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS Posts (
            id_post INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT,
            description TEXT,
			id_user INTEGER,
			FOREIGN KEY (id_user) REFERENCES Users(id_user)
        )
    `)
	if err != nil {
		log.Fatal(err)
	}
}

func createCommentsTable() {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS Comments (
            id_comment INTEGER PRIMARY KEY AUTOINCREMENT,
            description TEXT,
			id_user INTEGER,
			id_post INTEGER,
			FOREIGN KEY (id_user) REFERENCES Users(id_user),
			FOREIGN KEY (id_post) REFERENCES Posts(id_post)
        )
    `)
	if err != nil {
		log.Fatal(err)
	}
}

func createLikesTable() {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS Likes (
            id_like INTEGER PRIMARY KEY AUTOINCREMENT,
			id_user INTEGER,
            id_post INTEGER,
            id_comment INTEGER,
			FOREIGN KEY (id_user) REFERENCES Users(id_user),
            FOREIGN KEY (id_post) REFERENCES Posts(id_post) ON DELETE SET NULL,
            FOREIGN KEY (id_comment) REFERENCES Comments(id_comment) ON DELETE SET NULL,
            is_like INTEGER
        )
    `)
	if err != nil {
		log.Fatal(err)
	}
}

func addUser(email string, pseudo string, password string) {
	_, err := db.Exec(`
		INSERT INTO users (email, pseudo, password) VALUES ($1, $2, $3);
	`, email, pseudo, password)

	if err != nil {
		log.Fatal(err)
	}
}

func addPost(title string, description string, userID int) {
	_, err := db.Exec(`
		INSERT INTO Posts (title, description, id_user) VALUES ($1, $2, $3);
	`, title, description, userID)

	if err != nil {
		log.Fatal(err)
	}
}

func addComment(description string, userID int, postID int) {
	_, err := db.Exec(`
		INSERT INTO Comments (description, id_user, id_post) VALUES ($1, $2, $3);
	`, description, userID, postID)

	if err != nil {
		log.Fatal(err)
	}
}

func addLike(userID int, postID int, commentID int, isLike int) {
	likeValue := 0
	if isLike == 1 {
		likeValue = 1
	} else if isLike == -1 {
		likeValue = -1
	} else {
		log.Fatal("likeValue invalide")
	}

	_, err := db.Exec(`
		INSERT INTO Likes (id_post, id_comment, is_like) VALUES ($1, $2, $3);
	`, postID, commentID, likeValue)

	if err != nil {
		log.Fatal(err)
	}
}
