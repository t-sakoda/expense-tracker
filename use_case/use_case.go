package use_case

type UseCase interface {
	Execute(args ...interface{}) error
}
