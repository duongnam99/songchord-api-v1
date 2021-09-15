package models

type Song struct {
	Title    string
	Content  string
	Author   string
	Category string
	// Comment  []Comment
}

type Comment struct {
	Name    string
	Email   string
	Content string
}
