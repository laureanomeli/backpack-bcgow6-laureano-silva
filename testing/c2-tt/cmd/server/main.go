package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/laureanomeli/backpack-bcgow6-laureano-silva/testing/c2-tt/cmd/server/docs"
	"github.com/laureanomeli/backpack-bcgow6-laureano-silva/testing/c2-tt/cmd/server/handler"
	"github.com/laureanomeli/backpack-bcgow6-laureano-silva/testing/c2-tt/internal/transactions"
	"github.com/laureanomeli/backpack-bcgow6-laureano-silva/testing/c2-tt/pkg/store"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Test API
// @version 1.0
// @description this API Handle transactions
// @licence.name Apache 2.0
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error accesing .env")
	}

	db := store.NewStore("./transactions.json")

	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)

	t := handler.NewTransaction(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pathGroup := router.Group("/transactions")
	pathGroup.POST("/", t.Store())
	pathGroup.GET("/", t.GetAll())
	pathGroup.PUT("/:id", t.Update())
	pathGroup.DELETE("/:id", t.Delete())
	pathGroup.PATCH("/:id", t.UpdateCodeAndAmount())
	if err := router.Run(); err != nil {
		panic(1)
	}
}
