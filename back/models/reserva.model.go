package models

import "gorm.io/gorm"

type Reserva struct {
	gorm.Model

	ReservaID  uint `gorm:"not null;unique_index" json:"reservaID"`
	Usuario_ID uint `gorm:"not null;" json:"usuario_ID"`
	Carro_ID   uint `gorm:"not null" json:"carro_ID"`
}
