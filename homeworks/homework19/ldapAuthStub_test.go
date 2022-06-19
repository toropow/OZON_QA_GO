package homework19

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type stubLdap struct{}

func (r *stubLdap) ldapCheck(username string) bool {
	return true
}

func TestAuthStub(t *testing.T) {
	res := checkUsername("test1", new(stubLdap))
	assert.Equal(t, true, res, "Get: %v, Want: %v", "test", true)
}
