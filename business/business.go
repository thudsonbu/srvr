package business

import "errors"

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// Define an interface for the logger that our BusinessLogic will use. This
// allows us to easily swap out the logger later.
type Logger interface {
	Log(message string)
}

// Create an adapter from the Logger to our Business logic
type LoggerAdapter func(message string)

func (logger LoggerAdapter) Log(message string) {
	logger(message)
}

type BusinessLogic struct {
	logger Logger
	store  DataStore
}

func (businessLogic BusinessLogic) DoSomething() {
	businessLogic.logger.Log("Doing something")
}

func (businessLogic BusinessLogic) SayHello(userID string) (string, error) {
	name, ok := businessLogic.store.UserNameForID(userID)
	if !ok {
		return "", errors.New("No user found")
	}
	return "Hello, " + name, nil
}

func (businessLogic BusinessLogic) SayGoodbye(userID string) (string, error) {
	name, ok := businessLogic.store.UserNameForID(userID)
	if !ok {
		return "", errors.New("No user found")
	}
	return "Goodbye, " + name, nil
}

func NewBusinessLogic(logger Logger, store DataStore) BusinessLogic {
	return BusinessLogic{
		logger: logger,
		store:  store,
	}
}
