package main

import (
	"github.com/gin-gonic/gin"
	"inventory/internal/database"
	"inventory/internal/handler"
	"inventory/internal/repository"
	"inventory/internal/usecase"
)

func main() {
	db, err := database.InitSQLite()
	if err != nil {
		panic(err)
	}
	repo := repository.NewProductRepository(db)

	uc := usecase.NewProductUsecase(repo)

	h := handler.NewProductHandler(uc)

	r := gin.Default()
	h.RegisterRoutes(r)

	r.Run(":8081")
}
