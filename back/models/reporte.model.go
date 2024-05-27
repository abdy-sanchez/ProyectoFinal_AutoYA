package models

type Reporte struct { //Este modelo NO se guarda en DB

	PrecioTotal int64   `json:"precioTotal"`
	Detalles    []Carro `json:"detalles"`
}
