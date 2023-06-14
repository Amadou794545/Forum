package Database

import (
	"database/sql"
	"fmt"
	"log"
)

func GetUserID(identifier string) (string, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return "", err
	}

	query := "SELECT id_user FROM Users WHERE (pseudo = ? OR email = ?)"
	row := db.QueryRow(query, identifier, identifier)

	var userID string
	err = row.Scan(&userID)
	if err != nil {
		return "", err
	}

	db.Close()

	return userID, nil
}

func GetUserUsername(userID string) (string, error) {
	var db *sql.DB
	var err error

	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		return "", err
	}
	defer db.Close()

	var username string
	err = db.QueryRow("SELECT pseudo FROM Users WHERE id_user=?", userID).Scan(&username)
	if err != nil {
		return "", err
	}

	return username, nil
}

func GetAllPosts() ([]Post, error) {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id_post, title, description, image_path, id_user FROM Posts, id_hobbie FROM Hobbies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.ImagePath, &post.UserID, &post.HobbieID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	db.Close()

	return posts, nil
}

func GetPostByUser(userID int) {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	//TODO

	db.Close()
}

func GetPosts(offset, limit int) ([]Post, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Construct the SQL query with pagination
	query := fmt.Sprintf("SELECT id_post, title, description, imgPath, id_user, id_hobbie FROM Posts ORDER BY id_post DESC LIMIT %d OFFSET %d", limit, offset)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}

	// Iterate through the rows and create Post objects
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.ImagePath, &post.UserID, &post.HobbieID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("test ko")
		return nil, err
	}

	return posts, nil
}

func GetPostLikedByUser(userID int) {
	var db *sql.DB

	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	//TODO

	db.Close()
}
