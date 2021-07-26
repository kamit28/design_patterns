package creational

import (
	"errors"
	"fmt"
	"hash/fnv"
	"unicode"

	"github.com/asaskevich/govalidator"
)

type User struct {
	name     string
	age      int
	email    string
	password uint32
	status   bool // will be defaulted to false, ans will be changed to true once email is verified
}

func NewUser(name string, age int, email string, password string) (*User, error) {
	if age < 18 {
		return nil, errors.New("user age restrictions apply. Required Age is 18+")
	}
	if !isValidEmail(email) {
		return nil, fmt.Errorf("email was invalid: %s", email)
	}
	if !isAcceptablePassword(password) {
		return nil, errors.New("password does not meet the acceptable criteria")
	}
	u := User{name, age, email, hash(password), false}
	return &u, nil
}

func isValidEmail(email string) bool {
	return govalidator.IsEmail(email)
}

func isAcceptablePassword(p string) bool {
	var number, upper, special, eightOrMore bool
	letters := 0

	for _, c := range p {
		switch {
		case unicode.IsNumber(c):
			number = true
			letters++
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false, false
		}
	}
	eightOrMore = letters >= 8
	return number && upper && special && eightOrMore
}

func hash(p string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(p))
	return h.Sum32()
}
