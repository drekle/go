package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/drekle/go/rpc/urlValidator/go/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func (s server) Register(ctx context.Context, req *api.UrlRequest) (*api.UrlResponse, error) {
	var resp api.UrlResponse
	var err error

	url := req.Url
	if !strings.HasPrefix(url, "http") {
		//No scheme passed
		url = "http://" + url
	}
	log.Println("Received Request!")
	r, err := http.Get(url)
	if err != nil || r.StatusCode != http.StatusOK {
		resp.Valid = false
	} else {
		resp.Valid = true
	}
	return &resp, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterURLServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
