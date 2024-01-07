package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/util"
)

func HandlerTest(w http.ResponseWriter, r *http.Request) {

	userID := "1"

	type TestingUser struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}

	var testingUser TestingUser

	err := database.GetDB().QueryRow("SELECT * FROM public.testing WHERE id = $1", userID).Scan(&testingUser.ID, &testingUser.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.RespondWithError(w, http.StatusNotFound, "User not found")
		} else {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
		}
		return
	}

	// Respond with the user data in JSON format
	util.RespondWithJSON(w, http.StatusOK, testingUser)
}
