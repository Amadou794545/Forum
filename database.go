package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createUsersTable()
	createCommentsTable()
	createLikesTable()
	createPostsTable()
	createHobbiesTable()

	addHobbie("../pictures/cinema.png", "Cinema")
	addHobbie("../pictures/cuisine.png", "Cuisine")
	addHobbie("../pictures/informatique.png", "Informatique")
	addHobbie("../pictures/jeux.png", "Jeux")
	addHobbie("../pictures/lecture.png", "Lecture")
	addHobbie("../pictures/musique.png", "Musique")
	addHobbie("../pictures/sport.png", "Sport")
}

func createUsersTable() {
	_, err := db.Exec(`
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
}

func createPostsTable() {
	_, err := db.Exec(`
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
}

func createCommentsTable() {
	_, err := db.Exec(`
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
}

func createLikesTable() {
	_, err := db.Exec(`
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
}

func createHobbiesTable() {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS Hobbies (
			id_hobbie INTEGER PRIMARY KEY AUTOINCREMENT,
            img_path TEXT,
			description TEXT
        )
    `)
	if err != nil {
		log.Fatal(err)
	}
}

func addUser(email string, pseudo string, password string, imgPath string) {
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

func addPost(title string, imagePath string, description string, userID int, hobbieID int) {
	_, err := db.Exec(`
		INSERT INTO Posts (title, img_path, description, id_user, id_hobbie) VALUES ($1, $2, $3, $4);
	`, title, imagePath, description, userID, hobbieID)

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

func addHobbie(imgPath string, description string) {
	_, err := db.Exec(`
		INSERT INTO Hobbies (img_path, description) VALUES ($1, $2);
	`, imgPath, description)

	if err != nil {
		log.Fatal(err)
	}
}

func deleteUser(userID int) error {
	_, err := db.Exec("DELETE FROM Users WHERE id_user = ?", userID)
	if err != nil {
		return err
	}
	return nil
}

func deletePost(postID int) error {
	_, err := db.Exec("DELETE FROM Posts WHERE id_post = ?", postID)
	if err != nil {
		return err
	}
	return nil
}

func deleteComment(commentID int) error {
	_, err := db.Exec("DELETE FROM Comments WHERE id_comment = ?", commentID)
	if err != nil {
		return err
	}
	return nil
}

func getAllPosts() ([]Post, error) {
	rows, err := db.Query("SELECT id_post, title, description, image_path, id_user FROM Posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.ImagePath, &post.UserID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func countLikesPost(postID int) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Likes WHERE id_post = ? AND is_like = 1", postID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func countDislikesPost(postID int) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Likes WHERE id_post = ? AND is_like = -1", postID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func countLikesComment(commentID int) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Likes WHERE id_comment = ? AND is_like = 1", commentID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func countDislikesComment(commentID int) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Likes WHERE id_comment = ? AND is_like = -1", commentID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func updateUserIMG(path string, id_user int) {
	_, err := db.Exec(
		"UPDATE users SET imgPath = $1 WHERE id_user = $2", path, id_user)
	if err != nil {
		log.Fatal(err)
	}
}

func updateUsername(username string, id_user int) {
	_, err := db.Exec(
		"UPDATE users SET pseudo = $1 WHERE id_user = $2", username, id_user)
	if err != nil {
		log.Fatal(err)
	}
}
