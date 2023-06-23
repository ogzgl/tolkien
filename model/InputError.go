package model

import (
	"encoding/json"
	"log"
)

type InputError map[string][]string

func (e InputError) Error() string {
	js, err := json.Marshal(e)
	if err != nil {
		log.Print(err.Error())
		return err.Error()
	}

	return string(js)
}

func NewInputError(inputName, message string) InputError {
	return InputError{
		inputName: {message},
	}
}
