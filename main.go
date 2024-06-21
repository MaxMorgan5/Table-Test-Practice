package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func main() {
	s := "aaaaaa"
	sep := "h"

	result, err := SepStrings(s, sep)
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range result {
		fmt.Println(val)
	}
}

func SepStrings(s, sep string) ([]string, error) {

	if s == "" {
		return nil, errors.New("input string cannot be empty")
	}

	if sep == "" {
		return nil, errors.New("must provide a value for sep")
	}
	if !strings.Contains(s, sep) {
		return nil, errors.New("string does not contain sep")
	}

	var result []string

	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}

	if len(s) > 0 || len(result) == 0 {
		result = append(result, s)
	}

	return result, nil
}
