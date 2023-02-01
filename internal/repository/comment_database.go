package repository

import (
	"database/sql"
	"forum/internal/entity"
)

type TodoCommentDataBase struct {
	db *sql.DB
}

func NewTodoCommentDataBase(db *sql.DB) *TodoCommentDataBase {
	return &TodoCommentDataBase{db: db}
}

func (r *TodoCommentDataBase) CreateComment(user entity.User, postId int, comment string) (int, error) {
	query := `INSERT INTO comment (userId, postId, userName, comment) VALUES($1,$2,$3,$4)`
	res, err := r.db.Exec(query, user.Id, postId, user.Username, comment)
	if err != nil {
		return 0, err
	}
	commentId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(commentId), nil
}

func (r *TodoCommentDataBase) GetAllComments(postId int) ([]entity.Comment, error) {
	var comments []entity.Comment

	query := `SELECT id, userId, userName, postId, comment FROM comment WHERE postId = $1 ORDER BY id DESC`
	rows, err := r.db.Query(query, postId)
	if err != nil {
		return comments, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment entity.Comment
		err := rows.Scan(&comment.Id, &comment.UserId, &comment.Username, &comment.PostId, &comment.Comment)
		if err != nil {
			if err == sql.ErrNoRows {
				return comments, nil // чтобы не выводил ошибку если пост не комментировали
			}
			return comments, err
		}
		comments = append(comments, comment)
	}

	return comments, err
}

func (r *TodoCommentDataBase) GetCommentById(commentId int) (entity.Comment, error) {
	var comment entity.Comment
	query := `SELECT id, userId, userName, postId, comment FROM comment WHERE id = $1`
	err := r.db.QueryRow(query, commentId).Scan(&comment.Id, &comment.UserId, &comment.Username, &comment.PostId, &comment.Comment)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *TodoCommentDataBase) GetLikesByCommentId(commentId int) ([]entity.LikeDislike, error) {
	var likes []entity.LikeDislike
	query := `SELECT id, userId, postId, commentId FROM likes WHERE commentId = $1`
	rows, err := r.db.Query(query, commentId)
	if err != nil {
		return likes, err
	}
	defer rows.Close()
	for rows.Next() {
		var like entity.LikeDislike
		err := rows.Scan(&like.Id, &like.UserId, &like.PostId, &like.CommentId)
		if err != nil {
			if err == sql.ErrNoRows {
				return likes, nil // чтобы не выводил ошибку если у этого коммента нет лайков
			}
			return likes, err
		}
		likes = append(likes, like)
	}
	return likes, err
}

func (r *TodoCommentDataBase) GetDislikesByCommentId(commentId int) ([]entity.LikeDislike, error) {
	var dislikes []entity.LikeDislike
	query := `SELECT id, userId, postId, commentId FROM dislikes WHERE commentId = $1`
	rows, err := r.db.Query(query, commentId)
	if err != nil {
		return dislikes, err
	}
	defer rows.Close()
	for rows.Next() {
		var like entity.LikeDislike
		err := rows.Scan(&like.Id, &like.UserId, &like.PostId, &like.CommentId)
		if err != nil {
			if err == sql.ErrNoRows {
				return dislikes, nil // чтобы не выводил ошибку если у этого коммента нет лайков
			}
			return dislikes, err
		}
		dislikes = append(dislikes, like)
	}
	return dislikes, err
}
