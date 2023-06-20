package Database

type Post struct {
	ID          int
	Title       string
	Description string
	ImagePath   string
	UserID      int
	HobbieID    int
}

type Comments struct {
	Description string
	UserID      int
	PostID      int
}
