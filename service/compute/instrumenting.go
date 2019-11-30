package compute

import (
	"github.com/go-kit/kit/metrics"
	"time"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingMiddleware(counter metrics.Counter, latency metrics.Histogram, countResult metrics.Histogram, s Service) Service {
	return &instrumentingMiddleware{
		requestCount:   counter,
		requestLatency: latency,
		countResult:    countResult,
		next:           s,
	}
}

func (mw instrumentingMiddleware) ProcessInfixToPostfix(expression string) (infix string, postfix string, result int64, error error) {
	defer func(begin time.Time) {
		lvs := []string{"Method", "ProcessInfixToPostfix", "Input_expression", expression}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now().UTC())
	infix, postfix, result, error = mw.next.ProcessInfixToPostfix(expression)
	return
}
