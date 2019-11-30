package compute

import (
	"go-kit-base/infrastructure/validation"
)

type validationMiddleware struct {
	validator validation.Validator
	next      Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewValidationMiddleware(validator validation.Validator, s Service) Service {
	return &validationMiddleware{validator, s}
}

func (mw validationMiddleware) ProcessInfixToPostfix(exp string) (infix string, postfix string, result int64, error error) {
	//Expression validator
	infix = exp
	postfix = ""
	result = 0
	valid, err := mw.validator.ValidateExpression(exp)
	if err != nil {
		error = err
		panic(error)
		infix, postfix, result, error = mw.next.ProcessInfixToPostfix(exp)
	}
	if valid {
		infix, postfix, result, error = mw.next.ProcessInfixToPostfix(exp)
	}

	return
}
