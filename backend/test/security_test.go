package test

import (
	"testing"

	"github.com/iacopoGhilardi/amILate/pkg/security"
	"github.com/stretchr/testify/assert"
)

func TestPasswordHashing(t *testing.T) {
	plain := "mypassword"
	hash, err := security.HashPassword(plain)

	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	ok := security.CheckPasswordHash(plain, hash)
	assert.True(t, ok)
}
