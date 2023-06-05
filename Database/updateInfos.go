package main

import "log"

func UpdateUserIMG(path string, id_user int) {
	_, err := db.Exec(
		"UPDATE users SET imgPath = $1 WHERE id_user = $2", path, id_user)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateUsername(username string, id_user int) {
	_, err := db.Exec(
		"UPDATE users SET pseudo = $1 WHERE id_user = $2", username, id_user)
	if err != nil {
		log.Fatal(err)
	}
}
