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
            id_user INTEGER PRIMARY KEY,
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
            id_post INTEGER PRIMARY KEY,
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
            id_comment INTEGER PRIMARY KEY,
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
            id_like INTEGER PRIMARY KEY,
            id_post INTEGER,
            id_comment INTEGER,
            is_like INTEGER,
            FOREIGN KEY (id_post) REFERENCES Posts(id_post) ON DELETE CASCADE,
            FOREIGN KEY (id_comment) REFERENCES Comments(id_comment) ON DELETE CASCADE
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
