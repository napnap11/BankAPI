package server

import (
	"bankapi/bankaccount"
	"bankapi/user"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockUserService struct {
}

func (s *mockUserService) All() (users []user.User, err error) {
	users = []user.User{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
		},
	}
	return
}

func (s *mockUserService) FindByID(id int) (u *user.User, err error) {
	u = &user.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}
	return
}

func (s *mockUserService) Create(u user.User) (err error) {
	return
}

func (s *mockUserService) Update(id int, u user.User) (us *user.User, err error) {
	us = &user.User{
		ID:        id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
	return
}

func (s *mockUserService) Delete(id int) (err error) {
	return
}

type mockBankService struct{}

func (s *mockBankService) Create(id int, account bankaccount.BankAccount) (err error) {
	return
}

func TestUserAll(t *testing.T) {
	s := &Server{
		DB:          nil,
		UserService: &mockUserService{},
	}
	r := SetupRoute(s)
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.NoError(t, err)
	expectedJSON := `[{"id":1,"first_name":"John","last_name":"Doe"}]`
	assert.Equal(t, expectedJSON, string(body))
}

func TestUserByID(t *testing.T) {
	s := &Server{
		DB:          nil,
		UserService: &mockUserService{},
	}
	r := SetupRoute(s)
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/users/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.NoError(t, err)
	expectedJSON := `{"id":1,"first_name":"John","last_name":"Doe"}`
	assert.Equal(t, expectedJSON, string(body))
}

func TestCreateUser(t *testing.T) {
	s := &Server{
		DB:          nil,
		UserService: &mockUserService{},
	}
	r := SetupRoute(s)
	user := `{"first_name":"John","last_name":"Doe"}`
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/users", strings.NewReader(user))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resp := w.Result()
	_, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	s := &Server{
		DB:          nil,
		UserService: &mockUserService{},
	}
	r := SetupRoute(s)
	user := `{"first_name":"John","last_name":"Doe"}`
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:8080/users/1", strings.NewReader(user))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.NoError(t, err)
	expectedJSON := `{"id":1,"first_name":"John","last_name":"Doe"}`
	assert.Equal(t, expectedJSON, string(body))
}

func TestDeleteUser(t *testing.T) {
	s := &Server{
		DB:          nil,
		UserService: &mockUserService{},
	}
	r := SetupRoute(s)
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:8080/users/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resp := w.Result()
	_, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.NoError(t, err)
}

func TestCreateBankAccount(t *testing.T) {
	s := &Server{
		DB:          nil,
		UserService: &mockUserService{},
	}
	r := SetupRoute(s)
	account := `{"account_number":123456,"name":"John Doe"}`
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/users/1/bankAccounts", strings.NewReader(account))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resp := w.Result()
	_, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.NoError(t, err)
}
