package storage

import (
	"encoding/csv"
	"os"

	"go-crawler/internal/crawler"
)

type CSVStorage struct {
	file   *os.File
	writer *csv.Writer
}

func NewCSVStorage(filename string) (*CSVStorage, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	writer := csv.NewWriter(file)
	// Write header
	writer.Write([]string{"URL", "Title"})
	writer.Flush()

	return &CSVStorage{
		file:   file,
		writer: writer,
	}, nil
}

func (s *CSVStorage) Save(results []crawler.PageData) error {
	for _, r := range results {
		err := s.writer.Write([]string{r.URL, r.Title})
		if err != nil {
			return err
		}
	}
	s.writer.Flush()
	return nil
}

func (s *CSVStorage) Close() error {
	s.writer.Flush()
	return s.file.Close()
}
