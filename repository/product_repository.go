package repository

import (
	"errors"
	"gorm.io/gorm"
	"rest-api/configuration/logger"
	"rest-api/model"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	var products []model.Product
	if err := pr.db.Order("id ASC").Find(&products).Error; err != nil {
		logger.Error("Error retrieving products from database", err)
		return nil, err
	}
	return products, nil
}

func (pr *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	var product model.Product
	if err := pr.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		logger.Error("Error retrieving product from database", err)
		return nil, err
	}
	return &product, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	if err := pr.db.Create(&product).Error; err != nil {
		logger.Error("Error creating product in database", err)
		return 0, err
	}
	return product.ID, nil
}

func (pr *ProductRepository) UpdateProductByID(id int, data model.Product) (*model.Product, error) {
	var product model.Product
	if err := pr.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		logger.Error("Error retrieving product for update in database", err)
		return nil, err
	}

	product.Name = data.Name
	product.Price = data.Price

	if err := pr.db.Save(&product).Error; err != nil {
		logger.Error("Error saving updated product in database", err)
		return nil, err
	}
	return &product, nil
}

func (pr *ProductRepository) DeleteProductByID(id int) (*model.Product, error) {
	var product model.Product
	if err := pr.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		logger.Error("Error retrieving product for delete in database", err)
		return nil, err
	}

	if err := pr.db.Delete(&product).Error; err != nil {
		logger.Error("Error deleting product from database", err)
		return nil, err
	}
	return &product, nil
}
