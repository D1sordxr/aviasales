package order

import (
	"github.com/D1sordxr/aviasales/src/internal/application/order/dto"
	"github.com/D1sordxr/aviasales/src/internal/presentation/api/v1/controllers/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase
}

type UseCase interface {
	CreateOrderDTO(order dto.Order) (dto.Order, error)
}

func NewOrderHandler(useCase UseCase) *Handler {
	return &Handler{UseCase: useCase}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	c.JSON(418, response.CommonResponse{
		Message: "no message",
		Data:    nil,
	})
}
