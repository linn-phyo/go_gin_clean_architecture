package interfaces

import (
	"context"

	"github.com/linn-phyo/go_gin_clean_architecture/src/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]domain.Users, error)
	FindByID(ctx context.Context, id string) (domain.Users, error)
	Save(ctx context.Context, user domain.Users) (domain.Users, error)
	Delete(ctx context.Context, user domain.Users) error
}