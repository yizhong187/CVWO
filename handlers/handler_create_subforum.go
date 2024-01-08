package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"

	"github.com/yizhong187/CVWO/util"
)

func HandlerCreateSubforum(w http.ResponseWriter, r *http.Request) {

	userName := chi.URLParam(r, "name")

	godotenv.Load(".env")
	subforumTable := os.Getenv("DB_SUBFORUM_TABLE")
	if subforumTable == "" {
		log.Fatal("DB_SUBFORUM_TABLE is not set in the environment")
	}

	type CreateSubforumRequestData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	var requestData CreateSubforumRequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	query := fmt.Sprintf("INSERT INTO %s (name, description, created_by, created_at, updated_by) VALUES ($1, $2, $3, NOW(), $3)", subforumTable)
	_, err = database.GetDB().Exec(query, requestData.Name, requestData.Description, userName)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, struct{ Message string }{"Subforum created successfully"})
}
