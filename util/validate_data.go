package util

import (
	"fmt"

	"github.com/yizhong187/CVWO/models"
)

func UseCreationValidation(user models.User) error {
	if user.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}
