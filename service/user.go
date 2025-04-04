package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/apiforge/params"
	"github.com/kaibling/iggy/apperror"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/kaibling/iggy/pkg/crypto"
)

type userRepo interface {
	SaveUser(u entity.NewUser) (*entity.User, error)
	FetchUserByName(name string) (*entity.User, error)
	FetchByIDs(ids []string) ([]*entity.User, error)
	// FetchAll() ([]*entity.User, error)
	DeleteUser(id string) error
	IDQuery(query string) ([]string, error)
}

type UserService struct {
	ctx  context.Context
	repo userRepo
	cfg  config.Configuration
}

func NewUserService(ctx context.Context, u userRepo, cfg config.Configuration) *UserService {
	return &UserService{ctx: ctx, repo: u, cfg: cfg}
}

func (us *UserService) FetchUser(id string) (*entity.User, error) {
	loadedUser, err := us.FetchUsers([]string{id})
	if err != nil {
		return nil, err
	}

	return loadedUser[0], nil
}

func (us *UserService) FetchUsers(ids []string) ([]*entity.User, error) {
	loadedUsers, err := us.repo.FetchByIDs(ids)
	if err != nil {
		return nil, err
	}

	for _, u := range loadedUsers {
		u.Redact()
	}

	return loadedUsers, nil
}

func (us *UserService) CreateUser(u entity.NewUser) (*entity.User, error) {
	pwd, err := crypto.HashPassword(u.Password, us.cfg.App.PasswordCost)
	if err != nil {
		return nil, err
	}

	u.Password = pwd

	if u.ID == "" {
		u.ID = utils.NewULID().String()
	}

	return us.repo.SaveUser(u)
}

func (us *UserService) FetchByPagination(qp params.Pagination) ([]*entity.User, params.Pagination, error) {
	p := NewPagination(qp, "users")

	idQuery := p.GetCursorSQL()

	ids, err := us.repo.IDQuery(idQuery)
	if err != nil {
		return nil, params.Pagination{}, err
	}

	ids, pag := p.FinishPagination(ids)

	loadedUsers, err := us.repo.FetchByIDs(ids)
	if err != nil {
		return nil, params.Pagination{}, err
	}

	return loadedUsers, pag, nil
}

func (us *UserService) DeleteUser(id string) error {
	return us.repo.DeleteUser(id)
}

func (us *UserService) EnsureAdmin(password string) (string, error) {
	if _, err := us.repo.FetchUserByName(us.cfg.App.AdminUser); err != nil { //nolint: nestif
		if errors.Is(err, sql.ErrNoRows) {
			// create Admin user
			if password == "" {
				password = utils.NewULID().String()
			}

			pwdhash, _ := crypto.HashPassword(password, us.cfg.App.PasswordCost)

			if _, err = us.repo.SaveUser(entity.NewUser{
				ID:       utils.NewULID().String(),
				Username: us.cfg.App.AdminUser,
				Password: pwdhash,
				Active:   true,
			}); err != nil {
				return "", err
			}

			return password, nil
		}

		return "", err
	}

	return "", nil
}

func (us *UserService) Login(login entity.Login, ts *TokenService) (*entity.Token, error) {
	user, err := us.repo.FetchUserByName(login.Username)
	if err != nil {
		return nil, err
	}

	ok, err := crypto.CheckPasswordHash(login.Password, user.Password)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, apperror.ErrForbidden
	}

	expirationTime := time.Now().Add(time.Hour * time.Duration(us.cfg.App.TokenExpiration))

	return ts.CreateToken(entity.CreateNewToken(user.ID, expirationTime))
}

func (us *UserService) ValidateToken(token string, ts *TokenService) (*entity.User, error) {
	t, err := ts.ReadTokenByValue(token)
	if err != nil {
		return nil, err
	}

	// TODO validate expiration
	return us.FetchUser(t.User.ID)
}
