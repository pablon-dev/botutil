package botutil

import (
	"errors"
	"strings"
)

//Getcommand Recibe un comando proporcionado por el usuario, devuelve los distintos parámetros
func Getcommand(text string) (command []string, err error) {
	text = strings.Trim(text, " ")
	if text[0] != 47 {
		return nil, errors.New("Commands must start with: '/' ")
	}
	s := strings.Split(text, " ")
	return s, nil
}
