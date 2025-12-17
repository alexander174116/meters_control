package main

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// checkLogin check log+pass and return user info
func checkLogin(db *sql.DB, username, password string) (*User, error) {
	if username == "" || password == "" {
		return nil, errors.New("username or password can't be empty")
	}

	var u User
	err := db.QueryRow(
		"SELECT id, username, password_hash, role FROM users WHERE username = $1",
		username,
	).Scan(&u.Id, &u.Username, &u.Password, &u.Role)
	if err == sql.ErrNoRows {
		return nil, errors.New("user nor found")
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return &u, nil
}
