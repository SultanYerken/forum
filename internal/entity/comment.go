package entity

type Comment struct {
	Id           int    `json:"-"`
	UserId       int    `json:"userId"`
	PostId       int    `json:"postId"`
	Username     string `json:"username"`
	Comment      string `json:"comment"`
	LikeCount    int
	DislikeCount int
}
