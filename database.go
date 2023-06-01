package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createUsersTable(db)
	createCommentsTable(db)
	createLikesTable(db)
	createPostsTable(db)

	newUser(db)
}

func createUsersTable(db *sql.DB) {
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

func createPostsTable(db *sql.DB) {
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

func createCommentsTable(db *sql.DB) {
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

func createLikesTable(db *sql.DB) {
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

func newUser(db *sql.DB) {
	_, err := db.Exec(`
		INSERT INTO users (email, pseudo, password) VALUES ('test@gmail.com', 'itsMe', 'passw0rd');
	`)

	if err != nil {
		log.Fatal(err)
	}
}
