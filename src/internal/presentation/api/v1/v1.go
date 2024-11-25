package v1

import (
	useCase "github.com/D1sordxr/aviasales/src/internal/application"
	orderHandler "github.com/D1sordxr/aviasales/src/internal/presentation/api/v1/controllers/handlers/order"
	okHandler "github.com/D1sordxr/aviasales/src/internal/presentation/api/v1/controllers/handlers/statusOK"
	ticketHandler "github.com/D1sordxr/aviasales/src/internal/presentation/api/v1/controllers/handlers/ticket"
	orderRoutes "github.com/D1sordxr/aviasales/src/internal/presentation/api/v1/controllers/routes/order"
	okRoutes "github.com/D1sordxr/aviasales/src/internal/presentation/api/v1/controllers/routes/statusOK"
	ticketRoutes "github.com/D1sordxr/aviasales/src/internal/presentation/api/v1/controllers/routes/ticket"
	"github.com/gin-gonic/gin"
)

type RoutesV1 struct {
	RouterGroup *gin.RouterGroup
	UseCase     *useCase.UseCase
}

func NewRoutesV1(rg *gin.RouterGroup, uc *useCase.UseCase) {
	routes := &RoutesV1{
		RouterGroup: rg,
		UseCase:     uc,
	}
	routes.setupRoutesV1()
}

func (r *RoutesV1) setupRoutesV1() {
	// Main path
	v1 := r.RouterGroup.Group("/v1")

	// Status path
	okHandlers := okHandler.NewOkHandler()
	okRoutes.NewOkRoutes(v1, okHandlers)

	// Tickets path
	ticketHandlers := ticketHandler.NewTicketHandler(r.UseCase)
	ticketRoutes.NewTicketRoutes(v1, ticketHandlers)

	// Orders path
	orderHandlers := orderHandler.NewOrderHandler(r.UseCase)
	orderRoutes.NewOrderRoutes(v1, orderHandlers)
}
