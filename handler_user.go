package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func handlerUser(w http.ResponseWriter, r *http.Request) {
	// Assuming you're retrieving user by an ID passed as a query parameter
	// http://www.example.com/user?id=123

	//userID := r.URL.Query().Get("id")
	userID := "1"
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	var user User
	err := db.QueryRow("SELECT id, name FROM public.user_data WHERE id = $1", userID).Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Respond with the user data in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
