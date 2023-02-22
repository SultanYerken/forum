package repository

import (
	"database/sql"
	"log"
)

func NewSQLiteDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./internal/usecase/repo/database.db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	
	if err = db.Ping(); err != nil {
		log.Println("ping:", err)
		return nil, err
	}
	
	forTableConnection(db)
	return db, nil
}

// FOREIGN KEY doesn't work without it...
func forTableConnection(db *sql.DB) error {
	query := `PRAGMA foreign_key=1`
	res, err := db.Prepare(query)
	if err != nil {
		return err
	}
	res.Exec()
	return nil
}
