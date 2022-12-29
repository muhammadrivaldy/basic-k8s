package service

import (
	"basic-k8s/config"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	goutil "github.com/muhammadrivaldy/go-util"
	"gorm.io/gorm"
)

type service struct {
	gormDb *gorm.DB
}

func NewService() {

	// get config filename from environment
	filename := os.Getenv("CONFIG_FILENAME")
	if filename == "" {
		filename = "local.conf"
	}

	// open file config
	osFile, err := goutil.OpenFile("./config", filename)
	if err != nil {
		panic(err)
	}

	// load config
	var config config.Config
	if err := goutil.Configuration(osFile, &config); err != nil {
		panic(err)
	}

	// open connection to mysql db
	sqlDb, err := goutil.NewMySQL(config.Database.User, config.Database.Password, config.Database.Address, config.Database.Schema, nil)
	if err != nil {
		panic(err)
	}

	// migration database
	driver, err := mysql.WithInstance(sqlDb, &mysql.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		config.Database.Schema,
		driver)
	if err != nil {
		panic(err)
	}

	// do a migration up
	m.Up()

	// wrap the connection db within gorm
	gormDb, err := goutil.NewGorm(sqlDb, "mysql", 2)
	if err != nil {
		panic(err)
	}

	svc := service{gormDb: gormDb}

	// routing endpoints
	route := gin.New()
	route.Handle(http.MethodPost, "/api/v1/users", svc.CreateUser)
	route.Handle(http.MethodGet, "/api/v1/users/:id", svc.GetUser)
	route.Run(fmt.Sprintf(":%d", config.Port))

}
