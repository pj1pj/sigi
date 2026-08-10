package models

import "time"

type EstadoOrden string

// Los estados de la orden de compra reflejan su estado en el proceso de compra.
const (
	OrdenPendiente  EstadoOrden = "Pendiente"
	OrdenConfirmada EstadoOrden = "Confirmada"
	OrdenCancelada  EstadoOrden = "Cancelada"
)

// La orden de compra representa una solicitud de productos realizada por la empresa a un proveedor.
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

// AgregarProducto agrega un producto a la orden de compra y recalcula el total.
func (o *OrdenCompra) AgregarProducto(producto *Producto) {
	if producto == nil {
		return
	}

	o.productos = append(o.productos, producto)
	o.CalcularTotal()
}

// El total se calcula a partir de los subtotales de todos los productos.
// De esta forma, el total de la orden no depende de un valor ingresado
// directamente por el usuario.
func (o *OrdenCompra) CalcularTotal() {
	o.total = 0

	for _, producto := range o.productos {
		if producto != nil {
			o.total += producto.Subtotal()
		}
	}
}

// Confirmar cambia el estado de la orden a "Confirmada".
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
