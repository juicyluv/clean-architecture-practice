package postgres

import (
	"clean-arch/internal/entity"
	"clean-arch/internal/infrastructure/repository"
	"clean-arch/pkg/postgres"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepository(pg *postgres.Postgres) *UserRepository {
	return &UserRepository{pg}
}

func (r *UserRepository) Insert(ctx context.Context, user entity.User) (int64, error) {
	var id int64

	row := r.Pool.QueryRow(ctx, `
		INSERT INTO users (username, email, password) 
		VALUES($1, $2, $3) 
		RETURNING id`,
		user.Username,
		user.Email,
		user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, errors.Wrap(err, "inserting user")
	}

	return id, nil
}

func (r *UserRepository) SelectById(ctx context.Context, id int64) (*entity.User, error) {
	var user entity.User

	row := r.Pool.QueryRow(ctx, `
		SELECT id, username, email, created_at
		FROM users
		WHERE id = $1`, id)
	if err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrObjectNotFound
		}
		return nil, errors.Wrapf(err, "scanning user %d", id)
	}

	return &user, nil
}

func (r *UserRepository) SelectByLogin(ctx context.Context, email string, username string) (*entity.User, error) {
	var user entity.User

	row := r.Pool.QueryRow(ctx, `
		SELECT id, username, email, created_at
		FROM users
		WHERE lower(username) = lower($1) 
		   OR lower(email) = lower($2)`, username, email)
	if err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrObjectNotFound
		}
		return nil, errors.Wrapf(err, "scanning user")
	}

	return &user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.Pool.Query(ctx, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return errors.Wrapf(err, "deleting user %d", id)
	}

	return nil
}
