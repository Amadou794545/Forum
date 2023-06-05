package main

func DeleteUser(userID int) error {
	_, err := db.Exec("DELETE FROM Users WHERE id_user = ?", userID)
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(postID int) error {
	_, err := db.Exec("DELETE FROM Posts WHERE id_post = ?", postID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(commentID int) error {
	_, err := db.Exec("DELETE FROM Comments WHERE id_comment = ?", commentID)
	if err != nil {
		return err
	}
	return nil
}
