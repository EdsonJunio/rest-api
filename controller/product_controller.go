package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"rest-api/configuration/logger"
	"rest-api/configuration/rest_err"
	"rest-api/model"
	"rest-api/usecase"
	"strconv"
)

type ProductController struct {
	productUseCase *usecase.ProductUsecase
}

func NewProductController(usecase *usecase.ProductUsecase) *ProductController {
	return &ProductController{
		productUseCase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	logger.Info("Received GET /products request")

	products, err := p.productUseCase.GetProducts()
	if err != nil {
		logger.Error("Failed to retrieve products", err)
		restErr := rest_err.NewInternalServerError("Error retrieving products")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  products,
		"count": len(products),
	})
}

func (p *ProductController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("productId")
	logger.Info("Received GET /product/:productId request", zap.String("productId", id))

	if id == "" {
		restErr := rest_err.NewBadRequestError("productId is required")
		logger.Error("Product ID missing in request", nil)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	productId, convErr := strconv.Atoi(id)
	if convErr != nil {
		restErr := rest_err.NewBadRequestError("productId must be a number")
		logger.Error("Invalid productId parameter", convErr, zap.String("param", id))
		ctx.JSON(restErr.Code, restErr)
		return
	}

	product, err := p.productUseCase.GetProductByID(productId)
	if err != nil {
		restErr := rest_err.NewNotFoundError("product not found")
		logger.Error("Product not found", err, zap.Int("productId", productId))
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	logger.Info("Received POST /product request")

	var product model.Product
	if err := ctx.BindJSON(&product); err != nil {
		restErr := rest_err.NewBadRequestError("Invalid JSON payload")
		logger.Error("Invalid JSON payload in CreateProduct", err)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		restErr := rest_err.NewInternalServerError("Could not create product")
		logger.Error("Failed to create product", err)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": insertedProduct,
	})
}

func (p *ProductController) UpdateProductByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	logger.Info("Received PUT /products/:id request", zap.String("id", idParam))

	id, convErr := strconv.Atoi(idParam)
	if convErr != nil {
		restErr := rest_err.NewBadRequestError("id must be a number")
		logger.Error("Invalid id parameter", convErr, zap.String("param", idParam))
		ctx.JSON(restErr.Code, restErr)
		return
	}

	var input model.Product
	if err := ctx.ShouldBindJSON(&input); err != nil {
		restErr := rest_err.NewBadRequestError("invalid request body")
		logger.Error("Invalid request body in UpdateProductByID", err)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	updatedProduct, err := p.productUseCase.UpdateProductByID(id, input)
	if err != nil {
		if err.Error() == "product not found" {
			restErr := rest_err.NewNotFoundError("product not found")
			logger.Error("Product not found for update", err, zap.Int("id", id))
			ctx.JSON(restErr.Code, restErr)
			return
		}
		restErr := rest_err.NewInternalServerError("error updating product")
		logger.Error("Failed to update product", err, zap.Int("id", id))
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": updatedProduct,
	})
}

func (p *ProductController) DeleteProductByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	logger.Info("Received DELETE /products/:id request", zap.String("id", idParam))

	productId, convErr := strconv.Atoi(idParam)
	if convErr != nil {
		restErr := rest_err.NewBadRequestError("id must be a number")
		logger.Error("Invalid id parameter in DeleteProductByID", convErr, zap.String("param", idParam))
		ctx.JSON(restErr.Code, restErr)
		return
	}

	_, err := p.productUseCase.DeleteProductByID(productId)
	if err != nil {
		if err.Error() == "product not found" {
			restErr := rest_err.NewNotFoundError("product not found")
			logger.Error("Product not found for deletion", err, zap.Int("id", productId))
			ctx.JSON(restErr.Code, restErr)
			return
		}
		restErr := rest_err.NewInternalServerError("error deleting product")
		logger.Error("Failed to delete product", err, zap.Int("id", productId))
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "product deleted successfully",
	})
}
