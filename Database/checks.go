package main

func checkUsername(pseudo string) bool {
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

func checkEmail(email string) bool {
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
