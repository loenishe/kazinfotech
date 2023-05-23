package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"kazinfotech/internal/model"
	"kazinfotech/internal/service"
	"strconv"
)

type Handler struct {
	l   *logrus.Logger
	srv *service.Service
}

func NewHandler(l *logrus.Logger, srv *service.Service) *Handler {
	return &Handler{
		l:   l,
		srv: srv,
	}
}

func (h *Handler) LoginUser(c echo.Context) error {
	var user model.User
	ctx := c.Request().Context()
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(400, MakeResponse(user.Username, 400, err.Error()))
	}
	login, err := h.srv.LoginUser(ctx, user)
	if err != nil {
		return c.JSON(400, MakeResponse(user.Username, 400, err.Error()))

	}
	return c.JSON(200, MakeResponse(user.Username, 200, strconv.FormatBool(login)))
}
