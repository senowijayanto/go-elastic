package services

import (
	"go-elastic/models"
	"go-elastic/repositories"
)

type ProductService interface {
	CreateProduct(product models.Product) error
	GetAllProducts() ([]models.Product, error)
	UpdateProduct(id uint, product models.Product) error
	DeleteProduct(id uint) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) CreateProduct(product models.Product) error {
	return s.repo.Create(product)
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) UpdateProduct(id uint, product models.Product) error {
	return s.repo.Update(id, product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}
