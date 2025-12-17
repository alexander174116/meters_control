package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5555"
	user     = "namorbor"
	password = "gapega76"
	dbname   = "app_db"
	sslmode  = "disable"
)

type User struct {
	Id       int    `sql:"id"`
	Username string `sql:"username"`
	Password string `sql:"password_hash"`
	Role     string `sql:"role"`
}

type House struct {
	ID      int           `sql:"id"`
	Address string        `sql:"address"`
	UserUD  sql.NullInt64 `sql:"user_id"`
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(1, err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(2, err)
	}
	fmt.Printf("connected to %s...\n", dbname)
	// TEST FUNCS
	// err = createUser(db, "тестРу", "test_pass", "admin")
	// err = deleteUser(db, 18)
	// err = updateUsername(db, 18, "lol")
	// u, err := checkLogin(db, "тестРу", "test_pass")
	// _, err = checkLogin(db, "rinat_kosmonavt", "test_pass")
	testId := 29
	err = addHouse(db, "Космонав0тов 651г", &testId)
	// err = deleteHouse(db, 9)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Println()
	users, err := getAllUsers(db)
	if err != nil {
		log.Fatal(3, err)
	}
	for idx, val := range users {
		fmt.Printf("%d - %v\n", idx, val)
	}
	fmt.Println()
	houses, err := getAllHouses(db)
	if err != nil {
		log.Fatal(3, err)
	}
	for idx, val := range houses {
		fmt.Printf("%d - %v\n", idx, val)
	}
}
