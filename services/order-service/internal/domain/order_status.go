package domain

type OrderStatus string

const (
	OrderStatusCreated      OrderStatus = "created"
	OrderStatusProcessing   OrderStatus = "processing"
	OrderStatusShipped      OrderStatus = "shipped"
	OrderStatusDelivered    OrderStatus = "delivered"
	OrderStatusCompleted    OrderStatus = "completed"
	OrderStatusCancellation OrderStatus = "cancellation"
	OrderStatusCancelled    OrderStatus = "cancelled"
)

type OrderPaymentStatus string

const (
	OrderPaymentStatusAwaitingPayment  OrderPaymentStatus = "awaiting_payment"
	OrderPaymentStatusProcessing       OrderPaymentStatus = "processing"
	OrderPaymentStatusPaid             OrderPaymentStatus = "paid"
	OrderPaymentStatusRefundProcessing OrderPaymentStatus = "refund_processing"
	OrderPaymentStatusRefunded         OrderPaymentStatus = "refunded"
	OrderPaymentStatusFailed           OrderPaymentStatus = "failed"
	OrderPaymentStatusCancelled        OrderPaymentStatus = "cancelled"
)
