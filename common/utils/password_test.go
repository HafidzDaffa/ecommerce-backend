package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/ecommerce-go-vue/backend/common/utils"
)

func TestHashPassword(t *testing.T) {
	password := "mypassword123"

	hash, err := utils.HashPassword(password)

	assert.NoError(t, err, "HashPassword should not return error")
	assert.NotEmpty(t, hash, "Hash should not be empty")
	assert.NotEqual(t, password, hash, "Hash should not be equal to plain password")
}

func TestCheckPassword(t *testing.T) {
	password := "mypassword123"

	hash, err := utils.HashPassword(password)
	assert.NoError(t, err)

	isValid := utils.CheckPassword(password, hash)

	assert.True(t, isValid, "CheckPassword should return true for correct password")
}

func TestCheckPasswordWrongPassword(t *testing.T) {
	password := "mypassword123"
	wrongPassword := "wrongpassword"

	hash, err := utils.HashPassword(password)
	assert.NoError(t, err)

	isValid := utils.CheckPassword(wrongPassword, hash)

	assert.False(t, isValid, "CheckPassword should return false for wrong password")
}

func TestCheckPasswordInvalidHash(t *testing.T) {
	password := "mypassword123"
	invalidHash := "invalidhash123"

	isValid := utils.CheckPassword(password, invalidHash)

	assert.False(t, isValid, "CheckPassword should return false for invalid hash")
}

func TestHashPasswordDifferentResults(t *testing.T) {
	password := "mypassword123"

	hash1, err1 := utils.HashPassword(password)
	hash2, err2 := utils.HashPassword(password)

	assert.NoError(t, err1)
	assert.NoError(t, err2)

	assert.NotEqual(t, hash1, hash2, "Hashing the same password should produce different hashes due to salt")
}

func TestCheckPasswordBothHashes(t *testing.T) {
	password := "mypassword123"

	hash1, _ := utils.HashPassword(password)
	hash2, _ := utils.HashPassword(password)

	isValid1 := utils.CheckPassword(password, hash1)
	isValid2 := utils.CheckPassword(password, hash2)

	assert.True(t, isValid1, "First hash should validate correctly")
	assert.True(t, isValid2, "Second hash should validate correctly")
}
