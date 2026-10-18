package users

type User struct {
	ID             string   `json:"id"`
	Username       string   `json:"username"`
	DisplayName    string   `json:"displayName"`
	Bio            string   `json:"bio"`
	AvatarKey      string   `json:"avatarKey"`
	FollowerCount  int      `json:"followerCount"`
	FollowingCount int      `json:"followingCount"`
	VideoCount     int      `json:"videoCount"`
	Links          []string `json:"links"`
	IsPrivate      bool     `json:"isPrivate"`
	CreatedAt      int64    `json:"createdAt"`
	UpdatedAt      int64    `json:"updatedAt"`
}
