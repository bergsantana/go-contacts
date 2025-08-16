package repository

import "github.com/bergsantana/go-contacts/internal/entity"

type ContactRepository interface {
	GetAll() ([]entity.Contact, error)
	GetByID(id uint) (*entity.Contact, error)
	GetByCPF(cpf string) (*entity.Contact, error)
	GetByCNPJ(cnpj string) (*entity.Contact, error)
	GetByEmail(email string) (*entity.Contact, error)
	Create(contact *entity.Contact) error
	Update(contact *entity.Contact) error
	Delete(id uint) error
}
