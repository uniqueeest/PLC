package user

import (
	"time"

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