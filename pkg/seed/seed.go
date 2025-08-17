package seed

import (
	"log"

	"github.com/bergsantana/go-contacts/internal/entity"
	"github.com/bergsantana/go-contacts/internal/repository"
	"github.com/bergsantana/go-contacts/internal/usecase"
	"gorm.io/gorm"
)

func SeedContacts(db *gorm.DB) {
	repo := repository.NewContactGormRepository(db)
	uc := usecase.NewContactUsecase(repo)

	contacts := []entity.Contact{
		{
			Name:    "João Pereira",
			Email:   "joao.pereira@example.com",
			CPF:     ptrString("532.218.790-16"),
			CNPJ:    nil,
			Phone:   "11987654321",
			Address: ptrString("Rua das Palmeiras, 123, São Paulo - SP"),
		},
		{
			Name:    "Tech Solutions Ltda",
			Email:   "contato@techsolutions.com.br",
			CPF:     nil,
			CNPJ:    ptrString("70.958.874/0001-31"),
			Phone:   "11998567890",
			Address: ptrString("Avenida Paulista, 1500, São Paulo - SP"),
		},
		{
			Name:    "Maria Oliveira",
			Email:   "maria.oliveira@example.com",
			CPF:     nil,
			CNPJ:    nil,
			Phone:   "21987654321",
			Address: ptrString("Rua das Acácias, 45, Rio de Janeiro - RJ"),
		},
		{
			Name:    "Consultoria ABC",
			Email:   "abc.consultoria@example.com",
			CPF:     ptrString("904.278.000-21"),
			CNPJ:    ptrString("26.277.322/0001-76"),
			Phone:   "31987654321",
			Address: ptrString("Av. Central, 250, Belo Horizonte - MG"),
		},
		{
			Name:    "Carlos Souza",
			Email:   "carlos.souza@example.com",
			CPF:     ptrString("221.564.340-42"),
			CNPJ:    nil,
			Phone:   "11976543210",
			Address: ptrString("Rua Azul, 78, Campinas - SP"),
		},
	}

	for _, contact := range contacts {

		if err := uc.CreateContact(&contact); err != nil {
			log.Printf("Falha ao conectar ao banco de dados %s: %v\n", contact.Email, err)
		} else {
			log.Printf("Contato inserido com sucesso ->  %s.\n", contact.Email)
		}
	}
}

func ptrString(s string) *string {
	return &s
}
