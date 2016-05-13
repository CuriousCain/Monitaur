package models_test

import (
	"github.com/curiouscain/monitaur/models"
	"testing"
)

func TestHashPassword(t *testing.T) {
	user := models.User{
		UnsafePassword: "test",
	}

	user.HashPassword()

	if user.Password == nil {
		t.FailNow()
	}
}
