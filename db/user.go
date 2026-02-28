package db

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          int64     `json:"id" db:"id"`
	Email       string    `json:"email" db:"email"`
	DisplayName string    `json:"displayName" db:"display_name"`
	Password    string    `json:"-" db:"password"`
	CanLogin    bool      `json:"canLogin" db:"can_login"`
	IsAdmin     bool      `json:"isAdmin" db:"is_admin"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}

func GetUserByEmail(email string) (User, error) {
	var u User
	err := DB.QueryRow(
		`SELECT id, email, COALESCE(display_name, ''), password, can_login, is_admin FROM users WHERE email = $1 AND deleted_at IS NULL`,
		email,
	).Scan(&u.Id, &u.Email, &u.DisplayName, &u.Password, &u.CanLogin, &u.IsAdmin)
	return u, err
}

func GetUserByID(id int64) (User, error) {
	var u User
	err := DB.QueryRow(
		`SELECT id, email, COALESCE(display_name, ''), password, can_login, is_admin FROM users WHERE id = $1 AND deleted_at IS NULL`,
		id,
	).Scan(&u.Id, &u.Email, &u.DisplayName, &u.Password, &u.CanLogin, &u.IsAdmin)
	return u, err
}

func CreateUser(email, displayName, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = DB.Exec(
		`INSERT INTO users (email, display_name, password, created_at, updated_at) VALUES ($1, NULLIF($2, ''), $3, now(), now())`,
		email, displayName, string(hashed),
	)
	return err
}

func UpdateUser(id int64, email, displayName string) error {
	_, err := DB.Exec(
		`UPDATE users SET email = $2, display_name = NULLIF($3, ''), updated_at = now() WHERE id = $1`,
		id, email, displayName,
	)
	return err
}

// UpsertRootUser creates the admin account from env vars if it doesn't already exist.
func UpsertRootUser() error {
	email := getEnv("ROOT_USER_EMAIL", "admin@example.com")

	var count int
	if err := DB.QueryRow(`SELECT COUNT(*) FROM users WHERE email = $1`, email).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	password := getEnv("ROOT_USER_PASSWORD", "password")
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`INSERT INTO users 
		(email, password, can_login, is_admin, created_at, updated_at) 
		VALUES 
		($1, $2, true, true, now(), now())`,
		email, string(hashed),
	)
	return err
}
