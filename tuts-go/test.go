package main

import (
	"strings"
)

func getInitials(name string) (string, string) {
	splitNames := strings.Split(strings.ToUpper(name), " ")
	var initials []string
	for _, v := range splitNames {
		initials = append(initials, v[:1])
	}
	if len(initials) == 1 {
		return initials[0], "_"
	}
	return initials[0], initials[1]
}

var mapTest = map[string]uint8{
	"test1": 15,
	"test2": 50,
}
