package domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

var (
	ErrInvalidEMail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
	ErrHashError       = errors.New("failed to hash password")
	emailRegex         = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

//	passwordRegex      = regexp.MustCompile("^(((?=.*[a-z])(?=.*[A-Z]))|((?=.*[a-z])(?=.*[0-9]))|((?=.*[A-Z])(?=.*[0-9])))(?=.{6,})")
)

//easyjson:json
type User struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id"`
	EMail string             `json:"email" bson:"email"`
	Hash  []byte             `json:"hash" bson:"hash"`
	Admin bool               `json:"admin" bson:"admin"`
}

func isEMailValid(email string) bool {
	return emailRegex.MatchString(email)
}

func isPasswordValid(password string) bool {
	return true
	//return passwordRegex.MatchString(password)
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func NewAdmin(email, password string) (*User, error) {
	if !isEMailValid(email) {
		return nil, ErrInvalidEMail
	}

	if !isPasswordValid(password) {
		return nil, ErrInvalidPassword
	}

	hash, err := hashPassword(password)
	if err != nil {
		return nil, ErrHashError
	}

	return &User{
		ID:    primitive.NewObjectID(),
		EMail: email,
		Hash:  hash,
		Admin: true,
	}, nil
}

func NewUser(email, password string) (*User, error) {
	if !isEMailValid(email) {
		return nil, ErrInvalidEMail
	}

	if !isPasswordValid(password) {
		return nil, ErrInvalidPassword
	}

	hash, err := hashPassword(password)
	if err != nil {
		return nil, ErrHashError
	}

	return &User{
		ID:    primitive.NewObjectID(),
		EMail: email,
		Hash:  hash,
	}, nil
}
