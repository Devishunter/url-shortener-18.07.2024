package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestShortenHandler(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	origDB := db
	db = mockDB
	defer func() { db = origDB }()

	mock.ExpectExec("INSERT INTO urls").WithArgs("http://example.com", sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

	reqBody := ShortenRequest{URL: "http://example.com"}
	body, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/shorten", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(shortenHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
	var resp ShortenResponse
	json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NotEmpty(t, resp.ShortURL, "Expected non-empty short URL")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRedirectHandler(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	origDB := db
	db = mockDB
	defer func() { db = origDB }()

	mock.ExpectQuery("SELECT original_url FROM urls WHERE short_url = ?").
		WithArgs("shortURL123").
		WillReturnRows(sqlmock.NewRows([]string{"original_url"}).AddRow("http://example.com"))

	req, err := http.NewRequest("GET", "/shortURL123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/{shortURL}", redirectHandler).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusFound, rr.Code, "Expected status Found")
	assert.Equal(t, "http://example.com", rr.Header().Get("Location"), "Expected redirect to original URL")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
