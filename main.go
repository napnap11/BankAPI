package main

import (
	"bankapi/server"
	"bankapi/user"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        first_name TEXT,
        last_name TEXT
);

CREATE TABLE IF NOT EXISTS bankaccounts (
        id SERIAL PRIMARY KEY,
        user_id INT FOREIGN KEY REFERRENCES users(user_id),
        account_number INT UNIQUE,
        name TEXT,
        balance INT
);

CREATE TABLE IF NOT EXISTS keys (
        id SERIAL PRIMARY KEY,
        key TEXT
);
	`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}

	s := &server.Server{
		db: db,
		userService: &user.UserServiceImp{
			db: db,
		},
		bankService: &bankaccount.BankServiceImpl{
			db: db,
		},
	}

	r := server.SetupRoute(s)

	r.Run(":" + os.Getenv("PORT"))
}
