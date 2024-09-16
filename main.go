package main

import (
	"github.com/fakhrizalmus/belajargolang/controllers/productcontroller"
	"github.com/fakhrizalmus/belajargolang/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Edit)
	r.DELETE("/api/product/:id", productcontroller.Delete)

	r.Run()
}
