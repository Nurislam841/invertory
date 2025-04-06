package usecase

import (
	"inventory/internal/model"
	"inventory/internal/repository"
)

type ProductUsecase struct {
	repo *repository.ProductRepository
}

func NewProductUsecase(repo *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		repo: repo,
	}
}

func (u *ProductUsecase) Create(product *model.Product) error {
	return u.repo.Create(product)
}

func (u *ProductUsecase) GetByID(id int) (*model.Product, error) {
	return u.repo.GetByID(id)
}

func (u *ProductUsecase) Update(product *model.Product) error {
	return u.repo.Update(product)
}

func (u *ProductUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u *ProductUsecase) GetAll() ([]model.Product, error) {
	return u.repo.GetAll()
}
