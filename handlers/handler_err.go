package handlers

import (
	"net/http"

	"github.com/yizhong187/CVWO/util"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	util.RespondWithError(w, 400, "Something went wrong :(")
}
