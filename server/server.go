package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	pbAuth "github.com.tw/intri-type/api/auth"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// GCore ...
type GCore struct {
	wg sync.WaitGroup
}

// New creates new server GCore
func New() *GCore {
	return &GCore{}
}

// Start starts server
func (g *GCore) Start() {
	g.wg.Add(1)
	go func() {
		log.Fatal(g.startGRPC())
		g.wg.Done()
	}()
	g.wg.Add(1)
	go func() {
		log.Fatal(g.startREST())
		g.wg.Done()
	}()
}
func (g *GCore) WaitStop() {
	g.wg.Wait()
}
func (g *GCore) startGRPC() error {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pbAuth.RegisterAuthenticationServer(grpcServer, g)
	grpcServer.Serve(lis)
	return nil
}
func (g *GCore) startREST() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pbAuth.RegisterAuthenticationHandlerFromEndpoint(ctx, mux, ":8080", opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8090", mux)
}

// Login API
func (g *GCore) Login(ctx context.Context, r *pbAuth.LoginRequest) (*pbAuth.LoginReply, error) {
	// Maybe it can be moved to middleware
	if err := r.Validate(); err != nil {
		return nil, err
	}

	// An authentication check for acc & pwd
	var result = true
	log.Print(r.Account, " login - ", result)

	return &pbAuth.LoginReply{
		Result: result,
	}, nil
}

// Logout API
func (g *GCore) Logout(ctx context.Context, r *empty.Empty) (*pbAuth.LogoutReply, error) {

	var result = true
	log.Print("Logout - ", result)

	return &pbAuth.LogoutReply{
		Result: result,
	}, nil
}
