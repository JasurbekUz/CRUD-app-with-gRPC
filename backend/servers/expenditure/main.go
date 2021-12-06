package main

import (
	"log"
	"net"
	//"context"
	"wallet/database"
	"gorm.io/gorm"
	"google.golang.org/grpc"
	"wallet/proto_files/expenditure"
)

type Server struct {

	DB *gorm.DB
	expenditure.UnimplementedExpenditureServiceServer
}

func main () {

	listen, err := net.Listen("tcp", ":4002")

	if err != nil {
		log.Fatalf("listenning error: %v", err)
	}

	newServer := grpc.NewServer()

	db := database.InitDatabase()

	expenditure.RegisterExpenditureServiceServer(newServer, &Server{DB: db})

	log.Printf("server listening port: %v", listen.Addr())

	if err := newServer.Serve(listen); err != nil {

		log.Fatalf("newServer serving error: %v", err)
	}
}