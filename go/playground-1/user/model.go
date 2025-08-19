package user

import "time"

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserRequest struct {
		ID string `json:"id"`
	Name     *string `json:"name,omitempty"`
	Gender   *string `json:"gender,omitempty"`
	Age      *int `json:"age,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Age       int    `json:"age"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

type DeleteUserResponse struct {
	ID string `json:"id"`
}