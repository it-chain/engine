package api_gateway

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"encoding/json"

	"github.com/go-kit/kit/examples/shipping/cargo"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(bs TransactionQueryApi, logger kitlog.Logger) http.Handler {

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
	}

	findAllUncommittedTransactionsHandler := kithttp.NewServer(
		makeFindUncommittedTransactionsEndpoint(bs),
		decodeFindAllUncommittedTransactionsRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/transactions", findAllUncommittedTransactionsHandler).Methods("GET")

	return r
}

// this return nil because this request body is empty
func decodeFindAllUncommittedTransactionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// DTO
	return nil, nil
}

// serialize
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case cargo.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	//case ErrInvalidArgument:
	//	w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
