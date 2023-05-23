package storage

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"kazinfotech/internal/model"
)

type Store struct {
	db     *sqlx.DB
	logger *logrus.Logger
}

func NewStore(db *sqlx.DB, logger *logrus.Logger) *Store {
	return &Store{
		db:     db,
		logger: logger,
	}
}

func (s *Store) LoginUser(ctx context.Context, user model.User) (bool, error) {
	var data []model.User
	err := s.db.SelectContext(ctx, &data, "Select * from users where user_id = $1", user.Username)
	if err != nil {
		return false, err
	}
	h := md5.New()
	h.Write([]byte(user.Password))
	pas := hex.EncodeToString(h.Sum(nil))
	if len(data) > 0 {
		if data[0].Password == pas {
			return true, nil
		} else {
			return false, errors.New("Неправильный пароль")
		}
	} else {
		return false, errors.New("Пользователь не найден")
	}

}
