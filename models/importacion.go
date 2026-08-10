package models

import "time"

type EstadoImportacion string

// Los estados de la importación reflejan el progreso del proceso de importación.
const (
	ImportacionEnPreparacion EstadoImportacion = "En preparación"
	ImportacionEnTransito    EstadoImportacion = "En tránsito"
	ImportacionEnAduana      EstadoImportacion = "En aduana"
	ImportacionLlegadaBodega EstadoImportacion = "Llegó a bodega"
)

// La importación representa el proceso de traer productos desde un proveedor extranjero hasta la bodega de la empresa.
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

// NuevaImportacion crea una nueva instancia de Importacion con los datos proporcionados.
func NuevaImportacion(
	codigo string,
	orden *OrdenCompra,
	transporte *Transporte,
	ciudadOrigen string,
	ciudadDestino string,
) *Importacion {
	// La fecha de registro se establece automáticamente al momento de crear la importación.
	fechaRegistro := time.Now()

	return &Importacion{
		codigo:        codigo,
		orden:         orden,
		transporte:    transporte,
		ciudadOrigen:  ciudadOrigen,
		ciudadDestino: ciudadDestino,
		fechaRegistro: fechaRegistro,

		

		// la fecha de registro y agregando 30 días.
		fechaEstimadaLlegada: fechaRegistro.AddDate(0, 0, 30),

		\
		estado: ImportacionEnTransito,
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

// El estado se modifica mediante este método para mantener
// encapsulado el estado interno de la importación.
// La validación de si el cambio es permitido corresponderá
// posteriormente a la capa de servicios.
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
