package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Entry struct {
	Name string `json:"name"`
}

func addEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var e Entry
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec("INSERT INTO entries (name) VALUES (?)", e.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(e)
	}
}

func getEntries(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT name FROM entries")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var entries []Entry
	for rows.Next() {
		var e Entry
		err := rows.Scan(&e.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		entries = append(entries, e)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entries)
}

func main() {
	var err error
	db, err = sql.Open("mysql", "<username>:<password>@tcp(<RDS Endpoint>)/<database-name>")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/add", addEntry)
	http.HandleFunc("/entries", getEntries)

	http.ListenAndServe(":8080", nil)
}
