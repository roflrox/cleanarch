package domain

import "context"

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (User, error)
}
