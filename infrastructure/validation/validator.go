package validation

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)
//TODO Make it with a builder pattern
type Validator interface {
	ValidateExpression(exp string) (valid bool, error error)
}

type expressionValidator struct {
	exp string
}

func New() *expressionValidator {
	return &expressionValidator{exp: ""}
}

func (i expressionValidator) ValidateExpression(exp string) (valid bool, error error) {
	error = errors.New("")
	valid = true

	//Validate emptiness
	if exp == "" {
		valid = false
		exp = "EMPTY_INFIX_EXPRESSION"
		error = fmt.Errorf("the expression %s is invalid", exp)
		return
	}

	//Validate letters
	if IsLetter(exp) {
		valid = false
		error = fmt.Errorf("the expression %s is invalid", exp)
		return
	}

	//Validate parentheses
	if ContainsParentheses(exp) {
		valid = false
		error = fmt.Errorf("the expression %s is invalid", exp)
		return
	}

	//Validate special characters
	//if ContainsSpecialCharacters(exp) {
	//	valid = false
	//	error = fmt.Errorf("the expression %s is invalid", exp)
	//	return
	//}

	return
}

func IsLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func ContainsParentheses(s string) bool {
	if strings.ContainsRune(s, '(') || strings.ContainsRune(s, ')') {
		return true
	}
	return false
}

func ContainsSpecialCharacters(s string) bool {
	for _, r := range s {
		if unicode.IsSymbol(r) {
			return true
		}
	}
	return false
}
