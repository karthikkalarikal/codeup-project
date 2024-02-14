package data

import (
	"authentication/pkg/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

const dbTimeout = time.Second * 3
const (
	nullError     = "no empty values"
	errorPassword = "the passwords don't match"
)

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User: User{},
	}
}

type Models struct {
	User User
}

type User struct {
	Username        sql.NullString `json:"username"`
	Email           sql.NullString `json:"email"`
	Password        string         `json:"password"`
	ConfirmPassword string         `json:"confirmpassword"`
	Name            sql.NullString `json:"name"` // for proper representaion of null value in go sql.Null is used
}

type TempUser struct {
	Id              int            `json:"id"`
	FirstName       string         `json:"first_name"`
	LastName        string         `json:"last_name"`
	Username        sql.NullString `json:"username"`
	Email           string         `json:"email"`
	Password        string         `json:"password"`
	ConfirmPassword string         `json:"confirmpassword"`
	Name            sql.NullString `json:"name"`
}

func (u *User) GetByEmail(email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, password, created_at, updated_at from users where email = $1`

	var user domain.User

	row := db.QueryRowContext(ctx, query, email)
	fmt.Println("row  ", row)
	fmt.Println("email: ", email)
	err := row.Scan(
		&user.ID,
		&user.Email,
		// &user.Username,
		&user.Password,
		// &user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	fmt.Println("error in row. Scan", err)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) UserSignUp(user User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	if user.Password != user.ConfirmPassword {
		return 0, errors.New(errorPassword)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return 0, err
	}

	var newId int

	if user.Name.String == "" || user.Email.String == "" || user.Password == "" {
		return 0, errors.New(nullError)
	}

	//checking validity

	err = validator.New().Struct(user)
	if err != nil {

		return 0, err

	}
	stmt := `insert into users
	 (username,email,password,name,created_at,updated_at)
	 values ($1,$2,$3,$4,$5,$6,$7)`
	err = db.QueryRowContext(ctx, stmt,
		user.Username,
		user.Email,

		hashedPassword,
		user.Name,
		time.Now(),
		time.Now(),
	).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil

}
