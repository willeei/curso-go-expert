package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, user.Name, "John Doe")
	assert.Equal(t, user.Email, "j@j.com")
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe1", "j@j1.com", "123456")
	assert.Nil(t, err)
	assert.NotEqual(t, user.Password, "123456")
	assert.True(t, user.CheckPassword("123456"))
	assert.False(t, user.CheckPassword("1234567"))
}
