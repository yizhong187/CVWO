package handlers

import (
	"net/http"

	"github.com/yizhong187/CVWO/util"
)

func HandlerTesting(w http.ResponseWriter, r *http.Request) {

	util.RespondWithJSON(w, http.StatusOK, "OK")
}
