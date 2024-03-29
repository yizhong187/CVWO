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

func QueryUser(username string) (models.User, error) {
	godotenv.Load(".env")
	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		return models.User{}, errors.New("DB_USERS_TABLE is not set in the environment")
	}

	var user models.User
	query := fmt.Sprintf("SELECT id, name, type, created_at FROM %s WHERE name = $1", usersTable)
	err := database.GetDB().QueryRow(query, username).Scan(&user.ID, &user.Name, &user.Type, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("User not found! ")
		} else {
			return models.User{}, fmt.Errorf("Error retrieving user: %v", err)
		}
	}

	return user, nil
}

func QueryUserType(id string) (string, error) {
	godotenv.Load(".env")
	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		return "", errors.New("DB_USERS_TABLE is not set in the environment")
	}

	var userType string
	query := fmt.Sprintf("SELECT type FROM %s WHERE id = $1", usersTable)
	err := database.GetDB().QueryRow(query, id).Scan(&userType)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("User not found")
		} else {
			return "", fmt.Errorf("Error retrieving user type: %v", err)
		}
	}

	return userType, nil
}

func QueryUserID(username string) (string, error) {
	godotenv.Load(".env")
	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		return "", errors.New("DB_USERS_TABLE is not set in the environment")
	}

	var userID string
	query := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", usersTable)
	err := database.GetDB().QueryRow(query, username).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("User not found")
		} else {
			return "", fmt.Errorf("Error retrieving user ID: %v", err)
		}
	}

	return userID, nil
}

// [DEPRECATED]
func IsAdminOf(username string, subforumID string) (bool, error) {
	// Use QueryUser to find the user with the username
	user, err := QueryUser(username)
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

// [DEPRECATED]
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

func QueryUsername(id string) (string, error) {
	godotenv.Load(".env")
	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		return "", errors.New("DB_USERS_TABLE is not set in the environment")
	}

	var username string
	query := fmt.Sprintf("SELECT name FROM %s WHERE id = $1", usersTable)
	err := database.GetDB().QueryRow(query, id).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("User not found")
		} else {
			return "", fmt.Errorf("Error retrieving user ID: %v", err)
		}
	}

	return username, nil
}

func QueryReplyCount(threadID int) (int, error) {
	godotenv.Load(".env")
	repliesTable := os.Getenv("DB_REPLIES_TABLE")
	if repliesTable == "" {
		return -1, errors.New("DB_REPLIES_TABLE is not set in the environment")
	}

	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE thread_id = $1", repliesTable)
	err := database.GetDB().QueryRow(query, threadID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("Error retrieving reply count: %v", err)
	}

	return count, nil
}

func QueryUsernameTaken(username string) (bool, error) {
	godotenv.Load(".env")
	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		return false, errors.New("DB_USERS_TABLE is not set in the environment")
	}

	var taken bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE name = $1)", usersTable)
	err := database.GetDB().QueryRow(query, username).Scan(&taken)
	if err != nil {
		return false, fmt.Errorf("Error checking username taken: %v", err)
	}

	return taken, nil
}
