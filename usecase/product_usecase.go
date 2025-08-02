package usecase

import (
	"errors"
	"rest-api/configuration/rest_err"
	"rest-api/model"
	"rest-api/repository"
)

type ProductUsecase struct {
	repository *repository.ProductRepository
}

func NewProductUsecase(repo *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	products, err := pu.repository.GetProducts()
	if err != nil {
		return nil, rest_err.NewInternalServerError("Error retrieving products")
	}
	return products, nil
}

func (pu *ProductUsecase) GetProductByID(id int) (*model.Product, error) {
	product, err := pu.repository.GetProductByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return nil, rest_err.NewNotFoundError("product not found")
		}
		return nil, rest_err.NewInternalServerError("Error retrieving product")
	}
	return product, nil
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, rest_err.NewInternalServerError("Could not create product")
	}
	product.ID = productId
	return product, nil
}

func (pu *ProductUsecase) UpdateProductByID(id int, data model.Product) (*model.Product, error) {
	updatedProduct, err := pu.repository.UpdateProductByID(id, data)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return nil, rest_err.NewNotFoundError("product not found")
		}
		return nil, rest_err.NewInternalServerError("Error updating product")
	}
	return updatedProduct, nil
}

func (pu *ProductUsecase) DeleteProductByID(id int) (*model.Product, error) {
	deletedProduct, err := pu.repository.DeleteProductByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return nil, rest_err.NewNotFoundError("product not found")
		}
		return nil, rest_err.NewInternalServerError("Error deleting product")
	}
	return deletedProduct, nil
}
