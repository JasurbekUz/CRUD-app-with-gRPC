package main

import (
	"log"
	"net"
	//"context"
	"wallet/database"
	"gorm.io/gorm"
	"google.golang.org/grpc"
	"wallet/proto_files/cash"
)

type Server struct {

	DB *gorm.DB
	cash.UnimplementedCashServiceServer
}

func main () {

	listen, err := net.Listen("tcp", ":4001")

	if err != nil {

		log.Fatalf("listenning error: %v", err)
	}

	newServer := grpc.NewServer()

	db := database.InitDatabase()

	cash.RegisterCashServiceServer(newServer, &Server{DB: db})

	log.Printf("server listening port: %v", listen.Addr())

	if err := newServer.Serve(listen); err != nil {

		log.Fatalf("newServer serving error: %v", err)
	}
}