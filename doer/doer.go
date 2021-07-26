package doer

//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks github.com/Garriden/mock-testing/doer Doer

type Doer interface {
	DoSomething(int, string) error
}
