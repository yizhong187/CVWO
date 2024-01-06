package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/models"
)

func HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// Insert newUser into the database
	_, err = database.GetDB().Exec("INSERT INTO public.user_data (name) VALUES ($1)", newUser.Name)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created status
	json.NewEncoder(w).Encode(struct{ Message string }{"User created successfully"})
}
