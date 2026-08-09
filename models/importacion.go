package models

import "time"

type EstadoImportacion string

const (
	ImportacionEnPreparacion EstadoImportacion = "En preparación"
	ImportacionEnTransito    EstadoImportacion = "En tránsito"
	ImportacionEnAduana      EstadoImportacion = "En aduana"
	ImportacionLlegadaBodega EstadoImportacion = "Llegó a bodega"
)

type Importacion struct {
	codigo               string
	orden                *OrdenCompra
	transporte           *Transporte
	ciudadOrigen         string
	ciudadDestino        string
	fechaRegistro        time.Time
	fechaEstimadaLlegada time.Time
	estado               EstadoImportacion
}

func NuevaImportacion(
	codigo string,
	orden *OrdenCompra,
	transporte *Transporte,
	ciudadOrigen string,
	ciudadDestino string,
) *Importacion {
	fechaRegistro := time.Now()

	return &Importacion{
		codigo:               codigo,
		orden:                orden,
		transporte:           transporte,
		ciudadOrigen:         ciudadOrigen,
		ciudadDestino:        ciudadDestino,
		fechaRegistro:        fechaRegistro,
		fechaEstimadaLlegada: fechaRegistro.AddDate(0, 0, 30),
		estado:               ImportacionEnTransito,
	}
}

func (i *Importacion) Codigo() string {
	return i.codigo
}

func (i *Importacion) Orden() *OrdenCompra {
	return i.orden
}

func (i *Importacion) Transporte() *Transporte {
	return i.transporte
}

func (i *Importacion) CiudadOrigen() string {
	return i.ciudadOrigen
}

func (i *Importacion) CiudadDestino() string {
	return i.ciudadDestino
}

func (i *Importacion) FechaRegistro() time.Time {
	return i.fechaRegistro
}

func (i *Importacion) FechaEstimadaLlegada() time.Time {
	return i.fechaEstimadaLlegada
}

func (i *Importacion) Estado() EstadoImportacion {
	return i.estado
}

func (i *Importacion) ActualizarEstado(nuevoEstado EstadoImportacion) {
	i.estado = nuevoEstado
}

func (i *Importacion) EstaEnPreparacion() bool {
	return i.estado == ImportacionEnPreparacion
}

func (i *Importacion) EstaEnTransito() bool {
	return i.estado == ImportacionEnTransito
}

func (i *Importacion) EstaEnAduana() bool {
	return i.estado == ImportacionEnAduana
}

func (i *Importacion) LlegoABodega() bool {
	return i.estado == ImportacionLlegadaBodega
}
