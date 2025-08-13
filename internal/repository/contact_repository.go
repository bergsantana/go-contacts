package repository

import "github.com/bergsantana/go-contacts/internal/entity"

type ContactRepository interface {
	GetAll() ([]entity.Contact, error)
	GetByID(id uint) (*entity.Contact, error)
	Create(contact *entity.Contact) error
	Update(contact *entity.Contact) error
	Delete(id uint) error
}
