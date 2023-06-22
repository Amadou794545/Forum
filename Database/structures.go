package Database

type Post struct {
	ID          int
	Title       string
	Description string
	ImagePath   string
	UserID      int
	HobbieID    int
	Likes       int
	Dislikes    int
}

type Comments struct {
	CommentID   int
	Description string
	UserID      int
	PostID      int
	Likes       int
	Dislikes    int
}

type Comments struct {
	Description string
	UserID      int
	PostID      int
}
