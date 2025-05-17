package cmd

import (
	"fmt"
	"go-crawler/internal/crawler"
	"go-crawler/internal/storage"
	"log"
	"os"
	"strconv"
)

func Execute() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: go run main.go <startURL> <depth>")
	}

	startURL := os.Args[1]
	depth, err := strconv.Atoi(os.Args[2])
	if err != nil || depth < 1 {
		log.Fatalf("Depth must be a positive integer")
	}

	results := crawler.StartCrawling(startURL, depth)

	csvStorage, err := storage.NewCSVStorage("output.csv")
	if err != nil {
		log.Fatalf("Failed to initialize CSV storage: %v", err)
	}

	err = csvStorage.Save(results)
	if err != nil {
		log.Fatalf("Failed to save results: %v", err)
	}
	fmt.Println("Crawling completed. Output saved to output.csv")
}
