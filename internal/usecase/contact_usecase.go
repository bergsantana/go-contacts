package usecase

import (
	"errors"
	"fmt"

	"github.com/bergsantana/go-contacts/internal/entity"
	"github.com/bergsantana/go-contacts/internal/repository"
	"github.com/bergsantana/go-contacts/pkg/formatters"
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
	err := cleanAndValidateFields(contact)
	if err != nil {
		fmt.Println("Erro ao atualizar contato: ", err)
		return err
	}
	return uc.repo.Create(contact)
}

func (uc *ContactUsecase) UpdateContact(contact *entity.Contact) error {

	err := cleanAndValidateFields(contact)
	if err != nil {
		fmt.Println("Erro ao atualizar contato: ", err)
		return err
	}

	return uc.repo.Update(contact)
}

func (uc *ContactUsecase) DeleteContact(id uint) error {
	return uc.repo.Delete(id)
}

func cleanAndValidateFields(contact *entity.Contact) error {
	// Valida CPF
	if contact.CPF != nil && *contact.CPF != "" {
		if !validate.IsValidCPF(*contact.CPF) {
			return errors.New("invalid CPF")
		}
	}

	// Validata CNPJ
	if contact.CNPJ != nil && *contact.CNPJ != "" {
		if !validate.IsValidCNPJ(*contact.CNPJ) {
			return errors.New("invalid CNPJ")
		}
	}

	// Formatar telefone
	if contact.Phone != "" {
		formatted, err := formatters.FormatPhoneNumber(contact.Phone)
		if err != nil {
			return err
		}
		contact.Phone = formatted
	}
	return nil
}
