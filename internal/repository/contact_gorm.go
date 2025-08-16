package repository

import (
	"github.com/bergsantana/go-contacts/internal/entity"
	"gorm.io/gorm"
)

type contactGormRepository struct {
	db *gorm.DB
}

func NewContactGormRepository(db *gorm.DB) ContactRepository {
	return &contactGormRepository{db: db}
}

func (r *contactGormRepository) GetAll() ([]entity.Contact, error) {
	var contacts []entity.Contact
	err := r.db.Find(&contacts).Error
	return contacts, err
}

func (r *contactGormRepository) GetByID(id uint) (*entity.Contact, error) {
	var contact entity.Contact
	if err := r.db.First(&contact, id).Error; err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *contactGormRepository) Create(contact *entity.Contact) error {
	return r.db.Create(contact).Error
}

func (r *contactGormRepository) Update(contact *entity.Contact) error {
	return r.db.Save(contact).Error
}

func (r *contactGormRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Contact{}, id).Error
}

func (r *contactGormRepository) GetByCPF(cpf string) (*entity.Contact, error) {
	var contact entity.Contact
	result := r.db.Where("cpf = ?", cpf).First(&contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return &contact, nil
}

func (r *contactGormRepository) GetByCNPJ(cnpj string) (*entity.Contact, error) {
	var contact entity.Contact
	result := r.db.Where("cnpj = ?", cnpj).First(&contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return &contact, nil
}

func (r *contactGormRepository) GetByEmail(email string) (*entity.Contact, error) {
	var contact entity.Contact
	result := r.db.Where("email = ?", email).First(&contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return &contact, nil
}
