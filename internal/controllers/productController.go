package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitormoschetta/go-rest-api/internal/models"
	"github.com/vitormoschetta/go-rest-api/internal/repositories"
)

type ProductController struct {
	ProductRepository repositories.IProductRepository
}

func NewProductController(productRepository repositories.IProductRepository) *ProductController {
	return &ProductController{
		ProductRepository: productRepository,
	}
}

func (c *ProductController) FindAll(context *gin.Context) {
	products, err := c.ProductRepository.FindAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, products)
}

func (c *ProductController) FindByID(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	product, err := c.ProductRepository.FindByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if product.ID == "" {
		context.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	context.JSON(http.StatusOK, product)
}

func (c *ProductController) Create(context *gin.Context) {
	var product models.Product
	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product = models.NewProduct(product.Name, product.Price)
	msgs := product.Validate()
	if len(msgs) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": msgs})
		return
	}
	product, err := c.ProductRepository.Create(product)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, product)
}

func (c *ProductController) Update(context *gin.Context) {
	id := context.Param("id")
	product, err := c.ProductRepository.FindByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if product.ID == "" {
		context.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	var newProduct models.Product
	if err := context.ShouldBindJSON(&newProduct); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newProduct.ID = product.ID
	product.Update(newProduct.Name, newProduct.Price)
	msgs := product.Validate()
	if len(msgs) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": msgs})
		return
	}
	err = c.ProductRepository.Update(product)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, newProduct)
}

func (c *ProductController) Delete(context *gin.Context) {
	id := context.Param("id")
	product, err := c.ProductRepository.FindByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if product.ID == "" {
		context.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	err = c.ProductRepository.Delete(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
