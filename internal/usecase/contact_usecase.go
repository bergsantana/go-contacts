package usecase

import (
	"errors"

	"github.com/bergsantana/go-contacts/internal/entity"
	"github.com/bergsantana/go-contacts/internal/repository"
	"github.com/bergsantana/go-contacts/pkg/validate"
)

type ContactUsecase struct {
	repo repository.ContactRepository
}

func NewContactUsecase(repo repository.ContactRepository) *ContactUsecase {
	return &ContactUsecase{repo: repo}
}

func (uc *ContactUsecase) GetContacts() ([]entity.Contact, error) {
	return uc.repo.GetAll()
}

func (uc *ContactUsecase) GetContactByID(id uint) (*entity.Contact, error) {
	return uc.repo.GetByID(id)
}

func (uc *ContactUsecase) CreateContact(contact *entity.Contact) error {
	if contact.CPF != nil && *contact.CPF != "" {
		if !validate.IsValidCPF(*contact.CPF) {
			return errors.New("invalid CPF")
		}
	}
	return uc.repo.Create(contact)
}

func (uc *ContactUsecase) UpdateContact(contact *entity.Contact) error {
	if contact.CPF != nil && *contact.CPF != "" {
		if !validate.IsValidCPF(*contact.CPF) {
			return errors.New("invalid CPF")
		}
	}
	return uc.repo.Update(contact)
}

func (uc *ContactUsecase) DeleteContact(id uint) error {
	return uc.repo.Delete(id)
}
