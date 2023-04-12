package usecase

import (
	"clean-arch/internal/entity"
	"context"
)

type (
	User interface {
		Create(ctx context.Context, user entity.User) (*entity.User, error)
		GetById(ctx context.Context, id int64) (*entity.User, error)
		Delete(ctx context.Context, id int64) error
	}

	UserRepository interface {
		Insert(ctx context.Context, user entity.User) (int64, error)
		SelectById(ctx context.Context, id int64) (*entity.User, error)
		SelectByLogin(ctx context.Context, email string, username string) (*entity.User, error)
		Delete(ctx context.Context, id int64) error
	}
)
