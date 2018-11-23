package endpoints

import (
	"context"
	"sample/services"
	"sample/structures"

	"github.com/go-kit/kit/endpoint"
)

// MakeGreetingEndpoint creates Endpoint for Greeting method of GreetingService
func MakeGreetingEndpoint(svc services.GreetingService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(structures.GreetingRequest)
		v, err := svc.Greet(req.Name)
		if err != nil {
			return structures.GreetingResponse{v, err.Error()}, nil
		}
		return structures.GreetingResponse{v, ""}, nil
	}
}

// MakeIntEndpoint cerates Endpoint for GetInt method of GreetingService
func MakeIntEndpoint(svc services.GreetingService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v := svc.GetInt()
		return structures.IntResponse{v}, nil
	}
}
