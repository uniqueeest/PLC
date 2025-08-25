package user

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Service 사용자 서비스
type Service struct {
	repo *Repository
}

func UserService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(req CreateUserRequest) (CreateUserResponse, error) {
	users, err := s.repo.LoadUsers()
	if err != nil {
		return CreateUserResponse{}, err
	}

	// 새 사용자 생성
	newUser := User{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Gender:    req.Gender,
		Age:       req.Age,
		Email:     req.Email,
		Password:  req.Password, // TODO: 해시로 변경해야 함
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	users = append(users, newUser)

	err = s.repo.SaveUsers(users)
	if err != nil {
		return CreateUserResponse{}, err
	}

	return CreateUserResponse{
		ID:        newUser.ID,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}, nil
}

func (s *Service) GetUser(id string) (User, error) {
	users, err := s.repo.LoadUsers()
	if err != nil {
		return User{}, err
	}

	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return User{}, fmt.Errorf("user not found")
}

func (s *Service) UpdateUser(id string, req UpdateUserRequest) (UserResponse, error) {
	users, err := s.repo.LoadUsers()
	if err != nil {
		return UserResponse{}, err
	}

	now := time.Now().UTC()

	for i := range users {
		if users[i].ID != id {
			continue
	}

		if req.Name != nil {users[i].Name = *req.Name}
		if req.Gender != nil {users[i].Gender = *req.Gender}
		if req.Age != nil {users[i].Age = *req.Age}
		if req.Email != nil {users[i].Email = *req.Email}
		if req.Password != nil {users[i].Password = *req.Password}

		users[i].UpdatedAt = now

		err = s.repo.SaveUsers(users)
		if err != nil {
			return UserResponse{}, err
		}
		
		// 업데이트된 사용자 정보 반환
		return UserResponse{
			ID:        users[i].ID,
			Name:      users[i].Name,
			Gender:    users[i].Gender,
			Age:       users[i].Age,
			Email:     users[i].Email,
			CreatedAt: users[i].CreatedAt,
			UpdatedAt: users[i].UpdatedAt,
		}, nil
	}

	return UserResponse{}, fmt.Errorf("user not found")
}

func (s *Service) DeleteUser(id string) (DeleteUserResponse, error) {
	users, err := s.repo.LoadUsers()
	if err != nil {
		return DeleteUserResponse{}, err
	}

	for i := range users {
		if users[i].ID == id {
			users = append(users[:i], users[i+1:]...)
			
			err = s.repo.SaveUsers(users)
			if err != nil {
				return DeleteUserResponse{}, err
			}

			return DeleteUserResponse{
				ID: id,
			}, nil
		}
	}

	return DeleteUserResponse{}, fmt.Errorf("user not found")
}

// ValidateUser 이메일과 비밀번호로 사용자 검증
func (s *Service) ValidateUser(email, password string) (*UserResponse, error) {
	users, err := s.repo.LoadUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Email == email && user.Password == password {
			return &UserResponse{
				ID: user.ID,
				Name: user.Name,
				Gender: user.Gender,
				Age: user.Age,
				Email: user.Email,
			}, nil
		}
	}

	return nil, fmt.Errorf("invalid email or password")
}

func CreateToken(userid string) (string, error) {
  var err error
  //Creating Access Token
  os.Setenv("ACCESS_SECRET", "jwt-secret")
  atClaims := jwt.MapClaims{}
  atClaims["authorized"] = true
  atClaims["user_id"] = userid
  atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
  if err != nil {
     return "", err
  }
  return token, nil
}

func (s *Service) Login(req LoginRequest) (LoginResponse, error) {
	user, err := s.ValidateUser(req.Email, req.Password)
	if err != nil {
		return LoginResponse{}, err
	}

	token, err := CreateToken(user.ID)
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		Token:     token,
		User:      *user,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}, nil
}