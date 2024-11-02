package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/iggy/model"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/kaibling/iggy/pkg/crypto"
)

type userRepo interface {
	SaveUser(u model.NewUser) (*model.User, error)
	FetchUser(id string) (*model.User, error)
	FetchUserByName(name string) (*model.User, error)
	FetchAll() ([]*model.User, error)
	DeleteUser(id string) error
}

type UserService struct {
	ctx  context.Context
	repo userRepo
	cfg  config.Configuration
}

func NewUserService(ctx context.Context, u userRepo, cfg config.Configuration) *UserService {
	return &UserService{ctx: ctx, repo: u, cfg: cfg}
}

func (us *UserService) FetchUser(id string) (*model.User, error) {
	u, err := us.repo.FetchUser(id)
	if err != nil {
		return nil, err
	}
	u.Password = ""
	return u, nil
}

func (us *UserService) CreateUser(u model.NewUser) (*model.User, error) {
	pwd, err := crypto.HashPassword(u.Password, us.cfg.PasswordCost)
	if err != nil {
		return nil, err
	}
	u.Password = pwd
	u.ID = utils.NewULID().String()
	return us.repo.SaveUser(u)
}

func (us *UserService) FetchAll() ([]*model.User, error) {
	return us.repo.FetchAll()
}
func (us *UserService) DeleteUser(id string) error {
	return us.repo.DeleteUser(id)
}

func (us *UserService) EnsureAdmin(password string) (string, error) {
	if _, err := us.repo.FetchUserByName(us.cfg.AdminUser); err != nil {
		if err == sql.ErrNoRows {
			// create Admin user
			if password == "" {
				password = utils.NewULID().String()
			}
			pwdhash, _ := crypto.HashPassword(password, us.cfg.PasswordCost)
			if _, err = us.repo.SaveUser(model.NewUser{
				ID:       utils.NewULID().String(),
				Username: us.cfg.AdminUser,
				Password: pwdhash,
			}); err != nil {
				return "", err
			}
			return password, nil
		}
		return "", err
	}
	return "", nil
}

func (us *UserService) Login(login model.Login, ts *TokenService) (*model.Token, error) {
	user, err := us.repo.FetchUserByName(login.Username)
	if err != nil {
		return nil, err
	}
	ok, err := crypto.CheckPasswordHash(login.Password, user.Password)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("no")
	}
	expirationTime := time.Now().Add(time.Hour * time.Duration(us.cfg.TokenExpiration))
	return ts.CreateToken(model.CreateNewToken(user.ID, expirationTime))
}

func (us *UserService) ValidateToken(token string, ts *TokenService) (*model.User, error) {
	t, err := ts.ReadTokenByValue(token)
	if err != nil {
		fmt.Printf("_> %v\n", err.Error())
		return nil, err
	}

	// TODO validate expiration
	return us.FetchUser(t.User.ID)
}
