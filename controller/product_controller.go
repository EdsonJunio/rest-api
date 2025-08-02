package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  products,
		"count": len(products),
	})
}

func (p *ProductController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "productId is required",
		})
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "productId must be a number",
		})
		return
	}

	product, err := p.productUseCase.GetProductByID(productId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON payload",
		})
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create product",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": insertedProduct,
	})
}

func (p *ProductController) UpdateProductByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be a number",
		})
		return
	}

	var input model.Product
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	updatedProduct, err := p.productUseCase.UpdateProductByID(id, input)
	if err != nil {
		if err.Error() == "product not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "product not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error updating product",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": updatedProduct,
	})
}

func (p *ProductController) DeleteProductByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	productId, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be a number",
		})
		return
	}

	_, err = p.productUseCase.DeleteProductByID(productId)
	if err != nil {
		if err.Error() == "product not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "product not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error deleting product",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "product deleted successfully",
	})
}
