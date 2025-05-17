package storage

import "go-crawler/internal/crawler"

// Storage interface for different storage implementations
type Storage interface {
	Save([]crawler.PageData) error
	Close() error
}
