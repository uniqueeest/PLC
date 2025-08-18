package user

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Repository struct {
	filePath string
}

func UserRepository(dataDir string) *Repository {
	return &Repository{
		filePath: filepath.Join(dataDir, "users.json"),
	}
}

func (r *Repository) LoadUsers() ([]User, error) {
	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		return []User{}, nil
	}

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []User{}, nil
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) SaveUsers(users []User) error {
	dir := filepath.Dir(r.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0644)
}

