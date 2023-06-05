package main

func GetAllPosts() ([]Post, error) {
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

	return posts, nil
}

func GetPostByUser(userID int) {
	//TODO
}

func GetPostLikedByUser(userID int) {
	//TODO
}
