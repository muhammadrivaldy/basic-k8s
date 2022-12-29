package service

import (
	"basic-k8s/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	goutil "github.com/muhammadrivaldy/go-util"
)

type PayloadCreateUser struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (s service) CreateUser(c *gin.Context) {

	// get payload
	var payload PayloadCreateUser
	if err := c.Bind(&payload); err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, err, nil)
		return
	}

	// insert data
	if err := s.gormDb.Create(&models.User{Name: payload.Name, Address: payload.Address}).Error; err != nil {
		goutil.ResponseError(c, http.StatusInternalServerError, err, nil)
		return
	}

	goutil.ResponseOK(c, http.StatusCreated, nil)

}
