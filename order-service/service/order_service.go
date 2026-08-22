package service

import "errors"

type OrderService struct {}

func (o *OrderService) CreateOrder(product string)error{
	if product == "unavailable"{
		return errors.New("product is unavailable")
	}
	return nil
}