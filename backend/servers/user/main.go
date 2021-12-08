package main

import (
	"log"
	"net"
	"context"
	"errors"
	"wallet/database"
	"gorm.io/gorm"
	"google.golang.org/grpc"
	"wallet/proto_files/user"
)

type Server struct {

	DB *gorm.DB
	user.UnimplementedUserSeviceServer
}

var p_user user.User

func (s Server) PostNewUser(ctx context.Context, in *user.PostUser) (p_user *user.User, err error) {

	err = s.DB.Raw(database.CREATE_NEW_USER, in.GetFullname(), in.GetUsername(), in.GetPassword()).Scan(&p_user).Error

	errors.Is(err, gorm.ErrRecordNotFound)

	return p_user, err
}

func (s Server) GetUserBasicInfo (ctx context.Context, in *user.Identificator) (p_user *user.User, err error) {

	err = s.DB.Raw(database.SELECT_USER_BUDGET_INFO, in.GetUsername(), in.GetPassword()).Scan(&p_user).Error

	errors.Is(err, gorm.ErrRecordNotFound)

	return p_user, err
}

func main () {

	listen, err := net.Listen("tcp", ":4000")

	if err != nil {
		log.Fatalf("listenning error: %v", err)
	}

	newServer := grpc.NewServer()

	db := database.InitDatabase()

	user.RegisterUserSeviceServer(newServer, &Server{DB: db})

	log.Printf("server listening port: %v", listen.Addr())

	if err := newServer.Serve(listen); err != nil {

		log.Fatalf("newServer serving error: %v", err)
	}
}