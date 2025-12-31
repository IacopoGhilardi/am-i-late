package test

import (
	"testing"

	"fmt"

	"github.com/iacopoGhilardi/amILate/pkg/security"
	"github.com/stretchr/testify/assert"
)

func TestPasswordHashing(t *testing.T) {
	plain := "mypassword"
	hash, err := security.HashPassword(plain)
	fmt.Println(hash)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	ok := security.CheckPasswordHash(plain, hash)
	assert.True(t, ok)
}
