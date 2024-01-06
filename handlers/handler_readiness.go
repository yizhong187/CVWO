package handlers

import (
	"net/http"

	"github.com/yizhong187/CVWO/util"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {

	util.RespondWithJSON(w, 200, struct{}{})
}
