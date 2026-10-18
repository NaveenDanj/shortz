package users

type Users struct {
	Id             string   `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Username       string   `gorm:"unique;not null" json:"username"`
	DisplayName    string   `gorm:"not null" json:"displayName"`
	AvatarKey      string   `gorm:"not null" json:"avatarKey"`
	Bio            string   `gorm:"type:text" json:"bio"`
	Links          []string `gorm:"type:text[]" json:"links"`
	IsPrivate      bool     `gorm:"not null;default:false" json:"isPrivate"`
	FollowerCount  int      `gorm:"not null;default:0" json:"followerCount"`
	FollowingCount int      `gorm:"not null;default:0" json:"followingCount"`
	VideoCount     int      `gorm:"not null;default:0" json:"videoCount"`
	CreatedAt      int64    `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt      int64    `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt      *int64   `gorm:"index" json:"deletedAt,omitempty"`
}
