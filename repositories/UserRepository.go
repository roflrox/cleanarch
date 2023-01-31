package repositories

import (
	"cleanarch/domain"
	"context"
	"database/sql"
)

type userRepo struct {
	db *sql.DB
}

func (u *userRepo) GetByID(ctx context.Context, id int64) (domain.User, error) {
	query := `SELECT id, name, created_at, updated_at FROM users where id=?`
	return u.getOne(ctx, query, id)
}

func (u *userRepo) getOne(ctx context.Context, query string, args ...interface{}) (domain.User, error) {

	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return domain.User{}, err
	}
	row := stmt.QueryRowContext(ctx, args...)

	res := domain.User{}

	row.Scan(
		&res.ID,
		&res.Name,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	return res, nil
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepo{db: db}
}
