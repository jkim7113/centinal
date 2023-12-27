package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jkim7113/centinal/model"
	"github.com/jkim7113/centinal/util"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User)
	Update(ctx context.Context, user model.User)
	UpdatePassword(ctx context.Context, UUID string, password string)
	UpdateRole(ctx context.Context, UUID string, role string)
	VerifyEmail(ctx context.Context, UUID string)
	Delete(ctx context.Context, UUID string)
	FindById(ctx context.Context, UUID string) (model.User, error)
	FindByEmail(ctx context.Context, email string) (model.User, error)
}

type UserRepositoryImpl struct {
	Db *sql.DB
}

func NewUserRepository(Db *sql.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (repo *UserRepositoryImpl) Create(ctx context.Context, user model.User) {
	tx, err := repo.Db.Begin()
	util.PanicIfError(err)
	defer util.CommitOrRollback(tx)

	SQL := "INSERT INTO users (Username, Email, Pw, Bio, PFP) VALUES (?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Pw, user.Bio, user.PFP)
	util.PanicIfError(err)
}

func (repo *UserRepositoryImpl) Update(ctx context.Context, user model.User) {
	tx, err := repo.Db.Begin()
	util.PanicIfError(err)
	defer util.CommitOrRollback(tx)

	SQL := "UPDATE users SET Username = ?, Email = ?, Bio = ?, PFP = ? WHERE UUID = UNHEX(?)"
	_, err = tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Bio, user.PFP, user.UUID)
	util.PanicIfError(err)
}

func (repo *UserRepositoryImpl) UpdatePassword(ctx context.Context, UUID string, password string) {

}

func (repo *UserRepositoryImpl) UpdateRole(ctx context.Context, UUID string, role string) {

}

func (repo *UserRepositoryImpl) VerifyEmail(ctx context.Context, UUID string) {

}

func (repo *UserRepositoryImpl) Delete(ctx context.Context, UUID string) {
	tx, err := repo.Db.Begin()
	util.PanicIfError(err)
	defer util.CommitOrRollback(tx)

	SQL := "DELETE FROM users WHERE UUID = UNHEX(?)"
	_, err = tx.ExecContext(ctx, SQL, UUID)
	util.PanicIfError(err)
}

func (repo *UserRepositoryImpl) FindById(ctx context.Context, UUID string) (model.User, error) {
	tx, err := repo.Db.Begin()
	util.PanicIfError(err)
	defer util.CommitOrRollback(tx)

	SQL := "SELECT HEX(UUID), Username, Email, Pw, Bio, Date, PFP, Verified, Role FROM users WHERE UUID = UNHEX(?)"
	result, errQuery := tx.QueryContext(ctx, SQL, UUID)
	util.PanicIfError(errQuery)
	defer result.Close()

	user := model.User{}

	if result.Next() {
		err := result.Scan(&user.UUID, &user.Username, &user.Email, &user.Pw, &user.Bio, &user.Date, &user.PFP, &user.Verified, &user.Role)
		util.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("Couldn't find a user with such UUID")
	}
}

func (repo *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (model.User, error) {
	tx, err := repo.Db.Begin()
	util.PanicIfError(err)
	defer util.CommitOrRollback(tx)

	SQL := "SELECT HEX(UUID), Username, Email, Pw, Bio, Date, PFP, Verified, Role FROM users WHERE Email = ?"
	result, errQuery := tx.QueryContext(ctx, SQL, email)
	util.PanicIfError(errQuery)
	defer result.Close()

	user := model.User{}

	if result.Next() {
		err := result.Scan(&user.UUID, &user.Username, &user.Email, &user.Pw, &user.Bio, &user.Date, &user.PFP, &user.Verified, &user.Role)
		util.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("Couldn't find a user with such email address")
	}
}
