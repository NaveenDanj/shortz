package users

import (
	"context"
	apiContract "contract/api/users"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

// table users
const (
	tableUserDisplayName = "display_name"
	tableUserBio         = "bio"
	tableUserLinks       = "links"
	tableUserIsPrivate   = "is_private"
)

type Repository interface {
	FindByID(ctx context.Context, id string) (*Users, error)
	UpdateProfile(ctx context.Context, id string, profile *apiContract.UpdateUserRequest) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindByID(ctx context.Context, id string) (*Users, error) {
	var user Users
	err := r.db.WithContext(ctx).Scopes(notDeleted).First(&user, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apiContract.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) UpdateProfile(ctx context.Context, id string, profile *apiContract.UpdateUserRequest) error {
	// Updates(map) bypasses the model's serializer tag, so a []string would be
	// expanded into a SQL row. Encode links ourselves and cast to jsonb.
	links := profile.Links
	if links == nil {
		links = []string{}
	}
	linksJSON, err := json.Marshal(links)
	if err != nil {
		return err
	}

	result := r.db.WithContext(ctx).Model(&Users{}).Where("id = ?", id).Scopes(notDeleted).Updates(map[string]interface{}{
		tableUserDisplayName: profile.DisplayName,
		tableUserBio:         profile.Bio,
		tableUserLinks:       gorm.Expr("?::jsonb", string(linksJSON)),
		tableUserIsPrivate:   profile.IsPrivate,
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return apiContract.ErrNotFound
	}

	return nil
}

func notDeleted(db *gorm.DB) *gorm.DB { return db.Where("deleted_at IS NULL") }
