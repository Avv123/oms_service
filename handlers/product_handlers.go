package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oms/models"
	"oms/repository"
)

func CreateProductById(ctx *gin.Context) {

	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
	}
	result, err := repository.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"product": result})

}
func GetAllProducts(ctx *gin.Context) {
	//var products []models.Product
	products, err := repository.FindAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, products)
}

func GetProductByid(ctx *gin.Context) {
	id := ctx.Param("id")
	product, _ := repository.GetProduct(id)
	if product == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"product": product})

}
