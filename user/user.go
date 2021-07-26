package user

import "github.com/Garriden/mock-testing/doer"

type User struct {
	Doer doer.Doer
}

func (u *User) Use() error {
	// Hi!
	return u.Doer.DoSomething(123, "Hi!")
}
