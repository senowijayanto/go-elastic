package repositories

import (
	"go-elastic/config"
	"go-elastic/models"
)

type ProductRepository interface {
	Create(product models.Product) error
	GetAll() ([]models.Product, error)
	Update(id uint, product models.Product) error
	Delete(id uint) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Create(product models.Product) error {
	return config.DB.Create(&product).Error
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Find(&products).Error
	return products, err
}

func (r *productRepository) Update(id uint, product models.Product) error {
	return config.DB.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Product{}, id).Error
}
