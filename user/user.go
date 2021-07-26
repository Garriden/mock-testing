package user

import "example.com/m/v2/doer"

type User struct {
	Doer doer.Doer
}

func (u *User) Use() error {
	// Hi!
	return u.Doer.DoSomething(123, "Hi!")
}
