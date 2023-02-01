package repository

import (
	"database/sql"
	"forum/internal/entity"
)

type TodoLikeDislikeDataBase struct {
	db *sql.DB
}

func NewTodoLikeDislikeDataBase(db *sql.DB) *TodoLikeDislikeDataBase {
	return &TodoLikeDislikeDataBase{db: db}
}

func (r *TodoLikeDislikeDataBase) LikeInDB(userId, postId, commentId int) (int, error) {
	query := `INSERT INTO likes (userId, postId, commentId) VALUES($1,$2,$3)`
	res, err := r.db.Exec(query, userId, postId, commentId)
	if err != nil {
		return 0, err
	}

	likeId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(likeId), nil
}

func (r *TodoLikeDislikeDataBase) Checklike(userId, postId, commentId int) (bool, error) {
	var userIdDB, postIdDB, commentIdDB int

	query := `SELECT userId, postId, commentId FROM likes WHERE userId = $1 AND postId = $2 AND commentId = $3`
	err := r.db.QueryRow(query, userId, postId, commentId).Scan(&userIdDB, &postIdDB, &commentIdDB)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // дальше запись лайк или дизлайк
		}
		return false, err
	}

	querydel := `DELETE FROM likes WHERE userId = $1 AND postId = $2 AND commentId = $3`
	_, err = r.db.Exec(querydel, userIdDB, postIdDB, commentIdDB)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *TodoLikeDislikeDataBase) DislikeInDB(userId, postId, commentId int) (int, error) {
	query := `INSERT INTO dislikes (userId, postId, commentId) VALUES($1,$2,$3)`
	res, err := r.db.Exec(query, userId, postId, commentId)
	if err != nil {
		return 0, err
	}

	dislikeId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(dislikeId), nil
}

func (r *TodoLikeDislikeDataBase) Checkdislike(userId, postId, commentId int) (bool, error) {
	var userIdDB, postIdDB, commentIdDB int
	query := `SELECT userId, postId, commentId FROM dislikes WHERE userId = $1 AND postId = $2 AND commentId = $3`
	err := r.db.QueryRow(query, userId, postId, commentId).Scan(&userIdDB, &postIdDB, &commentIdDB)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // дальше запись лайк или дизлайк
		}
		return false, err
	}

	querydel := `DELETE FROM dislikes WHERE  userId = $1 AND postId = $2 AND commentId = $3`
	_, err = r.db.Exec(querydel, userIdDB, postIdDB, commentIdDB)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *TodoLikeDislikeDataBase) GetLikesByUserId(userId int) ([]entity.LikeDislike, error) {
	var likes []entity.LikeDislike
	query := `SELECT id, userId, postId, commentId FROM likes WHERE userId = $1`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return likes, err
	}
	defer rows.Close()
	for rows.Next() {
		var like entity.LikeDislike
		err := rows.Scan(&like.Id, &like.UserId, &like.PostId, &like.CommentId)
		if err != nil {
			if err == sql.ErrNoRows {
				return likes, nil // чтобы не выводил ошибку если user ничего не лайкал
			}
			return likes, err
		}
		likes = append(likes, like)
	}
	return likes, nil
}

func (r *TodoLikeDislikeDataBase) GetDislikesByUserId(userId int) ([]entity.LikeDislike, error) {
	var dislikes []entity.LikeDislike
	query := `SELECT id, userId, postId, commentId FROM dislikes WHERE userId = $1`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return dislikes, err
	}
	defer rows.Close()
	for rows.Next() {
		var dislike entity.LikeDislike
		err := rows.Scan(&dislike.Id, &dislike.UserId, &dislike.PostId, &dislike.CommentId)
		if err != nil {
			if err == sql.ErrNoRows {
				return dislikes, nil // чтобы не выводил ошибку если user ничего не дизлайкал
			}
			return dislikes, err
		}
		dislikes = append(dislikes, dislike)
	}
	return dislikes, nil
}
