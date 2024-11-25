package dao

import (
	"context"
	"errors"
	"github.com/D1sordxr/aviasales/src/internal/application/order/dto"
	"github.com/D1sordxr/aviasales/src/internal/db/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"time"
)

type OrderDAO struct {
	DB *pgx.Conn
}

func NewOrderDAO(conn *pgx.Conn) *OrderDAO {
	return &OrderDAO{DB: conn}
}

func (dao *OrderDAO) CreateOrder(order models.Order) (int, error) {
	isAvailable, err := dao.ticketsAvailabilityCheck(order.Tickets)
	if err != nil {
		return 0, err
	}
	if !isAvailable {
		return 0, errors.New("ticket is not available")
	}

	clientID := uuid.New()
	addressID := uuid.New()
	serialNumber := func() int {
		now := time.Now()
		return int(now.UnixNano() % 1000000000)
	}()

	tx, err := dao.DB.Begin(context.Background())
	if err != nil {
		return 0, err
	}

	var orderID int
	err = tx.QueryRow(context.Background(), `
		INSERT INTO orders (
                client_id, 
                address_id,
                order_status,
                payment_method,
                serial_number,
                closed,
                deleted,
                created_at,
                updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
        RETURNING id
	`, clientID, addressID, "Created", order.PaymentMethod, serialNumber, false, false).Scan(&orderID)

	if err != nil {
		_ = tx.Rollback(context.Background())
		return 0, err
	}

	for _, v := range order.Tickets {
		_, err = tx.Exec(context.Background(), `
			UPDATE tickets 
			SET 
			    order_id = $1, 
				passenger_name = $2,
				is_available = $3
			WHERE id = $4 
		`, orderID, clientID, false, v.ID)
		if err != nil {
			_ = tx.Rollback(context.Background())
			return 0, err
		}
	}

	if err = tx.Commit(context.Background()); err != nil {
		return 0, err
	}

	return orderID, nil
}

func (dao *OrderDAO) CreateOrderDTO(order dto.Order) (dto.Order, error) {
	return dto.Order{OrderStatus: order.OrderStatus}, nil
}

func (dao *OrderDAO) ticketsAvailabilityCheck(t []models.Ticket) (bool, error) {
	var id []*uint64

	for _, v := range t {
		id = append(id, v.ID)
	}

	rows, err := dao.DB.Query(context.Background(), ` 
		SELECT is_available FROM tickets WHERE ID = ANY($1::INTEGER[])
	`, id)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		var isAvailable bool

		err = rows.Scan(
			&isAvailable,
		)

		if err != nil {
			return false, err
		}

		if isAvailable == false {
			return false, nil
		}
	}

	return true, nil
}

func (dao *OrderDAO) GetOrder(id int) (models.Order, error) {
	var order models.Order

	rows, err := dao.DB.Query(context.Background(), `
	SELECT 
		o.id, 
		o.client_id, 
		o.address_id, 
		o.order_status, 
		o.payment_method, 
		o.serial_number, 
		o.closed, 
		o.deleted,
		t.id,
		t.order_id,
		t.passenger_name,
		t.destination,
		t.payment,
		t.is_available,
		t.dispatch_time,
		t.arrival_time
		FROM orders o
		LEFT JOIN tickets t ON o.id = t.order_id 
		WHERE o.id = $1
	`, id)
	if err != nil {
		return models.Order{}, err
	}
	defer rows.Close()

	var tickets []models.Ticket

	for rows.Next() {
		var ticket models.Ticket

		err = rows.Scan(
			&order.ID,
			&order.ClientID,
			&order.AddressID,
			&order.OrderStatus,
			&order.PaymentMethod,
			&order.SerialNumber,
			&order.Closed,
			&order.Deleted,
			&ticket.ID,
			&ticket.OrderID,
			&ticket.PassengerName,
			&ticket.Destination,
			&ticket.Payment,
			&ticket.IsAvailable,
			&ticket.DispatchTime,
			&ticket.ArrivalTime,
		)

		if err != nil {
			return models.Order{}, err
		}
		tickets = append(tickets, ticket)
	}
	order.Tickets = tickets

	return order, nil
}
