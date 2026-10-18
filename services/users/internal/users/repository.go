package users

import (
	"context"
	apiContract "contract/api/users"
	"encoding/json"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// unique constraints/indexes created by the migrations
const (
	constraintUsersPK        = "users_pkey"
	constraintUsernameUnique = "users_username_lower_uidx"
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
	Create(ctx context.Context, id string, profile *apiContract.CreateUserRequest) (*Users, error)
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

func (r *repository) Create(ctx context.Context, id string, profile *apiContract.CreateUserRequest) (*Users, error) {
	links := profile.Links
	if links == nil {
		links = []string{}
	}

	user := Users{
		Id:          id,
		Username:    profile.Username,
		DisplayName: profile.DisplayName,
		Bio:         profile.Bio,
		Links:       links,
		IsPrivate:   profile.IsPrivate,
	}

	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			switch pgErr.ConstraintName {
			case constraintUsersPK:
				return nil, apiContract.ErrProfileExists
			case constraintUsernameUnique:
				return nil, apiContract.ErrUsernameTaken
			}
		}
		return nil, err
	}

	return &user, nil
}

func (r *repository) UpdateProfile(ctx context.Context, id string, profile *apiContract.UpdateUserRequest) error {

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
