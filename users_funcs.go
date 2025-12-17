package main

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// getAllUsers returns slise of all users
func getAllUsers(db *sql.DB) ([]User, error) {
	var usrs []User
	rows, err := db.Query("SELECT id, username, role FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Username, &u.Role)
		if err != nil {
			return nil, err
		}
		usrs = append(usrs, u)
	}
	return usrs, nil
}

// createUser - add new user in DB
func createUser(db *sql.DB, username, password, role string) error {
	if username == "" || password == "" {
		return errors.New("username or password can't be empty")
	}
	//default value
	if role == "" {
		role = "user"
	}
	if role != "user" && role != "admin" {
		return errors.New("role can be only user or admin")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3)`,
		username,
		string(hash),
		role,
	)
	return err
}

// deleteUser - delete out DB
func deleteUser(db *sql.DB, id int) error {
	res, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	check, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if check == 0 {
		return errors.New("ошибка удаления: такого пользователя нет в базе, id") // remake to Logs
	}

	return nil
}

// updateUsername - update user name
func updateUsername(db *sql.DB, id int, newUsername string) error {
	if newUsername == "" {
		return errors.New("new name can't be empty")
	}
	_, err := db.Exec("UPDATE users SET username = $1 WHERE id = $2", newUsername, id)
	return err
}

// updatePassword - set new password
func updatePassword(db *sql.DB, id int, newPassword string) error {
	if newPassword == "" {
		return errors.New("new passworc can't be empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE users SET password_hash = $1 WHERE id = $2", string(hash), id)
	return err
}
