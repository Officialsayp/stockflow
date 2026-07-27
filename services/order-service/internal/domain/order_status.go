package domain

type OrderStatus string

const (
	OrderStatusCreated    OrderStatus = "created"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type OrderPaymentStatus string

const (
	OrderPaymentStatusAwaitingPayment OrderPaymentStatus = "awaiting_payment"
	OrderPaymentStatusProcessing      OrderPaymentStatus = "processing"
	OrderPaymentStatusPaid            OrderPaymentStatus = "paid"
	OrderPaymentStatusRefunded        OrderPaymentStatus = "refunded"
	OrderPaymentStatusFailed          OrderPaymentStatus = "failed"
)
