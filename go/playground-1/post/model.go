package post

import "time"

type Post struct {
	ID string `json:"id"`
	UserID string `json:"user_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreatePostRequest struct {
	Title string `json:"title"`
	Content string `json:"content"`
	UserID string `json:"user_id"`
}

type CreatePostResponse struct {
	ID string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatePostRequest struct {
	ID string `json:"id"`
	Title *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

type DeletePostRequest struct {
	ID string `json:"id"`
}