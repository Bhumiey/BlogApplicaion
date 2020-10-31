package Code

type Blog struct {
	BlogTitle    string   `json:"blogTitle"`
	BlogAuthor   Author   `json:"blogAuthor"`
	BlogCategory Category `json:"blogCategory"`
	BlogContent  string   `json:"blogContent"`
}
type Author struct {
	AuthorName string `json:"authorName"`
	AuthorID   string `json:"authorId"`
}
type Category struct {
	CategoryID   string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}
type BlogList struct {
	Blogs []Blog `json:"blogList"`
}
