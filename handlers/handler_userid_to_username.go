package handlers

import (
	"net/http"

	"github.com/yizhong187/CVWO/util"
)

func HandlerUseridToUsername(w http.ResponseWriter, r *http.Request) {

	// Extract the ID from the query string
	userID := r.URL.Query().Get("id")
	if userID == "" {
		util.RespondWithError(w, http.StatusBadRequest, "ID is required")
		return
	}

	username, err := util.QueryUsername(userID)
	if err != nil {
		if err.Error() == "User not found!" {
			util.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Respond with the user data in JSON format
	util.RespondWithJSON(w, http.StatusOK, username)
}
