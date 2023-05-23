package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"kazinfotech/internal/handler"
	"kazinfotech/internal/service"
	"kazinfotech/internal/storage"
	"log"
)

func main() {

	db, err := storage.DBCon()
	if err != nil {
		log.Fatal(err)
	}
	logger := logrus.New()

	store := storage.NewStore(db, logger)
	serv := service.NewService(store, logger)
	hand := handler.NewHandler(logger, serv)

	e := echo.New()
	e.POST("/user/login", hand.LoginUser)
	e.Logger.Fatal(e.Start(":8080"))
}
