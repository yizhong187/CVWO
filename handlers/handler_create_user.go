package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/models"
	"github.com/yizhong187/CVWO/util"
)

func HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	// Check if request is a POST request
	if r.Method != http.MethodPost {
		util.RespondWithError(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Error parsing request body")
		return
	}

	// Insert newUser into the database
	_, err = database.GetDB().Exec("INSERT INTO public.user_data (name) VALUES ($1)", newUser.Name)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Respond with success message
	util.RespondWithJSON(w, http.StatusCreated, struct{ Message string }{"User created successfully"})
}
