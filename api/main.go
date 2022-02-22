package main

import (
	"go_xstore/api/models/storeitem"
	"go_xstore/api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	_storeitems := storeitem.New()
	r := gin.Default()
	
	r.GET("/ping", handlers.PingGet())
	r.GET("/items", handlers.ItemsGet(_storeitems))
	r.POST("/item", handlers.ItemPost(_storeitems))

	r.Run()

}