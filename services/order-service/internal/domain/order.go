package domain

import (
	"errors"
	"time"
)

type Order struct {
	id              ID
	createdAt       time.Time
	items           []OrderItem
	status          OrderStatus
	paymentStatus   OrderPaymentStatus
	buyerID         BuyerID
	paymentMethod   PaymentMethod
	deliveryAddress DeliveryAddress
	deliveryCost    Money
	buyerComment    BuyerComment
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

	copyItems := make([]OrderItem, len(items))
	copy(copyItems, items)
	deliveryCost, err := NewMoney(0, CurrencyRUB)
	if err != nil {
		return Order{}, err
	}
	return Order{
		id:              id,
		createdAt:       time.Now(),
		items:           copyItems,
		status:          OrderStatusCreated,
		paymentStatus:   OrderPaymentStatusAwaitingPayment,
		buyerID:         buyerID,
		paymentMethod:   paymentMethod,
		deliveryAddress: deliveryAddress,
		deliveryCost:    deliveryCost,
		buyerComment:    buyerComment,
	}, nil
}
func (o *Order) Pay() error {
	switch {
	case o.status == OrderStatusCancelled:
		return errors.New("cancelled order cannot be paid")
	case o.status == OrderStatusCompleted:
		return errors.New("completed order cannot be completed")
	case o.paymentStatus != OrderPaymentStatusAwaitingPayment:
		return errors.New("order is not awaiting payment")
	}

	o.paymentStatus = OrderPaymentStatusPaid

	return nil
}
func (o *Order) Cancel() error {
	//TODO
}

func (o *Order) Shipped() error {
	//TODO
}
