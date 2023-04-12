package usecase

import (
	"clean-arch/internal/entity"
	"clean-arch/internal/infrastructure/repository"
	"clean-arch/internal/usecase/apperror"
	"context"
	"errors"
)

type UserUseCase struct {
	repo UserRepository
}

func NewUserUseCase(repo UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	_, err := uc.repo.SelectByLogin(ctx, user.Email, user.Username)
	if err != nil && !errors.Is(err, repository.ErrObjectNotFound) {
		return nil, err
	}
	if !errors.Is(err, repository.ErrObjectNotFound) {
		return nil, apperror.AppError{"Email or username already taken.", apperror.ErrorTypeInvalidRequest}
	}

	err = user.HashPassword()
	if err != nil {
		return nil, err
	}

	id, err := uc.repo.Insert(ctx, user)
	if err != nil {
		return nil, err
	}

	u, err := uc.repo.SelectById(ctx, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uc *UserUseCase) GetById(ctx context.Context, id int64) (*entity.User, error) {
	u, err := uc.repo.SelectById(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return nil, apperror.AppError{"User not found.", apperror.ErrorTypeNotFound}
		}
		return nil, err
	}

	return u, nil
}

func (uc *UserUseCase) Delete(ctx context.Context, id int64) error {
	_, err := uc.repo.SelectById(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return apperror.AppError{"User not found.", apperror.ErrorTypeNotFound}
		}
		return err
	}

	err = uc.repo.Delete(ctx, id)
	return err
}
