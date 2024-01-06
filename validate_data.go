package main

import (
	"fmt"

	"github.com/yizhong187/CVWO/models"
)

func useCreationValidation(user models.User) error {
	if user.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}
