package main

import (
	"auth_user/app/config"
	"auth_user/app/db"
	"auth_user/app/pb"
	"auth_user/app/services"
	"auth_user/app/utils"
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", fmt.Sprintf("localhost:%d", 50051))

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}
	rg := gin.Default()
	s.StartHttp(rg.Group("user/"))
	grpcServer := grpc.NewServer()
	go rg.Run(":8082")
	pb.RegisterAuthServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
