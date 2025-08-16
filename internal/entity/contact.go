package entity

type Contact struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	Name    string  `json:"name"`
	Email   string  `json:"email" gorm:"unique"`
	Phone   string  `json:"phone" gorm:"unique"`
	CPF     *string `json:"cpf,omitempty" gorm:"unique"`  // Nullable
	CNPJ    *string `json:"cnpj,omitempty" gorm:"unique"` // Nullable
	Address *string `json:"address,omitempty"`            // Nullable
}
