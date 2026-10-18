package users

import (
	"context"

	apiUsers "contract/api/users"
)

type Service interface {
	GetMe(ctx context.Context, userID string) (*apiUsers.User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetMe(ctx context.Context, userID string) (*apiUsers.User, error) {
	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	links := user.Links
	if links == nil {
		links = []string{}
	}

	return &apiUsers.User{
		ID:             user.Id,
		Username:       user.Username,
		DisplayName:    user.DisplayName,
		Bio:            user.Bio,
		AvatarKey:      user.AvatarKey,
		FollowerCount:  user.FollowerCount,
		FollowingCount: user.FollowingCount,
		VideoCount:     user.VideoCount,
		Links:          links,
		IsPrivate:      user.IsPrivate,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}, nil
}
