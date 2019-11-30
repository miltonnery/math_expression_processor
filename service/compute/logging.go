package compute

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingMiddleware(logger log.Logger, s Service) Service {
	return &loggingMiddleware{logger, s}
}

func (mw loggingMiddleware) ProcessInfixToPostfix(expression string) (infix string, postfix string, result int64, error error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "ProcessInfixToPostfix",
			"Infix_expression", expression,
			"PostFix_expression", postfix,
			"Computed_result", result,
			"Error?", error,
			"Took", time.Since(begin))
	}(time.Now().UTC())
	infix, postfix, result, error = mw.next.ProcessInfixToPostfix(expression)
	return
}
