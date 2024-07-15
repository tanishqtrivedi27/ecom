package user

import (
	"database/sql"
	"fmt"

	"github.com/tanishqtrivedi27/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES ($1, $2, $3, $4)", user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetAddresses(userId int) ([]*types.Address, error) {
	rows, err := s.db.Query("SELECT * FROM addresses WHERE userId = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	addresses := make([]*types.Address, 0)
	for rows.Next() {
		a, err := scanRowsIntoAddress(rows)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, a)
	}

	return addresses, nil
}

func (s *Store) CreateAddress(address types.Address) error {
	_, err := s.db.Exec("INSERT INTO addresses(userId, line1, line2, city, country) VALUES ($1, $2, $3, $4, $5)", address.UserID, address.Line1, address.Line2, address.City, address.Country)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateLastLogin(userId int) error {
	_, err := s.db.Exec("UPDATE users SET lastLogin = NOW() WHERE id = $1", userId)
	
	return err
}

func (s *Store) CheckIfValidAddress(userId, addressId int) (bool, error) {
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM addresses WHERE id=$1 AND userId=$2)", addressId, userId).Scan(&exists);
	if err != nil {
		return false, err
	}

	return exists, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.LastLogin,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func scanRowsIntoAddress(rows *sql.Rows) (*types.Address, error) {
	address := new(types.Address)

	err := rows.Scan(
		&address.Id,
		&address.UserID,
		&address.Line1,
		&address.Line2,
		&address.City,
		&address.Country,
	)

	if err != nil {
		return nil, err
	}

	return address, nil
}
