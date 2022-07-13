package Service

import (
	"shop/Models"
	"shop/Repo"
	"shop/dto"
)

var ProdSer ProductService

type ProductService interface {
	CreateProduct(data *dto.Product) (product *Models.Product, err error)
	UpdateProduct(newProduct Models.Product) (product *Models.Product, err error)
}

type prodService struct {
	repo Repo.ProdRepository
}

func NewService(prodRepository Repo.ProdRepository) {
	ProdSer = &prodService{repo: prodRepository}
}

func (s *prodService) CreateProduct(data *dto.Product) (product *Models.Product, err error) {
	model := &Models.Product{
		Name:     data.Name,
		Price:    data.Price,
		Quantity: data.Quantity,
	}
	err = s.repo.SaveProduct(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *prodService) UpdateProduct(newProduct Models.Product) (product *Models.Product, err error) {

	panic("implement me")
}
