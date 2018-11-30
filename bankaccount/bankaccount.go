package bankaccount

import (
	"database/sql"
)

type BankAccount struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	AccountNo int    `json:"account_number"`
	Name      string `json:"name"`
	Balance   int    `json:"balance"`
}

type BankService interface {
	Create(int, BankAccount) error
}

type BankServiceImpl struct {
	DB *sql.DB
}

func (s *BankServiceImpl) Create(userID int, a BankAccount) (err error) {
	row := s.DB.QueryRow("INSERT INTO bankaccounts (user_id, account_number,name,balance) values ($1, $2,$3,$4) RETURNING id", userID, a.AccountNo, a.Name, 0)
	err = row.Scan(&a.ID)
	return
}
