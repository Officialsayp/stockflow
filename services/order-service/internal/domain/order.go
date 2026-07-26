package domain

import (
	"errors"
	"time"
)

type Order struct {
	ID              ID
	CreatedAt       time.Time
	Items           []OrderItem
	Status          OrderStatus
	PaymentStatus   OrderPaymentStatus
	BuyerID         BuyerID
	PaymentMethod   PaymentMethod
	DeliveryAddress DeliveryAddress
	DeliveryCost    Money
	BuyerComment    BuyerComment
}

func NewOrder(id ID, items []OrderItem, buyerID BuyerID, paymentMethod PaymentMethod, deliveryAddress DeliveryAddress, buyerComment BuyerComment) (Order, error) {
	switch {
	case id == "":
		return Order{}, errors.New("id is empty")
	case len(items) == 0:
		return Order{}, errors.New("items is empty")
	case buyerID == "":
		return Order{}, errors.New("buyer id is empty")
	case paymentMethod == "":
		return Order{}, errors.New("payment method is empty")
	case deliveryAddress == "":
		return Order{}, errors.New("delivery address is empty")
	}

	return Order{}, nil
}
