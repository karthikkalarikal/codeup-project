package repository

import (
	"authentication/pkg/domain"
	"authentication/pkg/repository/interfaces"
	customErrors "authentication/pkg/utils/errors"
	"authentication/pkg/utils/request"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{
		DB: DB,
	}
}

//  create a transactions function to further develop the database operations

// ------------------- user signup ----------------- \\
func (u *userDatabase) UserSignUp(ctx context.Context, user request.UserSignUpRequest) (userDetails domain.User, err error) {

	ctxDeadline, cancel := context.
		WithTimeout(ctx, 5*time.Second)

	defer cancel()

	tx := u.DB.WithContext(ctxDeadline).Begin()
	userToCreate := domain.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// to check for duplicate values
	_, err = u.FindUserByEmail(ctxDeadline, user.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.User{}, errors.New(customErrors.UserAlreadyExists)
	}

	result := tx.Create(&userToCreate)
	if result.Error != nil {
		tx.Rollback()

		if errors.
			Is(result.Error, context.DeadlineExceeded) {
			return domain.User{}, errors.New(customErrors.DatabaseTimeOut)
		}

		// error checking for duplicate values to be added
		return domain.User{}, err
	}
	// useing transaction for data safty
	tx.Commit()

	userResponse, err := u.FindUserByEmail(ctxDeadline, user.Email)
	if err != nil {
		return domain.User{}, err
	}

	return userResponse, nil

}

// -----------------find user by email -------------------\\

func (u *userDatabase) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	ctxDeadline, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user domain.User

	tx := u.DB.WithContext(ctxDeadline).Begin()

	result := tx.
		Where("email = ?", user.Email).
		First(&user)

	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return user, nil
}
