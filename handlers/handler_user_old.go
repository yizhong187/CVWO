package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yizhong187/CVWO/util"
)

func HandlerUserOld(w http.ResponseWriter, r *http.Request) {

	userName := chi.URLParam(r, "name")
	if userName == "" {
		util.RespondWithError(w, http.StatusBadRequest, "User's name is required")
		return
	}

	user, err := util.QueryUser(userName)
	if err != nil {
		if err.Error() == "User not found!" {
			util.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Respond with the user data in JSON format
	util.RespondWithJSON(w, http.StatusOK, user)
}
