package entity

type Post struct {
	Id           int    `json:"-"`
	UserId       int    `json:"userId"`
	Username     string `json:"username"`
	Post         string `json:"post"`
	Category     string
	LikeCount    int
	DislikeCount int
}
