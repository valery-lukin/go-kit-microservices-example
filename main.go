package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sample/endpoints"
	"sample/services"
	"sample/structures"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := services.GreetingServiceImpl{}

	greetingHandler := httptransport.NewServer(
		endpoints.MakeGreetingEndpoint(svc),
		decodeGreetingRequest,
		encodeResponse)

	intHandler := httptransport.NewServer(
		endpoints.MakeIntEndpoint(svc),
		decodeIntRequest,
		encodeResponse)

	http.Handle("/greet", greetingHandler)
	http.Handle("/get-int", intHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func decodeGreetingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request structures.GreetingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeIntRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request structures.IntRequest
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
