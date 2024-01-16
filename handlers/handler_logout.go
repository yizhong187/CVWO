package handlers

import (
	"net/http"
	"time"

	"github.com/yizhong187/CVWO/util"
)

func HandlerLogout(w http.ResponseWriter, r *http.Request) {
	// Set the token in an HTTP-only cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Path:     "/", // Make sure the cookie is sent with every request to the server
	})

	util.RespondWithJSON(w, http.StatusOK, "Successfully logged out")
}