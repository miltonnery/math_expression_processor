package compute

import (
	"context"
	"encoding/json"
	"errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

func MakeHandler(s Service) http.Handler {

	// Define all endpoints
	infixProcessingEndpoint := makeInfixProcessingEndpoint(s)

	// Define all http servers
	infixProcessingHandler := kithttp.NewServer(
		infixProcessingEndpoint,
		decodeExpression,
		encodeResult)

	r := mux.NewRouter()
	r.Handle("/evaluate", infixProcessingHandler).Methods("POST")
	return r
}

// ENCODERS AND DECODERS
//GENERAL
// General request decoder
func decodeJSONRequest(to interface{}, r *http.Request) (interface{}, error) {
	d, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(d, &to)
	return to, err
}

// General response encoder
func encodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	d, err := json.Marshal(response)
	_, err = w.Write(d)
	return err
}

// General JSON error encoder
func encodeJSONError(_ context.Context, err error, w http.ResponseWriter) {

	if strings.Contains(err.Error(), "is invalid") {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
/*	_ = json.NewEncoder(w).Encode(postfixResponse{
		Infix:   "",
		Postfix: "",
		Result:  0,
		Error:   err.Error(),
	})*/
}

//SPECIFIC Decoders and encoders
//Infix expressions for calculation
func decodeExpression(_ context.Context, r *http.Request) (interface{}, error) {
	var req infixRequest
	return decodeJSONRequest(&req, r)
}
func encodeResult(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(postfixResponse)
	if res.Error != "" {
		encodeJSONError(ctx, errors.New(res.Error), w)
	}
	return encodeJSONResponse(ctx, w, res)
}
