package entity

type LikeDislike struct {
	Id        int `json:"-"`
	UserId    int `json:"userId"`
	PostId    int `json:"postId"`
	CommentId int `json:"commentId"`
}
