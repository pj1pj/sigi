package models

type EstadoInventario string

// Los estados del inventario reflejan la disponibilidad de los productos en bodega.
const (
	InventarioDisponible EstadoInventario = "Disponible"
	InventarioAgotado    EstadoInventario = "Agotado"
)

type Inventario struct {
	producto    *Producto
	cantidad    int
	orden       *OrdenCompra
	proveedor   *Proveedor
	importacion *Importacion
	ubicacion   string
	estado      EstadoInventario
}

func NuevoInventario(
	producto *Producto,
	cantidad int,
	orden *OrdenCompra,
	proveedor *Proveedor,
	importacion *Importacion,
	ubicacion string,
	// El estado del inventario se determina automáticamente en función de la cantidad.
) *Inventario {
	estado := InventarioDisponible

	if cantidad <= 0 {
		estado = InventarioAgotado
	}

	return &Inventario{
		producto:    producto,
		cantidad:    cantidad,
		orden:       orden,
		proveedor:   proveedor,
		importacion: importacion,
		ubicacion:   ubicacion,
		estado:      estado,
	}
}

func (i *Inventario) Producto() *Producto {
	return i.producto
}

func (i *Inventario) Cantidad() int {
	return i.cantidad
}

func (i *Inventario) Orden() *OrdenCompra {
	return i.orden
}

func (i *Inventario) Proveedor() *Proveedor {
	return i.proveedor
}

func (i *Inventario) Importacion() *Importacion {
	return i.importacion
}

func (i *Inventario) Ubicacion() string {
	return i.ubicacion
}

func (i *Inventario) Estado() EstadoInventario {
	return i.estado
}

// AgregarCantidad aumenta la cantidad disponible en el inventario.
func (i *Inventario) AgregarCantidad(cantidad int) {
	if cantidad <= 0 {
		return
	}

	i.cantidad += cantidad
	i.actualizarEstado()
}

// RetirarCantidad disminuye la cantidad disponible en el inventario.
func (i *Inventario) RetirarCantidad(cantidad int) bool {
	if cantidad <= 0 || cantidad > i.cantidad {
		return false
	}

	i.cantidad -= cantidad
	i.actualizarEstado()

	return true
}

// El estado del inventario depende de la cantidad disponible.
// La actualización se mantiene dentro de la entidad para
// evitar que otras capas tengan que modificar el estado directamente.
func (i *Inventario) actualizarEstado() {
	if i.cantidad > 0 {
		i.estado = InventarioDisponible
		return
	}

	i.estado = InventarioAgotado
}
