package structs

type Post struct {
	Title    string
	FilePath string
}

type Blog struct {
	Posts []Post
}

type PostData struct {
	Title string
	HTML  []byte
}
