package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"kazinfotech/internal/model"
	"kazinfotech/internal/storage"
)

type Service struct {
	store  *storage.Store
	logger *logrus.Logger
}

func NewService(store *storage.Store, logger *logrus.Logger) *Service {
	return &Service{
		store:  store,
		logger: logger,
	}
}

func (s *Service) LoginUser(ctx context.Context, user model.User) (bool, error) {
	return s.store.LoginUser(ctx, user)
}
