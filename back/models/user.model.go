package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model

	Email         string    `gorm:"not null;unique_index" json:"email"`
	NombreUsuario string    `gorm:"not null;type:varchar(32)" json:"nombreUsuario"`
	Password      string    `gorm:"not null" json:"password"`
	Reservas      []Reserva `json:"reservas"`
}
