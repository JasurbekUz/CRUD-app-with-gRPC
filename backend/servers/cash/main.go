package main

import (
	"log"
	"net"
	"context"
	"errors"
	"wallet/database"
	"gorm.io/gorm"
	"google.golang.org/grpc"
	"wallet/proto_files/cash"
)

type Response struct {

	Id int
	Amo float64
	Sum string
	Time string
}

type Server struct {

	DB *gorm.DB
	cash.UnimplementedCashServiceServer
}

func (s Server) PostNewCash (ctx context.Context, in *cash.PostCash) (newCash *cash.Cash, err error) {

	log.Println(in)
	

	err = s.DB.Raw(
		database.CREATE_NEW_CASH,
		in.GetIdentificator().GetUsername(),
		in.GetIdentificator().GetPassword(),
		in.GetAmount(),
		in.GetSummary(),
		).Scan(&newCash).Error

	errors.Is(err, gorm.ErrRecordNotFound)
	
	return newCash, err
}

func (s Server) GetListOfCashe (ctx context.Context, in *cash.Identificator) (cashList *cash.ListOfCashe, err error) {

	res := s.DB.Raw(database.SELECT_USER_CASH_LIST, in.GetUsername(), in.GetPassword()).Scan(&cashList.Cashes)

	log.Println(res.RowsAffected)

	return nil, nil
}

func main () {

	listen, err := net.Listen("tcp", ":4001")

	if err != nil {

		log.Fatalf("listenning error: %v", err)
	}

	newServer := grpc.NewServer()

	db := database.InitDatabase()

	cash.RegisterCashServiceServer(newServer, &Server{DB: db})

	log.Printf("cash server is listening port: %v", listen.Addr())

	if err := newServer.Serve(listen); err != nil {

		log.Fatalf("newServer serving error: %v", err)
	}
}