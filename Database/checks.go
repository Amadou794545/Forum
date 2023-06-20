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
