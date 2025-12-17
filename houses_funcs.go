package main

import (
	"database/sql"
	"errors"
)

func addHouse(db *sql.DB, address string, userID *int) error {
	var uid sql.NullInt64
	if userID != nil {
		uid = sql.NullInt64{
			Int64: int64(*userID),
			Valid: true,
		}
	} else {
		uid = sql.NullInt64{Valid: false}
	}

	res, err := db.Exec(
		"INSERT INTO houses (address, user_id) VALUES ($1, $2)",
		address, uid,
	)
	if res == nil {
		return errors.New("такой адрес уже есть в базе")
	}
	return err
}

func getAllHouses(db *sql.DB) ([]House, error) {
	var houses []House
	rows, err := db.Query("SELECT id, address, user_id FROM houses")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var h House
		err := rows.Scan(&h.ID, &h.Address, &h.UserUD)
		if err != nil {
			return nil, err
		}
		houses = append(houses, h)
	}
	return houses, nil
}

func deleteHouse(db *sql.DB, id int) error {
	res, err := db.Exec("DELETE FROM houses WHERE id = $1", id)
	if err != nil {
		return err
	}
	check, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if check == 0 {
		return errors.New("ошибка удаления: такого пользователя нет в базе, id")
	}
	return nil
}
