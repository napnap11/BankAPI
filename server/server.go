package server

import (
	"bankapi/bankaccount"
	"bankapi/user"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	DB          *sql.DB
	UserService user.UserService
	BankService bankaccount.BankService
}

func addSecretKey() {

}

func checkKey(key string) bool {
	return true
}

func (s *Server) getAllUser(c *gin.Context) {
	users, err := s.UserService.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (s *Server) getUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := s.UserService.FindByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (s *Server) createUser(c *gin.Context) {
	var user user.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}
	err = s.UserService.Create(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}
	c.Status(http.StatusOK)
}

func (s *Server) updateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user user.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}
	u, err := s.UserService.Update(id, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, *u)
}

func (s *Server) deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := s.UserService.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}
	c.Status(http.StatusOK)
}
func (s *Server) CreateBankAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var account bankaccount.BankAccount
	err := c.ShouldBindJSON(&account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}
	err = s.BankService.Create(id, account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}
	c.Status(http.StatusOK)
}

func SetupRoute(s *Server) *gin.Engine {
	r := gin.Default()
	addSecretKey()
	r.Use(func(c *gin.Context) {
		key := c.Request.Header.Get("X-Auth-Token")
		if checkKey(key) {
			return
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	})
	r.GET("/users", s.getAllUser)
	r.GET("/users/:id", s.getUserByID)
	r.POST("/users", s.createUser)
	r.PUT("/users/:id", s.updateUser)
	r.DELETE("/users/:id", s.deleteUser)
	r.POST("/users/:id/bankAccounts", s.CreateBankAccount)
	return r
}
