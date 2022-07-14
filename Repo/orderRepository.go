package Repo

import (
	"shop/Config"
	"shop/Models"
)

type OrderRepository interface {
	CreateOrder(order *Models.Order) (err error)
	FindOrder(model *Models.Order, orderId int) (err error)
}

type orderRepo struct{}

func NewOrderRepo() OrderRepository {
	return &orderRepo{}
}
func (o *orderRepo) CreateOrder(order *Models.Order) (err error) {
	db := Config.DB.Save(order)
	err = db.Error
	if err != nil {
		return err

	}
	return nil
}

func (r *orderRepo) FindOrder(order *Models.Order, id int) (err error) {
	if err = Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}
