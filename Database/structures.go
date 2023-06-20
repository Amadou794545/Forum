package Database

type Comments struct {
	Description string
	UserID      int
	PostID      int
}
type Post struct {
	ID          int
	Title       string
	Description string
	ImagePath   string
	UserID      int
	HobbieID    int
	//description Comments
}
