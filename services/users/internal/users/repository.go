package users

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrNotFound = errors.New("user not found")

type Repository interface {
	FindByID(ctx context.Context, id string) (*Users, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindByID(ctx context.Context, id string) (*Users, error) {
	var user Users
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").First(&user, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
