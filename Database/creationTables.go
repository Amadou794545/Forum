package Database

import (
	"database/sql"
	"log"
)

func CreateUsersTable() {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Users (
            id_user INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT,
            pseudo TEXT,
            password TEXT,
			imgPath TEXT
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

func CreatePostsTable() {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Posts (
            id_post INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT,
			imgPath TEXT,
            description TEXT,
			id_user INTEGER,
			nbr_like INTEGER,
			nbr_dislike INTEGER,
			id_hobbie INTEGER,
			FOREIGN KEY (id_user) REFERENCES Users(id_user),
			FOREIGN KEY (id_hobbie) REFERENCES Hobbies(id_hobbie) ON DELETE SET NULL
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

func CreateCommentsTable() {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Comments (
            id_comment INTEGER PRIMARY KEY AUTOINCREMENT,
            description TEXT,
			id_user INTEGER,
			id_post INTEGER,
			nbr_like INTEGER,
			nbr_dislike INTEGER,
			FOREIGN KEY (id_user) REFERENCES Users(id_user),
			FOREIGN KEY (id_post) REFERENCES Posts(id_post)
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

func CreateLikesTable() {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Likes (
            id_like INTEGER PRIMARY KEY AUTOINCREMENT,
			id_user INTEGER,
            id_post INTEGER,
            id_comment INTEGER,
            is_like INTEGER,
			FOREIGN KEY (id_user) REFERENCES Users(id_user),
            FOREIGN KEY (id_post) REFERENCES Posts(id_post) ON DELETE SET NULL,
            FOREIGN KEY (id_comment) REFERENCES Comments(id_comment) ON DELETE SET NULL
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

func CreatePostLikesTable() {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS PostsLikes (
			id_user INTEGER,
            id_post INTEGER,
			FOREIGN KEY (id_user) REFERENCES Users(id_user),
            FOREIGN KEY (id_post) REFERENCES Posts(id_post) ON DELETE SET NULL
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

func CreatePostDislikesTable() {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS PostsDislikes (
			id_user INTEGER,
            id_post INTEGER,
			FOREIGN KEY (id_user) REFERENCES Users(id_user),
            FOREIGN KEY (id_post) REFERENCES Posts(id_post) ON DELETE SET NULL
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

func CreateCommentsLikesTable() {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS CommentsLikes (
			id_user INTEGER,
            id_comment INTEGER,
			FOREIGN KEY (id_user) REFERENCES Users(id_user),
            FOREIGN KEY (id_comment) REFERENCES Comments(id_comment) ON DELETE SET NULL
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

func CreateCommentsDislikesTable() {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS CommentsDislikes (
			id_user INTEGER,
            id_comment INTEGER,
			FOREIGN KEY (id_user) REFERENCES Users(id_user),
            FOREIGN KEY (id_comment) REFERENCES Comments(id_comment) ON DELETE SET NULL
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

func CreateHobbiesTable() {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Hobbies (
			id_hobbie INTEGER PRIMARY KEY AUTOINCREMENT,
            img_path TEXT,
			description TEXT
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}
