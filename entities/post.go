package entities

type Post struct {
	ID 		int 		`json:"id"`
	Title string 	`json:"title"`
	Text 	string 	`json:"text"`
}