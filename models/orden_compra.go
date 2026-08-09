package models

import "time"

type EstadoOrden string

const (
	OrdenPendiente  EstadoOrden = "Pendiente"
	OrdenConfirmada EstadoOrden = "Confirmada"
	OrdenCancelada  EstadoOrden = "Cancelada"
)

type OrdenCompra struct {
	codigo        string
	proveedor     *Proveedor
	productos     []*Producto
	total         float64
	fechaCreacion time.Time
	estado        EstadoOrden
}

func NuevaOrdenCompra(codigo string, proveedor *Proveedor) *OrdenCompra {
	return &OrdenCompra{
		codigo:        codigo,
		proveedor:     proveedor,
		productos:     make([]*Producto, 0),
		total:         0,
		fechaCreacion: time.Now(),
		estado:        OrdenPendiente,
	}
}

func (o *OrdenCompra) Codigo() string {
	return o.codigo
}

func (o *OrdenCompra) Proveedor() *Proveedor {
	return o.proveedor
}

func (o *OrdenCompra) Productos() []*Producto {
	return o.productos
}

func (o *OrdenCompra) Total() float64 {
	return o.total
}

func (o *OrdenCompra) FechaCreacion() time.Time {
	return o.fechaCreacion
}

func (o *OrdenCompra) Estado() EstadoOrden {
	return o.estado
}

func (o *OrdenCompra) AgregarProducto(producto *Producto) {
	if producto == nil {
		return
	}

	o.productos = append(o.productos, producto)
	o.CalcularTotal()
}

func (o *OrdenCompra) CalcularTotal() {
	o.total = 0

	for _, producto := range o.productos {
		if producto != nil {
			o.total += producto.Subtotal()
		}
	}
}

func (o *OrdenCompra) Confirmar() {
	o.estado = OrdenConfirmada
}

func (o *OrdenCompra) Cancelar() {
	o.estado = OrdenCancelada
}

func (o *OrdenCompra) EstaConfirmada() bool {
	return o.estado == OrdenConfirmada
}

func (o *OrdenCompra) EstaCancelada() bool {
	return o.estado == OrdenCancelada
}
