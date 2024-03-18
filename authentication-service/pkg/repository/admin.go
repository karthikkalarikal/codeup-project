package repository

import (
	"authentication/pkg/domain"
	"authentication/pkg/repository/interfaces"
	"context"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type adminRepoImpl struct {
	DB *gorm.DB
}

func NewAdminRepo(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepoImpl{
		DB: DB,
	}
}

func (a *adminRepoImpl) Transactions(tx func(repo interfaces.AdminRepository) error) error {
	fmt.Println("here in transactions")
	trx := a.DB.Begin()
	repo := NewAdminRepo(trx)
	err := tx(repo)
	if err != nil {
		return err
	}
	if err := trx.Commit().Error; err != nil {
		return err
	}
	fmt.Println("err ", err)
	return nil
}

func (a *adminRepoImpl) BlockUser(ctx context.Context, id int) (domain.User, error) {
	fmt.Println("here block", id)

	user, err := a.GetUserById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	fmt.Println("user ", user)

	// var block bool
	if user.Email == "" {
		return domain.User{}, errors.New("no entry of that name")

	}
	query := `UPDATE users
			SET blocked = $1
			WHERE id = $2
	`
	fmt.Println("Generated Query:", query)
	fmt.Println("user block", user.Blocked)
	err = a.DB.WithContext(ctx).Exec(query, !user.Blocked, id).Error
	fmt.Println("err ", err)
	if err != nil {
		return domain.User{}, err
	}
	log.Println("here after blocking")
	user, err = a.GetUserById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}

// add transactions to improve all these functions - todo
func (a *adminRepoImpl) GetUserById(ctx context.Context, id int) (user domain.User, err error) {
	query := `select * from users where id = ?`
	// user := new(domain.User)
	err = a.DB.WithContext(ctx).Raw(query, id).Scan(&user).Error
	if err != nil {
		return
	}
	return user, nil
}
