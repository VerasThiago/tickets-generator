package repository

import (
	"github.com/verasthiago/tickets-generator/shared/models"
)

type Repository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) error
	DeleteUser(userID string) error
	UpdateUser(user *models.User) error
	UpdateUserPasswordByEmail(email string, password string) error
	VerifyUserAccountByID(id string) error

	CreateTicket(ticket *models.Ticket) error
	GetTicketListByUserID(userID string) ([]*models.Ticket, error)
	GetTicketByID(id string) (*models.Ticket, error)
	GetTicketByCPF(cpf string) (*models.Ticket, error)
	DeleteTicketByID(id string) error
	ValidateTicketByCPF(cpf string) error
}
