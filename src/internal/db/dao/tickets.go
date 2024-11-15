package dao

import (
	"aviasales/src/internal/db/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"time"
)

type TicketDAO struct {
	DB *pgx.Conn
}

func NewTicketDAO(conn *pgx.Conn) *TicketDAO {
	return &TicketDAO{DB: conn}
}

func (dao *TicketDAO) GetTickets() ([]models.Ticket, error) {
	var ticket models.Ticket
	rows, err := dao.DB.Query(context.Background(), `
	SELECT id, passenger_name, destination, payment, dispatch_time, arrival_time FROM tickets
	 `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := make([]models.Ticket, 0, 10)
	for rows.Next() {
		if err = rows.Scan(
			&ticket.ID, &ticket.PassengerName, &ticket.Destination,
			&ticket.Payment, &ticket.DispatchTime, &ticket.ArrivalTime,
		); err != nil {
			return nil, err
		}
		data = append(data, ticket)
	}

	return data, nil
}

func (dao *TicketDAO) CreateTicket(t models.Ticket) error {
	_, err := dao.DB.Exec(context.Background(), `
		INSERT INTO tickets(passenger_name, destination, payment, dispatch_time, arrival_time, created_at) 
		VALUES ($1, $2, $3, $4, $5, NOW())`,
		t.PassengerName, t.Destination, t.Payment, t.DispatchTime, t.ArrivalTime,
	)
	if err != nil {
		return err
	}

	return nil
}

func (dao *TicketDAO) UpdateTicket(t models.Ticket) error {
	query :=
		`UPDATE tickets
	SET passenger_name = COALESCE(NULLIF($1::TEXT, ''), passenger_name),
	destination = COALESCE(NULLIF($2::TEXT, ''), destination),
	payment = COALESCE($3::INTEGER, payment),
	dispatch_time = COALESCE($4::TIMESTAMP, dispatch_time),
	arrival_time = COALESCE($5::TIMESTAMP, arrival_time),
	updated_at = NOW()
	WHERE id = $6::INTEGER`

	passengerName := ""
	if t.PassengerName != nil {
		passengerName = *t.PassengerName
	}

	destination := ""
	if t.Destination != nil {
		destination = *t.Destination
	}

	var payment *uint64
	if t.Payment != nil {
		payment = t.Payment
	}

	var dispatchTime *time.Time
	if t.DispatchTime != nil {
		dispatchTime = t.DispatchTime
	}

	var arrivalTime *time.Time
	if t.ArrivalTime != nil {
		arrivalTime = t.ArrivalTime
	}

	_, err := dao.DB.Exec(context.Background(), query,
		passengerName, destination, payment,
		dispatchTime, arrivalTime, t.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update ticket: %w", err)
	}

	return nil
}

func (dao *TicketDAO) DeleteTicket(id string) error {
	_, err := dao.DB.Exec(context.Background(), `
		DELETE FROM tickets WHERE ID = $1
	`, id)
	if err != nil {
		return err
	}

	return nil
}
