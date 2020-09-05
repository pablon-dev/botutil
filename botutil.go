package botutil

import (
	"errors"
	"fmt"
	"strings"
)

const slash byte = 47

//Getcommand Recieve command and return parameters
func Getcommand(text string, separ byte) (command []string, err error) {
	text = strings.Trim(text, " ")
	if text[0] != separ {
		return nil, fmt.Errorf("Commands must start with: '%c' ", separ)
	}
	s := strings.Split(text, " ")
	return s, nil
}

//GetcommandB Recieve command and arguments and return parameters
func GetcommandB(comm string, args string) (command []string, err error) {
	if comm == "" {
		return nil, errors.New("Not valid Command ")
	}
	command = append(command, "/"+comm)
	if args != "" {
		s := strings.Split(args, " ")
		command = append(command, s...)
	}
	return command, nil
}

//SplitIgnore Recieve a string and return an array of substrings without empty space
func SplitIgnore(text string) (parts []string) {
	splitF := func(c rune) bool {
		return c == ' '
	}
	parts = strings.FieldsFunc(text, splitF)
	return
}

//URLform give the html/markdown form of a given url and body
func URLform(url, body string) (result string) {
	result = fmt.Sprintf(`[%s](%s)`, body, url)
	return
}

//NewLine append y to x in new Line
func NewLine(x, y string) string {
	if x == "" {
		return y
	}
	return x + "\n" + y

}
