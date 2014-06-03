package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
)

type User struct {
	Username  string `json:"userName"			bson:"username"`
	Password  string `json:"password"			bson:"password"`
	FirstName string `json:"firstName"			bson:"firstname"`
	LastName  string `json:"lastName"			bson:"lastname"`
	IsAdmin   bool   `json:"isAdmin"			bson:"isadmin"`
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *User) Validate(v *revel.Validation) {
	v.Check(user.Username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	ValidatePassword(v, user.Password).
		Key("user.Password")

	v.Check(user.Username,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}
