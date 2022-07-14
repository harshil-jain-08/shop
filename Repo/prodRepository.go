package Repo

import (
	"shop/Config"
	"shop/Models"
)

type ProdRepository interface {
	SaveProduct(product *Models.Product) (err error)
	FindProduct(product *Models.Product, id int) (err error)
	GetAllProducts(products *[]Models.Product) (err error)
}

type prodRepo struct{}

func NewProductRepo() ProdRepository {
	return &prodRepo{}
}

func (r *prodRepo) SaveProduct(product *Models.Product) (err error) {
	db := Config.DB.Save(product)
	err = db.Error
	if err != nil {
		return err

	}
	return nil
}

func (r *prodRepo) FindProduct(product *Models.Product, id int) (err error) {
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func (r *prodRepo) GetAllProducts(products *[]Models.Product) (err error) {
	if err = Config.DB.Find(products).Error; err != nil {
		return err
	}
	return nil
}
