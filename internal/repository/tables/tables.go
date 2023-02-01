package tables

import (
	"database/sql"
	"log"
)

const (
	user = `CREATE TABLE IF NOT EXISTS user (                   
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userName VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(100) NOT NULL
	);`
	post = `CREATE TABLE IF NOT EXISTS post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId INTEGER NOT NULL,
		userName VARCHAR(50),
		post TEXT NOT NULL,
		category TEXT,
		FOREIGN KEY (userId) REFERENCES user (id),
		FOREIGN KEY (userName) REFERENCES user (userName)
	);`
	comment = `CREATE TABLE IF NOT EXISTS comment (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId INTEGER NOT NULL,
		userName VARCHAR(50),
		postId INTEGER NOT NULL,
		comment TEXT NOT NULL,
		FOREIGN KEY (userId) REFERENCES user (id),
		FOREIGN KEY (userName) REFERENCES user (userName),
		FOREIGN KEY (postId) REFERENCES post (id)
	);`
	category = `CREATE TABLE IF NOT EXISTS category (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		categoryName VARCHAR(100) NOT NULL UNIQUE
	);`
	postCategory = `CREATE TABLE IF NOT EXISTS postCategory (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		postId INTEGER NOT NULL,
		categoryId INTEGER NOT NULL,
		FOREIGN KEY (postId) REFERENCES post (id),
		FOREIGN KEY (categoryId) REFERENCES category (id)
	);`
	likes = `CREATE TABLE IF NOT EXISTS likes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId INTEGER NOT NULL,
		postId INTEGER NOT NULL,
		commentId INTEGER NOT NULL,
		FOREIGN KEY (userId) REFERENCES user (id),
		FOREIGN KEY (postId) REFERENCES post (id),
		FOREIGN KEY (commentId) REFERENCES comment (id)
	);`
	dislikes = `CREATE TABLE IF NOT EXISTS dislikes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId INTEGER NOT NULL,
		postId INTEGER NOT NULL,
		commentId INTEGER NOT NULL,
		FOREIGN KEY (userId) REFERENCES user (id),
		FOREIGN KEY (postId) REFERENCES post (id),
		FOREIGN KEY (commentId) REFERENCES comment (id)
	);`
	session = `CREATE TABLE IF NOT EXISTS session (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId INTEGER NOT NULL,
		token TEXT,
		FOREIGN KEY (userId) REFERENCES user (id)
	);`
)

func CreateTables(db *sql.DB) {
	create(session, db)
	create(user, db)
	create(category, db)
	create(postCategory, db)
	create(post, db)
	create(comment, db)
	create(likes, db)
	create(dislikes, db)

	writeCategory(db)
}

func create(s string, db *sql.DB) error {
	query, err := db.Prepare(s)
	if err != nil {
		log.Println("Its tables", err)
		return err
	}
	query.Exec()
	return nil
}

func writeCategory(db *sql.DB) error {
	categories := []string{"human", "history", "book", "movie", "engineering", "other"}

	for i := range categories {
		query := "INSERT INTO category (categoryName) VALUES ($1)"
		_, err := db.Exec(query, categories[i])
		if err != nil {
			return err
		}
	}
	return nil
}
