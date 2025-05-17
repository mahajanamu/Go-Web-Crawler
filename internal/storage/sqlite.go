package storage

import "go-crawler/internal/crawler"

type SQLiteStorage struct {
	// db *sql.DB
}

func NewSQLiteStorage(path string) *SQLiteStorage {
	return &SQLiteStorage{}
}

func (s *SQLiteStorage) Save(pages []crawler.PageData) error {
	// Implement SQLite logic here
	return nil
}
