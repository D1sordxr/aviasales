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
	var order dto.Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(400, response.CommonResponse{Message: "error", Data: err.Error()})
	}

	order, err = h.UseCase.CreateOrderDTO(order)
	if err != nil {
		c.JSON(400, response.CommonResponse{Message: "error", Data: err.Error()})
	}

	c.JSON(200, response.CommonResponse{
		Message: "Successfully created!",
		Data:    order,
	})
}
