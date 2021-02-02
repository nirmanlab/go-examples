package comment

import "time"

// Comment represents model for comment entity
type Comment struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
}

// NewComment creates comment object
func NewComment(text string) Comment {
	return Comment{
		ID:   time.Now().Unix(),
		Text: "Hello",
	}
}
