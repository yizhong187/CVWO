package util

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/models"
)

func QueryUserType(userName string) (string, error) {
	godotenv.Load(".env")
	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		return "", errors.New("DB_USERS_TABLE is not set in the environment")
	}

	var userType string
	query := fmt.Sprintf("SELECT type FROM %s WHERE name = $1", usersTable)
	err := database.GetDB().QueryRow(query, userName).Scan(&userType)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("User not found")
		} else {
			return "", fmt.Errorf("Error retrieving user type: %v", err)
		}
	}

	return userType, nil
}

func QueryUser(userName string) (models.User, error) {
	godotenv.Load(".env")
	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		return models.User{}, errors.New("DB_USERS_TABLE is not set in the environment")
	}

	var user models.User
	query := fmt.Sprintf("SELECT id, name, type, created_at FROM %s WHERE name = $1", usersTable)
	err := database.GetDB().QueryRow(query, userName).Scan(&user.ID, &user.Name, &user.Type, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("User not found! ")
		} else {
			return models.User{}, fmt.Errorf("Error retrieving user: %v", err)
		}
	}

	return user, nil
}
