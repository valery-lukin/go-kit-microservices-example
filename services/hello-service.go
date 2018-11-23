package services

import "errors"

// GreetingService interface
type GreetingService interface {
	Greet(string) (string, error)
	GetInt() int
}

type GreetingServiceImpl struct{}

func (GreetingServiceImpl) Greet(name string) (string, error) {
	if name == "" {
		return "", ErrEmpty
	}
	return "Hello " + name, nil
}

func (GreetingServiceImpl) GetInt() int {
	return 22
}

// ErrEmpty sample error type
var ErrEmpty = errors.New("Empty string")
