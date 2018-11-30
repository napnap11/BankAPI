package user

import "database/sql"

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserService interface {
	All() ([]User, error)
	FindByID(int) (*User, error)
	Create(User) error
	Update(int, User) (*User, error)
	Delete(int) error
}

type UserServiceImpl struct {
	DB *sql.DB
}

func (s *UserServiceImpl) All() (users []User, err error) {
	rows, err := s.DB.Query("SELECT id, first_name, last_name FROM users")
	if err != nil {
		return
	}
	users = []User{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

func (s *UserServiceImpl) FindByID(id int) (user *User, err error) {
	rows, err := s.DB.Query("SELECT id, first_name, last_name FROM users")
	if err != nil {
		return
	}
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName)
	}
	return
}

func (s *UserServiceImpl) Create(u User) (err error) {
	row := s.DB.QueryRow("INSERT INTO users (first_name, last_name) values ($1, $2) RETURNING id", u.FirstName, u.LastName)
	err = row.Scan(&u.ID)
	return
}

func (s *UserServiceImpl) Update(id int, u User) (user *User, err error) {
	stmt := "UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3"
	_, err = s.DB.Exec(stmt, u.FirstName, u.LastName, id)
	if err != nil {
		return
	}
	return s.FindByID(id)
}

func (s *UserServiceImpl) Delete(id int) (err error) {
	_, err = s.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return
}
