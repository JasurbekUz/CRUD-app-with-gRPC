package main

import (
	"log"
	"context"
	"google.golang.org/grpc"
	"wallet/proto_files/user"
	"wallet/proto_files/cash"
	"wallet/proto_files/expenditure"

	//"http/net"
	"github.com/gin-gonic/gin"
)

var newUserClient user.UserSeviceClient
var newCashClient cash.CashServiceClient
var newExpenditureClient expenditure.ExpenditureServiceClient

// USER CONTROLLERS
func PostUser (ctx *gin.Context) {

	var newUser user.PostUser

	ctx.BindJSON(&newUser)

	result, err := newUserClient.PostNewUser(context.Background(), &newUser)

	if err != nil {
		
		ctx.JSON(403, gin.H {

			"message": "user name is exixts",
		})
	} else {

		ctx.JSON(201, result)

	}	
}

func GetUserInfo (ctx *gin.Context) {

	var identificator user.Identificator

	ctx.BindJSON(&identificator)

	result, err := newUserClient.GetUserBasicInfo(context.Background(), &identificator)

	if err != nil {
		
		ctx.JSON(403, gin.H{
			"message": "username or password error",
		})

	} else {

		ctx.JSON(200, result)
	}
}

// CASH CONTROLLERS
func Income (ctx *gin.Context) {

	var newCash cash.PostCash

	ctx.BindJSON(&newCash)

	result, err := newCashClient.PostNewCash(context.Background(), &newCash)

	if err != nil {
		
		ctx.JSON(403, gin.H {

			"message": "username or password error",
			"message2": "invalid operation!!!",
		})
	} else {

		ctx.JSON(201, result)

	}	
}

func GetListOfIncome (ctx *gin.Context) {

	var identificator cash.Identificator

	ctx.BindJSON(&identificator)

	result, err := newCashClient.GetListOfCashe(context.Background(), &identificator)

	if err != nil {

		ctx.JSON(403, gin.H{
			"message": "username or password error",
		})

	} else {

		ctx.JSON(200, result.Cashes)
	}
}

// EXPENDITURE CONTROLLERS
func ToSpand(ctx *gin.Context) {

	var exp expenditure.PostExpenditure

	ctx.BindJSON(&exp)

	result, err := newExpenditureClient.PostNewExpenditure(context.Background(), &exp)

	if err != nil {

		ctx.JSON(403, gin.H{
			"message": "username or password error",
		})

	} else {

		ctx.JSON(200, result)
	}
}

func main () {

	// CONNECTING WITH USER SERVICE
	connectionWithUserService, err := grpc.Dial("localhost:4000", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Couldn't connect %v", err)
	}

	defer connectionWithUserService.Close()

	newUserClient = user.NewUserSeviceClient(connectionWithUserService)

	// CONNECTING WITH CASH SERVICE
	connectionWithCashService, err := grpc.Dial("localhost:4001", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Couldn't connect %v", err)
	}

	defer connectionWithCashService.Close()

	newCashClient = cash.NewCashServiceClient(connectionWithCashService)

	// CONNECTING WITH CASH SERVICE
	connectionWithExpenditureService, err := grpc.Dial("localhost:4002", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Couldn't connect %v", err)
	}

	defer connectionWithExpenditureService.Close()

	newExpenditureClient = expenditure.NewExpenditureServiceClient(connectionWithExpenditureService)
	// API
	router := gin.Default()

	router.POST("/create_wallet", PostUser)
	router.GET("/get_balace", GetUserInfo)

	router.POST("/income", Income)
	// bu xizmat pulli
	//router.GET("/list_of_income", GetListOfIncome)

	router.POST("/to_spend", ToSpand)
	// bu xizmat pulli
	//router.GET("/list_of_expenditures", )

	router.Run("localhost:8080")
}