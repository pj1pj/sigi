package models

type Producto struct {
	nombre         string
	cantidad       int
	precioUnitario float64
	subtotal       float64
}

func NuevoProducto(nombre string, cantidad int, precioUnitario float64) *Producto {
	return &Producto{
		nombre:         nombre,
		cantidad:       cantidad,
		precioUnitario: precioUnitario,
		subtotal:       float64(cantidad) * precioUnitario,
	}
}

func (p *Producto) Nombre() string {
	return p.nombre
}

func (p *Producto) Cantidad() int {
	return p.cantidad
}

func (p *Producto) PrecioUnitario() float64 {
	return p.precioUnitario
}

func (p *Producto) Subtotal() float64 {
	return p.subtotal
}

func (p *Producto) ActualizarCantidad(cantidad int) {
	p.cantidad = cantidad
	p.calcularSubtotal()
}

func (p *Producto) ActualizarPrecio(precioUnitario float64) {
	p.precioUnitario = precioUnitario
	p.calcularSubtotal()
}

func (p *Producto) calcularSubtotal() {
	p.subtotal = float64(p.cantidad) * p.precioUnitario
}
