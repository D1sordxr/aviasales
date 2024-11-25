package application

import (
	orderDTO "github.com/D1sordxr/aviasales/src/internal/application/order/dto"
	order "github.com/D1sordxr/aviasales/src/internal/application/order/interfaces/dao"
	ticketDTO "github.com/D1sordxr/aviasales/src/internal/application/ticket/dto"
	ticket "github.com/D1sordxr/aviasales/src/internal/application/ticket/interfaces/dao"
)

type UseCase struct {
	TicketDAO ticket.TicketDAO
	OrderDAO  order.OrderDAO
}

func NewUseCase(ticketDAO ticket.TicketDAO, orderDAO order.OrderDAO) *UseCase {
	return &UseCase{
		TicketDAO: ticketDAO,
		OrderDAO:  orderDAO,
	}
}

func (uc *UseCase) GetTicketsDTO() (ticketDTO.Tickets, error) {
	return uc.TicketDAO.GetTicketsDTO()
}

func (uc *UseCase) GetTicketByIDDTO(id string) (ticketDTO.Ticket, error) {
	return uc.TicketDAO.GetTicketByIDDTO(id)
}

func (uc *UseCase) CreateTicketDTO(t ticketDTO.Ticket) (ticketDTO.Ticket, error) {
	return uc.TicketDAO.CreateTicketDTO(t)
}

func (uc *UseCase) UpdateTicketDTO(t ticketDTO.Ticket) (ticketDTO.Ticket, error) {
	return uc.TicketDAO.UpdateTicketDTO(t)
}

func (uc *UseCase) DeleteTicketDTO(id string) (ticketDTO.Ticket, error) {
	return uc.TicketDAO.DeleteTicketDTO(id)
}

func (uc *UseCase) CreateOrderDTO(order orderDTO.Order) (orderDTO.Order, error) {
	return uc.OrderDAO.CreateOrderDTO(order)
}
