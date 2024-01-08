package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/models"
	"github.com/yizhong187/CVWO/util"
)

// HandlerAllSubforums handles the request to retrieve all subforums from the database.
func HandlerAllSubforums(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	subforumsTable := os.Getenv("DB_SUBFORUMS_TABLE")
	query := fmt.Sprintf("SELECT * FROM %s", subforumsTable)

	// Execute sql query and return a rows result set
	rows, err := database.GetDB().Query(query)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}
	defer rows.Close()

	// Initialise a slice of subforums, scan each row into a subforum struct and append into the slice
	var subforums []models.Subforum
	for rows.Next() {
		var subforum models.Subforum
		err := rows.Scan(&subforum.ID, &subforum.Name, &subforum.Description, &subforum.CreatedBy, &subforum.CreatedAt, &subforum.UpdatedAt)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error scanning row: \n%v", err))
			return
		}
		subforums = append(subforums, subforum)
	}

	if err = rows.Err(); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Row error: \n%v", err))
		return
	}

	// Respond with list of subforums in JSON format
	util.RespondWithJSON(w, http.StatusOK, subforums)
}
