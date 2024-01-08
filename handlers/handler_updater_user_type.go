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

func HandlerUpdateUserType(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		log.Fatal("usersTable is not set in the environment")
	}

	adminsTable := os.Getenv("DB_ADMINS_TABLE")
	if adminsTable == "" {
		log.Fatal("adminsTable is not set in the environment")
	}

	// Query the database for the user
	userName := chi.URLParam(r, "name")
	user, err := util.QueryUser(userName)
	if err != nil {
		if err.Error() == "User not found" {
			util.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Check if user has "superuser" type to edit user type
	if user.Type != "superuser" {
		util.RespondWithError(w, http.StatusForbidden, "User does not have the required permissions")
		return
	}

	// Decode the JSON request body into EditUserTypeRequestData struct
	type CreateSubforumRequestData struct {
		Name       string `json:"name"`
		NewType    string `json:"newType"`
		SubforumID int    `json:"subforumID"`
	}
	var requestData CreateSubforumRequestData
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Execute the update query at UsersTable
	updateQuery := fmt.Sprintf("UPDATE %s SET type = $1 WHERE name = $2", usersTable)
	_, err = database.GetDB().Exec(updateQuery, requestData.NewType, requestData.Name)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error updating user: \n%v", err))
		return
	}

	//Execute the update query at AdminsTable
	if requestData.NewType == "admin" {
		// Promote to admin: Insert into AdminsTable
		insertQuery := fmt.Sprintf("INSERT INTO %s (admin_id, subforum_id) VALUES ($1, $2)", adminsTable)
		_, err = database.GetDB().Exec(insertQuery, user.ID, requestData.SubforumID)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error promoting user to admin: \n%v", err))
			return
		}
	} else {
		// Demote from admin: Delete from AdminsTable
		deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE admin_id = $1 and subforum_id = $2)", adminsTable)
		_, err = database.GetDB().Exec(deleteQuery, user.ID, requestData.SubforumID)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error demoting user from admin: \n%v", err))
			return
		}
	}

	// Respond with a success message
	if requestData.NewType == "admin" {
		util.RespondWithJSON(w, http.StatusOK, struct{ Message string }{"User promoted successfully"})
	} else {
		util.RespondWithJSON(w, http.StatusOK, struct{ Message string }{"User demoted successfully"})
	}

}
