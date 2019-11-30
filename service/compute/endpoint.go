package compute

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	"go-kit-base/infrastructure/validation"
)

// LOAD PRODUCT TRANSACTIONS

// Struct definition
// productTransactionRequest represents an HTTP request to read a specific product and user ID
type infixRequest struct {
	Expression string `json:"exp"`
}

// productTransactionsResponse represents an HTTP response containing set of Transactions found when fetching
type postfixResponse struct {
	Infix   string `json:"infix,omitempty"`
	Postfix string `json:"postfix,omitempty"`
	Result  int64  `json:"result,omitempty"`
	Error   string `json:"error,omitempty"`
}

// Endpoint creation
func makeInfixProcessingEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*infixRequest)

		//TODO Applying the validation in here is a mistake
		// A better approach is to use a Validation middleware which acts as a Circuit Breaker for bad requests
		//Expression validator
		validator := validation.New()
		valid, err := validator.ValidateExpression(req.Expression)
		if err != nil && valid {
			inf, pos, res, pErr := s.ProcessInfixToPostfix(req.Expression)
			if pErr != nil {
				pErr = errors.New("")
			}
			return postfixResponse{
				Infix:   inf,
				Postfix: pos,
				Result:  res,
				Error:   err.Error(),
			}, nil
		} else {
			return postfixResponse{
				Infix:   req.Expression,
				Postfix: "",
				Result:  0,
				Error:   err.Error(),
			}, nil
		}

	}
}
