package validator_test

import (
	"testing"
	vt "validator"
)

type User struct {
	Name    string
	Age     int
	Email   string
	Address string
	Website string
}

func (u *User) Validate() error {
	return vt.ValidateAll(
		vt.String(u.Name, "name").
			Required().
			NotOneOf([]string{"admin", "root"}).
			Max(100),

		vt.Number(u.Age, "age").
			Required("age is required").
			Min(18, "user must be adult").
			Max(100),

		vt.String(u.Email, "email").
			Required("Email is required").
			Email("Email is invalid"),

		vt.String(u.Address, "address").Required(),
		vt.String(u.Website, "website").Url(),
	)
}

func TestValidator(t *testing.T) {
	u := User{
		Name:    "kevin",
		Age:     110,
		Email:   "not an email",
		Website: "d",
	}
	err := u.Validate()
	if err != nil {
		t.Error(err)
	}
}
