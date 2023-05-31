package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "ma_base_de_donnees.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTableSQLUser := `
		CREATE TABLE IF NOT EXISTS Users (
			id_user INTEGER PRIMARY KEY,
			email TEXT,
			pseudo TEXT,
			mdp TEXT,
		);
	`

	_, err = db.Exec(createTableSQLUser)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQLPosts := `
		CREATE TABLE IF NOT EXISTS Posts (
			id_post INTEGER PRIMARY KEY,
			img IMAGE,
			title TEXT,
			description TEXT,
			id_user INTEGER,
			FOREIGN KEY (id_user) REFERENCES User(id_user)
		);
	`

	_, err = db.Exec(createTableSQLPosts)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQLComments := `
		CREATE TABLE IF NOT EXISTS Comments (
			id_comment INTEGER PRIMARY KEY,
			description TEXT,
			id_user INTEGER,
			FOREIGN KEY (id_user) REFERENCES Users(id_user),
			id_post INTEGER,
			FOREIGN KEY (id_post) REFERENCES Posts(id_post)
		);
	`

	_, err = db.Exec(createTableSQLComments)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQLLikes := `
		CREATE TABLE IF NOT EXISTS Likes (
			id_like INTEGER PRIMARY KEY,
			id_post INTEGER,
			FOREIGN KEY (id_post) REFERENCES Posts(id_post) ON DELETE SET NULL,
			id_comment INTEGER,
			FOREIGN KEY (id_comment) REFERENCES Comments(id_comment) ON DELETE SET NULL,
			is_like INTEGER
		);
	`

	_, err = db.Exec(createTableSQLLikes)
	if err != nil {
		log.Fatal(err)
	}

	/*insertUserSQL := `
		INSERT INTO utilisateurs (nom) VALUES ('John Doe');
	`

	_, err = db.Exec(insertUserSQL)
	if err != nil {
		log.Fatal(err)
	}*/
}
