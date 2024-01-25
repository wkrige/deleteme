package main

import (
	greetv1RPC "buf.build/gen/go/willkrige/deleteme/grpc/go/proto/helloworld/v1/helloworldv1grpc"
	"buf.build/gen/go/willkrige/deleteme/protocolbuffers/go/proto/helloworld/v1"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	logrus.Info("hello-world server")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3030))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetv1RPC.RegisterHelloWorldServiceServer(s, &HelloWorldGRPCServer{})

	log.Printf("hello-world server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	logrus.Info("goodbye world")
}

type HelloWorldGRPCServer struct {
	greetv1RPC.UnimplementedHelloWorldServiceServer
}

func (s *HelloWorldGRPCServer) GetHelloWorld(ctx context.Context, request *helloworldv1.GetHelloWorldRequest) (*helloworldv1.GetHelloWorldResponse, error) {
	logrus.Info("get-hello-world called")

	greeting := ""

	switch request.GetLanguage() {
	case helloworldv1.LanguageType_LANGUAGE_TYPE_ENGLISH:
		greeting = "Hello world!"
	case helloworldv1.LanguageType_LANGUAGE_TYPE_ITALIAN:
		greeting = "Ciao mondo!"
	case helloworldv1.LanguageType_LANGUAGE_TYPE_AFRIKAANS:
		greeting = "Hallo wereld!"
	default:
		return &helloworldv1.GetHelloWorldResponse{Greeting: "test"}, errors.New("unknown language")
	}

	return &helloworldv1.GetHelloWorldResponse{Greeting: greeting}, nil
}
