package homework19

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthTable(t *testing.T) {

	t.Run("check gen func via gomock", func(t *testing.T) {
		//Arrange
		ctrl := gomock.NewController(t)
		mockLdapCheck := NewMockldap(ctrl)

		type UserLdapTest struct {
			name        string
			username    string
			expectedVal bool
		}

		tests := []UserLdapTest{
			{name: "First test", username: "Test", expectedVal: true},
			{name: "Second test", username: "test1", expectedVal: false},
			{name: "Third test", username: "Test2", expectedVal: true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockLdapCheck.EXPECT().ldapCheck(tt.username).Return(tt.expectedVal)
				res := checkUsername(tt.username, mockLdapCheck)
				assert.Equal(t, res, tt.expectedVal, "Get: %v, Want: %v", res, tt.expectedVal)
			})
		}

	})

}
