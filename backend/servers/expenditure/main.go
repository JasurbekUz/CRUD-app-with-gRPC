package main

import (
	"log"
	"net"
	"errors"
	"context"
	"wallet/database"
	"gorm.io/gorm"
	"google.golang.org/grpc"
	"wallet/proto_files/expenditure"
)

type Server struct {

	DB *gorm.DB
	expenditure.UnimplementedExpenditureServiceServer
}

func (s Server) PostNewExpenditure (ctx context.Context, in *expenditure.PostExpenditure) (exp *expenditure.Expenditure, err error) {

	err = s.DB.Raw(database.CREATE_NEW_EXP,
		in.GetIdentificator().GetUsername(), 
		in.GetIdentificator().GetPassword(),
		in.GetAmount(),
		in.GetSummary(),
		).Scan(&exp).Error

	errors.Is(err, gorm.ErrRecordNotFound)

	return exp, err
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