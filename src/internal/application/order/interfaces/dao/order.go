package dao

import (
	"github.com/D1sordxr/aviasales/src/internal/application/order/dto"
)

type OrderDAO interface {
	CreateOrderDTO(order dto.Order) (dto.Order, error)
}
