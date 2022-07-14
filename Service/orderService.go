package Service

import (
	"shop/Models"
	"shop/Repo"
	"shop/dto"
)

var OrderSer OrderService

type OrderService interface {
	CreateOrder(data *dto.Order) (order *Models.Order, err error)
	FindOrder(id int) (order *Models.Order, err error)
}

type orderService struct {
	repo Repo.OrderRepository
}

func NewOrderService(orderRepository Repo.OrderRepository) {
	OrderSer = &orderService{repo: orderRepository}
}

func (s *orderService) CreateOrder(data *dto.Order) (order *Models.Order, err error) {
	model := &Models.Order{
		CustomerId: data.CustomerId,
		ProductId:  data.ProductId,
		Quantity:   data.Quantity,
	}
	err = s.repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}
	return model, nil
}
func (s *orderService) FindOrder(id int) (order *Models.Order, err error) {
	model := &Models.Order{}
	if err = s.repo.FindOrder(model, id); err != nil {
		return nil, err
	}
	return order, nil
}
