package botutil

import (
	"errors"
	"strings"
)

const slash int = 47

//Getcommand Recieve command and return parameters
func Getcommand(text string) (command []string, err error) {
	text = strings.Trim(text, " ")
	if text[0] != slash {
		return nil, errors.New("Commands must start with: '/' ")
	}
	s := strings.Split(text, " ")
	return s, nil
}
