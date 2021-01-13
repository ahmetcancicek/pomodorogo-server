package auth

import (
	"context"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
)

// UserService represent the user's service
type Service interface {
	FindByID(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Save(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
}
