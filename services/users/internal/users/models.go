package users

type Users struct {
	Id             string   `gorm:"primaryKey;type:text"`
	Username       string   `gorm:"unique;not null"`
	DisplayName    string   `gorm:"not null"`
	AvatarKey      string   `gorm:"not null"`
	Bio            string   `gorm:"type:text"`
	Links          []string `gorm:"type:jsonb;serializer:json;not null"`
	IsPrivate      bool     `gorm:"not null;default:false"`
	FollowerCount  int      `gorm:"not null;default:0"`
	FollowingCount int      `gorm:"not null;default:0"`
	VideoCount     int      `gorm:"not null;default:0"`
	CreatedAt      int64    `gorm:"autoCreateTime"`
	UpdatedAt      int64    `gorm:"autoUpdateTime"`
	DeletedAt      *int64   `gorm:"index"`
}

type Follows struct {
	FollowerId string `gorm:"primaryKey;type:text;not null"`
	FolloweeId string `gorm:"primaryKey;type:text;not null"`
	CreatedAt  int64  `gorm:"autoCreateTime"`
}
