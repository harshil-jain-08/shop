package Repo

import (
	"shop/Config"
	"shop/Models"
)

type ProdRepository interface {
	SaveProduct(product *Models.Product) (err error)
}

type prodRepo struct{}

func NewRepo() ProdRepository {
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
