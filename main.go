package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"github.com/uchijo/walica-clone-backend/data"
	"github.com/uchijo/walica-clone-backend/presenter"
	apipb "github.com/uchijo/walica-clone-backend/proto/proto/api"
	"github.com/uchijo/walica-clone-backend/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	util.LoadEnv()
	util.ConnectToDB()
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("faled to listen:", err)
	}

	s := grpc.NewServer()
	repository := data.RepositoryImpl{}
	apipb.RegisterWalicaCloneApiServer(s, presenter.NewServer(repository))
	log.Println("serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = apipb.RegisterWalicaCloneApiHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	withCors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}).Handler(gwmux)

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: withCors,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
