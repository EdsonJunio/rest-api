package usecase

import (
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
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) GetProductByID(id int) (*model.Product, error) {
	return pu.repository.GetProductByID(id)
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = productId
	return product, nil
}

func (pu *ProductUsecase) UpdateProductByID(id int, data model.Product) (*model.Product, error) {
	return pu.repository.UpdateProductByID(id, data)
}

func (pu *ProductUsecase) DeleteProductByID(id int) (*model.Product, error) {
	return pu.repository.DeleteProductByID(id)
}
