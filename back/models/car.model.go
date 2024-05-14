package models

import "gorm.io/gorm"

type Carro struct {
	gorm.Model

	CarroID       uint   `gorm:"not null;unique_index" json:"carroID"`
	EstadoReserva bool   `gorm:"default:false" json:"estadoReserva"`
	Precio        uint   `gorm:"not null" json:"precio"`
	ModeloYear    uint   `gorm:"not null" json:"modeloYear"`
	ModeloNombre  string `gorm:"not null" json:"modeloNombre"`
	Marca         string `gorm:"not null" json:"marca"`
	Transmision   string `gorm:"not null" json:"transmision"`
	Combustible   string `gorm:"not null" json:"combustible"`
	Tipo          string `gorm:"not null" json:"tipo"`
	Puertas       uint   `gorm:"not null" json:"puertas"`
	ABS           bool   `gorm:"not null" json:"abs"`
}
