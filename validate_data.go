package main

import (
	"fmt"
)

func useCreationValidation(user User) error {
	if user.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}
