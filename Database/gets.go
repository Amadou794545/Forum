package Database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetUserID(identifier string) (int, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()

	query := "SELECT id_user FROM Users WHERE (pseudo = ? OR email = ?)"
	row := db.QueryRow(query, identifier, identifier)

	var userID int
	err = row.Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil // Aucune ligne trouvée: userID = 0 sans erreur
		}
		return 0, err
	}

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

func GetUserImg(userID string) (string, error) {
	var imgPath string

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.QueryRow("SELECT imgPath FROM Users WHERE id_user = ?", userID).Scan(&imgPath)
	if err != nil {
		return "", err
	}

	return imgPath, nil
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
	query := fmt.Sprintf("SELECT id_post, title, description, imgPath, id_user, id_hobbie FROM Posts LIMIT %d OFFSET %d", limit, offset)

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
