package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/longjoy/micro-go-course/register/endpoint"
	"github.com/pkg/errors"
	"net/http"
	"os"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

// MakeHttpHandler make http handler use mux

func MakeHttpHandler(ctx context.Context, endpoints * endpoint.RegisterEndpoints) http.Handler {
	r := mux.NewRouter()
	kitLog := log.NewLogfmtLogger(os.Stderr)
	kitLog = log.With(kitLog, "ts", log.DefaultTimestampUTC)
	kitLog = log.With(kitLog, "caller", log.DefaultCaller)
	options := []kitHttp.ServerOption{
		kitHttp.ServerErrorHandler(transport.NewLogErrorHandler(kitLog)),
		kitHttp.ServerErrorEncoder(encodeError),
	}
	r.Methods("GET").Path("/health").Handler(kitHttp.NewServer(
			endpoints.HealthCheckEndpoint,
			decodeHealthCheckRequest,
			encodeJSONResponse,
			options...
		))

	r.Methods("GET").Path("/discovery/name").Handler(kitHttp.NewServer(
			endpoints.DiscoveryEndpoint,
			decodeHealthCheckRequest,
			encodeJSONResponse,
			options...
		))
	return r
}

func decodeDiscoveryRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	serviceName := r.URL.Query().Get("serviceName")

	if serviceName == ""{
		return nil, ErrorBadRequest
	}
	return endpoint.DiscoveryRequest{
		ServiceName:serviceName,
	}, nil
}

func decodeHealthCheckRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return endpoint.HealthRequest{}, nil
}



func encodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}


func encodeError(_ context.Context, err error, w http.ResponseWriter)  {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}