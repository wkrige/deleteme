package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	greetv1RPC "buf.build/gen/go/willkrige/deleteme/grpc/go/proto/helloworld/v1/helloworldv1grpc"
	"buf.build/gen/go/willkrige/deleteme/protocolbuffers/go/proto/helloworld/v1"

	"log"
	"time"
)

func main() {
	logrus.Info("hello-world client")

	conn, err := grpc.Dial("localhost:3030", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := greetv1RPC.NewHelloWorldServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetHelloWorld(ctx, &helloworldv1.GetHelloWorldRequest{Language: helloworldv1.LanguageType_LANGUAGE_TYPE_ENGLISH})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetGreeting())

	r, err = c.GetHelloWorld(ctx, &helloworldv1.GetHelloWorldRequest{Language: helloworldv1.LanguageType_LANGUAGE_TYPE_ITALIAN})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetGreeting())

	r, err = c.GetHelloWorld(ctx, &helloworldv1.GetHelloWorldRequest{Language: helloworldv1.LanguageType_LANGUAGE_TYPE_AFRIKAANS})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetGreeting())

	r, err = c.GetHelloWorld(ctx, &helloworldv1.GetHelloWorldRequest{Language: helloworldv1.LanguageType_LANGUAGE_TYPE_UNSPECIFIED})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetGreeting())
}
