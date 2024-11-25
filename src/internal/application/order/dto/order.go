package dto

import "github.com/google/uuid"

type Order struct {
	ID            int       `json:"id"`
	ClientID      uuid.UUID `json:"client_id"`
	AddressID     uuid.UUID `json:"address_id"`
	OrderStatus   string    `json:"order_status"`
	PaymentMethod string    `json:"payment_method"`
	SerialNumber  int       `json:"serial_number"`
	Closed        bool      `json:"closed"`
	Deleted       bool      `json:"deleted"`
	TotalPrice    int       `json:"total_price"`
}

type Orders struct {
	Orders []Order
}
