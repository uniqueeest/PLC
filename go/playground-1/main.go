package main

import (
	"main/user"

	"github.com/gin-gonic/gin"
)

func main() {
	// user 인스턴스 초기화
	repo := user.UserRepository("./db")
	service := user.UserService(repo)
	handler := user.UserHandler(service)

	router := gin.Default()

	handler.SetupRoutes(router)

	router.Run(":8080")
}