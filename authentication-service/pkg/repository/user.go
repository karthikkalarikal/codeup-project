package repository

import (
	"authentication/pkg/domain"
	"authentication/pkg/repository/interfaces"
	customErrors "authentication/pkg/utils/errors"
	"authentication/pkg/utils/request"
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return domain.User{}, err
	}

	tx := u.DB.WithContext(ctxDeadline).Begin()
	userToCreate := domain.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Password:  string(hashedPassword),
		Email:     user.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	fmt.Println("user", tx)

	// to check for duplicate values
	fmt.Println("email1", user.Email)
	_, err = u.FindUserByEmail(ctxDeadline, user.Email)
	fmt.Println("err ", err, "gorm error:", gorm.ErrRecordNotFound)
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
	fmt.Println("email2", user.Email)
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
	fmt.Println("email3", email)
	var user domain.User

	tx := u.DB.WithContext(ctxDeadline).Begin()

	result := tx.
		Where("email = ?", email).
		First(&user)

	if result.Error != nil {
		fmt.Println("err ", result.Error)
		return domain.User{}, result.Error
	}
	fmt.Println("user", user)
	return user, nil
}

// get all user repo
func (u *userDatabase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	users := new([]domain.User)

	err := u.DB.WithContext(ctx).Raw("select * from users").Scan(users).Error
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("database query timed out")
		}
		return nil, err
	}
	return *users, nil
}

// get user by email
func (u *userDatabase) SearchUserByEmail(ctx context.Context, email string) ([]domain.User, error) {
	user := new([]domain.User)

	query := "select * from users where email like ?"

	err := u.DB.WithContext(ctx).Raw(query, "%"+email+"%").Scan(user).Error
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("database query timed out: %w", err)
		}
		return nil, err
	}

	return *user, nil
}

// get user by username
func (u *userDatabase) SearchUserByUsername(ctx context.Context, username string) ([]domain.User, error) {
	users := new([]domain.User)

	query := "select * from users where username like ?"
	fmt.Println(query, username)
	err := u.DB.WithContext(ctx).Raw(query, "%"+username+"%").Scan(users).Error
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("database query timed out: %w", err)
		}
		return nil, err
	}

	return *users, nil
}
