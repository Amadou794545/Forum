package main

func CountLikesPost(postID int) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Likes WHERE id_post = ? AND is_like = 1", postID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountDislikesPost(postID int) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Likes WHERE id_post = ? AND is_like = -1", postID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountLikesComment(commentID int) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Likes WHERE id_comment = ? AND is_like = 1", commentID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountDislikesComment(commentID int) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Likes WHERE id_comment = ? AND is_like = -1", commentID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
