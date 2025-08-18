package usecase

import (
	"errors"
	"fmt"

	"github.com/bergsantana/go-contacts/internal/entity"
	"github.com/bergsantana/go-contacts/internal/repository"
	"github.com/bergsantana/go-contacts/pkg/formatters"
	"github.com/bergsantana/go-contacts/pkg/sanitize"
	"github.com/bergsantana/go-contacts/pkg/validate"
	"github.com/gofiber/fiber/v2"
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

func (uc *ContactUsecase) GetByCPF(cpf string) (*entity.Contact, error) {
	if cpf == "" {
		return nil, errors.New("cpf cannot be empty")
	}
	return uc.repo.GetByCPF(cpf)
}

func (uc *ContactUsecase) GetByCNPJ(cnpj string) (*entity.Contact, error) {
	if cnpj == "" {
		return nil, errors.New("cnpj cannot be empty")
	}
	return uc.repo.GetByCNPJ(cnpj)
}

func (uc *ContactUsecase) CreateContact(contact *entity.Contact, c *fiber.Ctx) error {
	err := cleanAndValidateFields(contact, uc, c)
	if err != nil {
		fmt.Println("Erro ao criar contato: ", err)
		return err
	}
	return uc.repo.Create(contact)
}

func (uc *ContactUsecase) UpdateContact(contact *entity.Contact, c *fiber.Ctx) error {

	err := cleanAndValidateFields(contact, uc, c)
	if err != nil {
		fmt.Println("Erro ao atualizar contato: ", err)
		return err
	}

	return uc.repo.Update(contact)
}

func (uc *ContactUsecase) DeleteContact(id uint) error {
	return uc.repo.Delete(id)
}

func cleanAndValidateFields(contact *entity.Contact, uc *ContactUsecase, c *fiber.Ctx) error {
	// Sanitização de XSS
	contact.Name = sanitize.StrictHTML(contact.Name)
	contact.Email = sanitize.StrictHTML(contact.Email)
	contact.Phone = sanitize.StrictHTML(contact.Phone)

	// Sanitizar de SQL Injection
	contact.Name = sanitize.SanitizeSQLInput(contact.Name)
	contact.Email = sanitize.SanitizeSQLInput(contact.Email)
	if contact.Address != nil && *contact.Address != "" {
		*contact.Address = sanitize.SanitizeSQLInput(*contact.Address)
	}

	// Valida CPF
	if contact.CPF != nil && *contact.CPF != "" {
		if !validate.IsValidCPF(*contact.CPF, c.UserContext()) {
			return errors.New("CPF invalido: " + *contact.CPF)
		}
		existingCPF, _ := uc.repo.GetByCPF(*contact.CPF)
		if existingCPF != nil {
			return errors.New("CPF já utilizado")
		}
	}

	// Validata CNPJ
	if contact.CNPJ != nil && *contact.CNPJ != "" {
		if !validate.IsValidCNPJ(*contact.CNPJ, c.UserContext()) {
			return errors.New("CNPJ INVALIDO: " + *contact.CNPJ)
		}
		existingCNPJ, _ := uc.repo.GetByCNPJ(*contact.CNPJ)
		if existingCNPJ != nil {
			return errors.New("CNPJ já utilizado")
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

	// Certifique que o email é único
	existing, _ := uc.repo.GetByEmail(contact.Email)
	if existing != nil {
		return errors.New("email já utilizado")
	}

	return nil
}
