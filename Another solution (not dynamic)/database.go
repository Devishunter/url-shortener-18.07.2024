package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "yourusername:0000@tcp(127.0.0.1:3306)/urlshortener")

	if err != nil {
		log.Fatalf("Could not open database: %s\n", err)
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS urls (
        id INT AUTO_INCREMENT PRIMARY KEY,
        original_url TEXT NOT NULL,
        short_url VARCHAR(255) NOT NULL UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf("Could not create table: %s\n", err)
	}
}

func ShortenURL(originalURL string) (string, error) {
	shortURL := generateShortURL()
	_, err := db.Exec("INSERT INTO urls (original_url, short_url) VALUES (?, ?)", originalURL, shortURL)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func ExpandURL(shortURL string) (string, error) {
	var originalURL string
	err := db.QueryRow("SELECT original_url FROM urls WHERE short_url = ?", shortURL).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("original URL not found for short URL: %s", shortURL)
		}
		return "", err
	}
	return originalURL, nil
}

func generateShortURL() string {
	return time.Now().Format("20060102150405")
}
