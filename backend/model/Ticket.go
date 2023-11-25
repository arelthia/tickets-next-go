package model

type Ticket struct {
	Id        int    `json:"id" db:"id"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Issue     string `json:"issue" db:"issue"`
	Priority  string `json:"priority" db:"priority"`
	CreatedAt string `json:"createdAt" db:"created_at"`
}

type TicketData interface {
	GetTickets() ([]Ticket, error)
	GetTicket(id int) (Ticket, error)
	CreateTicket(t Ticket) (int64, error)
	UpdateTicket(id int, t Ticket) (int64, error)
	DeleteTicket(id int) (int64, error)
}
