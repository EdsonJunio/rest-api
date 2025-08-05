package usecase

import (
	"errors"
	"rest-api/configuration/rest_err"
	"rest-api/model"
	"rest-api/repository"
)

type ProductUsecase struct {
	repo *repository.ProductRepository
}

func NewProductUsecase(repo *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (u *ProductUsecase) GetProducts() ([]model.Product, error) {
	return u.repo.GetProducts()
}

func (u *ProductUsecase) GetProductByID(id int) (*model.Product, error) {
	if id <= 0 {
		return nil, rest_err.NewBadRequestError("invalid ID")
	}

	product, err := u.repo.GetProductByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return nil, rest_err.NewNotFoundError("product not found")
		}
		return nil, rest_err.NewInternalServerError("could not retrieve product")
	}

	return product, nil
}

func (u *ProductUsecase) CreateProduct(p model.Product) (model.Product, error) {
	if err := validateProduct(p); err != nil {
		return model.Product{}, err
	}

	id, err := u.repo.CreateProduct(p)
	if err != nil {
		return model.Product{}, rest_err.NewInternalServerError("could not create product")
	}

	p.ID = id
	return p, nil
}

func (u *ProductUsecase) UpdateProductByID(id int, p model.Product) (*model.Product, error) {
	if id <= 0 {
		return nil, rest_err.NewBadRequestError("invalid ID")
	}

	if err := validateProduct(p); err != nil {
		return nil, err
	}

	_, err := u.repo.GetProductByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return nil, rest_err.NewNotFoundError("product not found")
		}
		return nil, rest_err.NewInternalServerError("could not fetch existing product")
	}

	updated, err := u.repo.UpdateProductByID(id, p)
	if err != nil {
		return nil, rest_err.NewInternalServerError("could not update product")
	}
	return updated, nil
}

func (u *ProductUsecase) DeleteProductByID(id int) (*model.Product, error) {
	if id <= 0 {
		return nil, rest_err.NewBadRequestError("invalid ID")
	}

	_, err := u.repo.GetProductByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return nil, rest_err.NewNotFoundError("product not found")
		}
		return nil, rest_err.NewInternalServerError("could not fetch product to delete")
	}

	deleted, err := u.repo.DeleteProductByID(id)
	if err != nil {
		return nil, rest_err.NewInternalServerError("could not delete product")
	}
	return deleted, nil
}

func validateProduct(p model.Product) *rest_err.RestErr {
	if p.Name == "" {
		return rest_err.NewBadRequestError("product name is required")
	}
	if p.Price <= 0 {
		return rest_err.NewBadRequestError("product price must be greater than zero")
	}
	return nil
}
