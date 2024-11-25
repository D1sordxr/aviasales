package order

import (
	"github.com/D1sordxr/aviasales/src/internal/presentation/api/v1/controllers/handlers/order"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	RouterGroup *gin.RouterGroup
	Handler     *order.Handler
}

func NewOrderRoutes(rg *gin.RouterGroup, h *order.Handler) {
	routes := Routes{
		RouterGroup: rg,
		Handler:     h,
	}
	routes.setupOrderRoutes()
}

func (r *Routes) setupOrderRoutes() {
	tickets := r.RouterGroup.Group("/orders")
	{
		tickets.POST("/order", r.Handler.CreateOrder)
		// TODO: get orders
		// TODO: get order by id
		// TODO: update order
	}
}
