package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MessageResponse struct {
	Response string `json:"response"`
}

type User struct {
	ID         int64   `json:"id"`
	FirstName  string  `json:"first_name"`
	MiddleName *string `json:"middle_name"`
	LastName   *string `json:"last_name"`
	Username   string  `json:"username"`
	Address    string  `json:"address"`
	City       string  `json:"city"`
	CreatedAt  time.Time
}

type dbS struct {
	db *sql.DB
}

func main() {

	mux := http.NewServeMux()

	dsn := "root:jujun_sql_123@tcp(127.0.0.1:3306)/latihan_4?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err.Error())
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(3 * time.Minute)

	mux.Handle("/test", myFirstApp(db))

	http.ListenAndServe(":8080", mux)

}

func myFirstApp(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(MessageResponse{Response: "Method tidak diizinkan"})
			return
		}

		req := &User{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(MessageResponse{Response: err.Error()})
			return
		}

		err := Create(req, r.Context(), db)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(MessageResponse{Response: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(MessageResponse{Response: "Success"})
	})
}

func Create(u *User, ctx context.Context, db *sql.DB) error {

	query := "INSERT INTO users (first_name, middle_name, last_name, username, address, city, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	r, err := db.ExecContext(ctx, query, u.FirstName, u.MiddleName, u.LastName, u.Username, u.Address, u.City, time.Now())

	if err != nil {
		return err
	}

	u.ID, _ = r.LastInsertId()

	return nil
}