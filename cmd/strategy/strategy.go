package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/statistico/statistico-grpc-gateway/proto/gen"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

func main() {
	host := os.Getenv("STATISTICO_STRATEGY_HOST")
	port := os.Getenv("STATISTICO_STRATEGY_PORT")

	conn, err := grpc.Dial(host + ":" + port, grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}

	mux := runtime.NewServeMux()

	err = statistico.RegisterStrategyServiceHandler(context.Background(), mux, conn)

	if err != nil {
		log.Fatalln(err)
	}

	log.Fatal(http.ListenAndServe(":8080", mux))
}
