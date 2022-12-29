package service

import (
	"basic-k8s/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (s service) GetUser(c *gin.Context) {

	// get payload
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		goutil.ResponseError(c, http.StatusBadRequest, err, nil)
		return
	}

	// get data from database
	var user models.User
	if err := s.gormDb.Where("id = ?", userId).First(&user).Error; err != nil {
		goutil.ResponseError(c, http.StatusInternalServerError, err, nil)
		return
	}

	goutil.ResponseOK(c, http.StatusOK, user)

}
