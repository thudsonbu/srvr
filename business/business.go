package business

import "errors"

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (l LoggerAdapter) Log(message string) {
	l(message)
}

type BusinessLogic struct {
	l  Logger
	ds DataStore
}

func (bl BusinessLogic) DoSomething() {
	bl.l.Log("Doing something")
}

func (bl BusinessLogic) SayHello(userID string) (string, error) {
	name, ok := bl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("No user found")
	}
	return "Hello, " + name, nil
}

func (bl BusinessLogic) SayGoodbye(userID string) (string, error) {
	name, ok := bl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("No user found")
	}
	return "Goodbye, " + name, nil
}
