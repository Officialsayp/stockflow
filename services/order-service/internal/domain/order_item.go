package domain

type OrderItem struct {
	SKU              SKU
	Quantity         int64
	UnitPrice        Money
	Discount         Money
	ReservedQuantity int64
}
