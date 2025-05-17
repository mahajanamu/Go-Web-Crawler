# 🕷️ Go Web Crawler

A simple yet powerful **concurrent web crawler** written in Go. This tool crawls websites starting from a given URL, traverses links up to a specified depth, extracts page titles, and saves the results to a CSV file.

---

## 📌 Features

- 🔗 Crawl any website from a start URL
- ⚡ Concurrent crawling with goroutines
- 🎯 Depth-based traversal to control recursion level
- 🔐 Prevents duplicate visits with a thread-safe visited map
- 🧠 Extracts `<title>` tags from HTML pages
- 📄 Saves crawl results (URL + Title) to `output.csv`
- 📦 Modular architecture using idiomatic Go packages
- 🛠️ CLI-based usage for flexibility

### ✅ Prerequisites

- Go 1.19 or higher
- Internet connection

### 📦 Install Dependencies

```bash
go mod tidy
```

### Run the crawler

```bash
go run main.go https://example.com 2
```

## 📌 To-Do (Improvements)

- [ ] Respect `robots.txt`
- [ ] Add rate limiting to avoid server overload
- [ ] Add SQLite or JSON output formats
- [ ] Add domain filter or link validator
- [ ] Add unit and integration tests

---

## 🧑‍💻 Author

**Amey Mahajan**  
[LinkedIn](https://www.linkedin.com/in/mahajanamu/)
