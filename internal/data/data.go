package data

import (
	"context"
	"userauth/internal/conf"
	"userauth/internal/biz"
        "github.com/google/wire"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

type User struct {
	ID       uint   `gorm:"primarykey"`
	Username string `gorm:"uniqueIndex"`
	Password string
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (r *userRepo) AddUser(ctx context.Context, username, password string) error {
	user := &User{Username: username, Password: password}
	return r.data.db.WithContext(ctx).Create(user).Error
}

func (r *userRepo) AuthenticateUser(ctx context.Context, username, password string) (bool, error) {
	var user User
	if err := r.data.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return false, err
	}
	return user.Password == password, nil
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type Data struct {
	db *gorm.DB
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := gorm.Open(postgres.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return &Data{db: db}, func() {}, nil
}

