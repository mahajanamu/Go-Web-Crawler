package main

import (
	"fmt"
	"go-crawler/cmd"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("❌ No URL provided. Usage: go run main.go <URL>")
		return
	}
	fmt.Println("✅ Starting crawler...") // <-- Debug line
	cmd.Execute()
	fmt.Println("✅ Crawler finished.") // <-- Debug line
}
