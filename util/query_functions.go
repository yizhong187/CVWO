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

func QueryUserID(userName string) (string, error) {
	godotenv.Load(".env")
	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		return "", errors.New("DB_USERS_TABLE is not set in the environment")
	}

	var userID string
	query := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", usersTable)
	err := database.GetDB().QueryRow(query, userName).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("User not found")
		} else {
			return "", fmt.Errorf("Error retrieving user ID: %v", err)
		}
	}

	return userID, nil
}

func IsAdminOf(userName string, subforumID string) (bool, error) {
	// Use QueryUser to find the user with the username
	user, err := QueryUser(userName)
	if err != nil {
		return false, err
	}

	// Check user type
	switch user.Type {
	case "normal": // Normal users are not admins
		return false, nil
	case "superuser": // Superusers are admins of all subforums
		return true, nil
	case "admin": // Check if the user is an admin of the specific subforum
		return checkAdminOfSubforum(user.ID, subforumID)
	default:
		return false, errors.New("unrecognized user type")
	}
}

func checkAdminOfSubforum(userID, subforumID string) (bool, error) {
	godotenv.Load(".env")
	adminsTable := os.Getenv("DB_ADMINS_TABLE")
	if adminsTable == "" {
		return false, errors.New("DB_ADMINS_TABLE is not set in the environment")
	}

	// Query to check if a specific record exists in the admin table
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE admin_id = $1 AND subforum_id = $2)", adminsTable)
	var exists bool
	err := database.GetDB().QueryRow(query, userID, subforumID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("Error querying admin table: %v", err)
	}

	return exists, nil
}
