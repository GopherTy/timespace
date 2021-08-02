package common

// IRegister common modules should implement this interface
type IRegister interface {
	Name() string
	CheckIn() error
}
