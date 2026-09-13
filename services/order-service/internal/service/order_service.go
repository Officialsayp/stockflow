package service

import "errors"

var ErrProductUnavailable = errors.New("product is unavailable")

type OrderService struct{}

func (o *OrderService) CreateOrder(product string) error {
	if product == "unavailable" {
		return ErrProductUnavailable
	}
	return nil
}
