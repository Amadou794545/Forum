package main

import "log"

func CheckUsername(pseudo string) bool {
	rows, err := db.Query("SELECT pseudo FROM Users WHERE pseudo = $1", pseudo)
	if err != nil {
		return false
	}
	defer rows.Close()
	if rows.Next() {
		return true // exist
	}

	return false // does not exist
}

func CheckEmail(email string) bool {
	rows, err := db.Query("SELECT email FROM Users WHERE email = $1", email)
	if err != nil {
		return false
	}
	defer rows.Close()
	if rows.Next() {
		return true // exist
	}

	return false // does not exist
}

func CheckLogin(identifier string, password string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Users WHERE (email = ? OR pseudo = ?) AND password = ?", identifier, identifier, password).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}
