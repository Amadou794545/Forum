package Database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GetUserID(identifier string) (string, error) {
	db, err := sql.Open("sqlite3", "./database.db")
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

func GetPostsDislikes(PostID int) (int, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()
	var counter int
	err = db.QueryRow("SELECT COUNT(*) FROM PostsDislikes WHERE id_post = $1", PostID).Scan(&counter)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

func GetCommentsLikes(PostID int) (int, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()
	var counter int
	err = db.QueryRow("SELECT COUNT(*) FROM CommentsLikes WHERE id_comment = $1", PostID).Scan(&counter)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

func GetCommentsDislikes(PostID int) (int, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()
	var counter int
	err = db.QueryRow("SELECT COUNT(*) FROM CommentsDislikes WHERE id_comment = $1", PostID).Scan(&counter)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

func GetPostsLikes(PostID int) (int, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var counter int
	err = db.QueryRow("SELECT COUNT(*) FROM PostsLikes WHERE id_post = $1", PostID).Scan(&counter)
	if err != nil {
		return 0, err
	}

	return counter, nil
}

func GetUserUsername(userID string) (string, error) {
	var db *sql.DB
	var err error

	db, err = sql.Open("sqlite3", "./database.db")
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
	db, err := sql.Open("sqlite3", "./database.db")
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

func GetPosts(offset, limit int, filters []int) ([]Post, error) {
	db, err := sql.Open("sqlite3", "./database.db")
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

func GetUserPosts(user_id int) ([]Post, error) {
	posts := make([]Post, 0)
	db, err := sql.Open("sqlite3", "./database.db")
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

<<<<<<< HEAD
=======
func GetUserLikedPosts(user_id int) ([]Post, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id_post FROM PostsLikes WHERE id_user = $1", user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likedPosts []int
	for rows.Next() {
		var id_post int
		if err := rows.Scan(&id_post); err != nil {
			return nil, err
		}
		likedPosts = append(likedPosts, id_post)
	}

	// Query the posts using the likedPosts IDs
	var likedPostIDs []string
	for _, id := range likedPosts {
		likedPostIDs = append(likedPostIDs, strconv.Itoa(id))
	}

	query := fmt.Sprintf("SELECT * FROM Posts WHERE id IN (%s)", strings.Join(likedPostIDs, ","))
	rows, err = db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likedPostsList []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.ImagePath, &post.UserID, &post.HobbieID); err != nil {
			return nil, err
		}
		likedPostsList = append(likedPostsList, post)
	}

	return likedPostsList, nil
}

>>>>>>> 38d9ecbcf868f56be88f6cdd6c27dccd2a8cf89e
func GetComment(postID string) ([]Comments, error) {
	db, err := sql.Open("sqlite3", "./database.db")
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
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	//TODO
	db.Close()
}
