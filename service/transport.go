package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
)

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V string `json:"v"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

type publishMessageRequest struct {
	S string `json:"s"`
}

type publishMessageResponse struct {
	V string `json:"v"`
}

type sdResponse struct {
	P string `json:"pong"`
}

func makeUppercaseEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v}, err
		}
		return uppercaseResponse{v}, nil
	}
}

func makeCountEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}

func makePublishMessageEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(publishMessageRequest)
		svc.PublishMessage(req.S)
		return publishMessageResponse{req.S}, nil
	}
}

func makeServiceDiscoveryEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return sdResponse{"pong"}, nil
	}
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodePublishMessageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request publishMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeServiceDiscoveryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
