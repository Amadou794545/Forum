package main

import (
	"fmt"
	"log"
)

func AddUser(email string, pseudo string, password string, imgPath string) {
	if imgPath == "" {
		imgPath = "default.jpg"
	}

	HashedPassword, error := GenerateFromPassword(password)

	if error != nil {
		fmt.Println(error)
	}

	_, err := db.Exec(`
		INSERT INTO users (email, pseudo, password, imgPath) VALUES ($1, $2, $3, $4);
	`, email, pseudo, HashedPassword, imgPath)

	if err != nil {
		log.Fatal(err)
	}
}

func AddPost(title string, imagePath string, description string, userID int, hobbieID int) {
	_, err := db.Exec(`
		INSERT INTO Posts (title, img_path, description, id_user, id_hobbie) VALUES ($1, $2, $3, $4);
	`, title, imagePath, description, userID, hobbieID)

	if err != nil {
		log.Fatal(err)
	}
}

func AddComment(description string, userID int, postID int) {
	_, err := db.Exec(`
		INSERT INTO Comments (description, id_user, id_post) VALUES ($1, $2, $3);
	`, description, userID, postID)

	if err != nil {
		log.Fatal(err)
	}
}

func AddLike(userID int, postID int, commentID int, isLike int) {
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

func AddHobbie(imgPath string, description string) {
	_, err := db.Exec(`
		INSERT INTO Hobbies (img_path, description) VALUES ($1, $2);
	`, imgPath, description)

	if err != nil {
		log.Fatal(err)
	}
}
