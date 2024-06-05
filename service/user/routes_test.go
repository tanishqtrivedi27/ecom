package user

import (
	// "testing"

	"github.com/tanishqtrivedi27/ecom/types"
)

// func TestUserServiceHandlers(t *testing.T) {
// 	userStore := &mockUserStore{}

// 	// t.Run("")

// }

type mockUserStore struct {}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

	
func (m *mockUserStore) GetUserById(Id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
