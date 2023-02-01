package repository

import (
	"database/sql"
	"forum/internal/entity"
)

type TodoPostDataBase struct {
	db *sql.DB
}

func NewTodoPostDataBase(db *sql.DB) *TodoPostDataBase {
	return &TodoPostDataBase{db: db}
}

func (r *TodoPostDataBase) GetUser(userId int) (entity.User, error) {
	var user entity.User
	query := `SELECT id, userName, email FROM user WHERE id = $1`
	err := r.db.QueryRow(query, userId).Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *TodoPostDataBase) CreatePost(user entity.User, post, category string) (int, error) {
	query := `INSERT INTO post (userId, userName, post, category) VALUES($1,$2,$3,$4)`
	res, err := r.db.Exec(query, user.Id, user.Username, post, category)
	if err != nil {
		return 0, err
	}
	postId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(postId), nil
}

func (r *TodoPostDataBase) DeletePost(postId int) error {
	query := `DELETE FROM post WHERE id = $1`
	_, err := r.db.Exec(query, postId)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoPostDataBase) GetAllPost() ([]entity.Post, error) {
	var posts []entity.Post
	query := `SELECT  id, userId, userName, post, category FROM post ORDER BY id DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return posts, err
	}
	defer rows.Close()
	for rows.Next() {
		var post entity.Post
		err := rows.Scan(&post.Id, &post.UserId, &post.Username, &post.Post, &post.Category)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, err
}

func (r *TodoPostDataBase) GetPostById(postId int) (entity.Post, error) {
	var post entity.Post
	query := `SELECT id, userId, userName, post, category FROM post WHERE id = $1`
	err := r.db.QueryRow(query, postId).Scan(&post.Id, &post.UserId, &post.Username, &post.Post, &post.Category)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (r *TodoPostDataBase) GetCategoryId(categoryName string) (int, error) {
	var id int
	query := `SELECT id FROM category WHERE categoryName = $1`
	err := r.db.QueryRow(query, categoryName).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TodoPostDataBase) WriteInPostCategory(postId, categoryId int) (int, error) {
	query := `INSERT INTO postCategory (postId, categoryId) VALUES ($1,$2)`
	res, err := r.db.Exec(query, postId, categoryId)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *TodoPostDataBase) GetPostIdByCategoryId(categoryId int) ([]int, error) {
	var postsId []int
	query := `SELECT postId FROM postCategory WHERE categoryId = $1 ORDER BY id DESC`
	rows, err := r.db.Query(query, categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var postId int
		err := rows.Scan(&postId)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil // чтобы не выдавал ошибку если в этой категории нет постов
			}
			return nil, err
		}
		postsId = append(postsId, postId)
	}
	return postsId, nil
}

func (r *TodoPostDataBase) GetPostsByUserId(userId int) ([]entity.Post, error) {
	var posts []entity.Post
	query := `SELECT id, userId, userName, post, category FROM post WHERE userId =$1 ORDER BY id DESC`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return posts, err
	}
	defer rows.Close()
	for rows.Next() {
		var post entity.Post
		err := rows.Scan(&post.Id, &post.UserId, &post.Username, &post.Post, &post.Category)
		if err != nil {
			if err == sql.ErrNoRows {
				return posts, nil // чтобы не выводил ошибку если user ничего не постил
			}
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *TodoPostDataBase) GetLikesByPostId(postId int) ([]entity.LikeDislike, error) {
	var likes []entity.LikeDislike
	query := `SELECT id, userId, postId, commentId FROM likes WHERE postId = $1 ORDER BY id DESC`
	rows, err := r.db.Query(query, postId)
	if err != nil {
		return likes, err
	}
	defer rows.Close()
	for rows.Next() {
		var like entity.LikeDislike
		err := rows.Scan(&like.Id, &like.UserId, &like.PostId, &like.CommentId)
		if err != nil {
			if err == sql.ErrNoRows {
				return likes, nil // чтобы не выводил ошибку если у этого поста нет лайков
			}
			return likes, err
		}
		likes = append(likes, like)
	}
	return likes, err
}

func (r *TodoPostDataBase) GetDislikesByPostId(postId int) ([]entity.LikeDislike, error) {
	var dislikes []entity.LikeDislike
	query := `SELECT id, userId, postId, commentId FROM dislikes WHERE postId = $1 ORDER BY id DESC`
	rows, err := r.db.Query(query, postId)
	if err != nil {
		return dislikes, err
	}
	defer rows.Close()
	for rows.Next() {
		var dislike entity.LikeDislike
		err := rows.Scan(&dislike.Id, &dislike.UserId, &dislike.PostId, &dislike.CommentId)
		if err != nil {
			if err == sql.ErrNoRows {
				return dislikes, nil // чтобы не выводил ошибку если у этого поста нет лайков
			}
			return dislikes, err
		}
		dislikes = append(dislikes, dislike)
	}
	return dislikes, err
}
