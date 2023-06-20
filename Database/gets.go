package Database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
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

func GetUserPosts(user_id int) ([]Post, error) {
	posts := make([]Post, 0)
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id_post, title, description, imgPath, id_user, id_hobbie FROM posts WHERE id_user = $1", user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func GetPosts(offset, limit int, filters []int) ([]Post, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Construct the base SQL query with pagination
	query := "SELECT id_post, title, description, imgPath, id_user, id_hobbie FROM Posts"

	// Create a placeholder string for the filter values
	filterPlaceholders := make([]string, len(filters))
	filterArgs := make([]interface{}, len(filters))
	for i, filter := range filters {
		filterPlaceholders[i] = "?"
		filterArgs[i] = filter
	}

	// Add WHERE clause to filter by selected hobbies if filters are provided
	if len(filters) > 0 {
		// Join the filter placeholders with commas
		filterValues := strings.Join(filterPlaceholders, ", ")

		// Add the WHERE clause to the query
		query += " WHERE id_hobbie IN (" + filterValues + ")"
	}

	// Add ORDER BY and LIMIT/OFFSET clauses for pagination
	query += fmt.Sprintf(" ORDER BY id_post DESC LIMIT %d OFFSET %d", limit, offset)

	rows, err := db.Query(query, filterArgs...)
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

func GetComment(postID string) ([]Comments, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT description, id_user, id_post FROM Comments WHERE id_Post = $1", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comments{}
	for rows.Next() {
		var comment Comments
		if err := rows.Scan(&comment.Description, &comment.UserID, &comment.PostID); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
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
