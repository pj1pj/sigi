package models

type EstadoProveedor string

const (
	ProveedorActivo   EstadoProveedor = "Activo"
	ProveedorInactivo EstadoProveedor = "Inactivo"
)

type Proveedor struct {
	codigo   string
	empresa  string
	pais     string
	contacto string
	telefono string
	correo   string
	estado   EstadoProveedor
}

func NuevoProveedor(codigo, empresa, pais, contacto, telefono, correo string) *Proveedor {
	return &Proveedor{
		codigo:   codigo,
		empresa:  empresa,
		pais:     pais,
		contacto: contacto,
		telefono: telefono,
		correo:   correo,
		estado:   ProveedorActivo,
	}
}

func (p *Proveedor) Codigo() string {
	return p.codigo
}

func (p *Proveedor) Empresa() string {
	return p.empresa
}

func (p *Proveedor) Pais() string {
	return p.pais
}

func (p *Proveedor) Contacto() string {
	return p.contacto
}

func (p *Proveedor) Telefono() string {
	return p.telefono
}

func (p *Proveedor) Correo() string {
	return p.correo
}

func (p *Proveedor) Estado() EstadoProveedor {
	return p.estado
}

// El estado se modifica mediante métodos para mantener encapsulado
// el valor interno del proveedor y evitar cambios directos desde otras capas.
func (p *Proveedor) Activar() {
	p.estado = ProveedorActivo
}

func (p *Proveedor) Desactivar() {
	p.estado = ProveedorInactivo
}

func (p *Proveedor) EstaActivo() bool {
	return p.estado == ProveedorActivo
}

func (p *Proveedor) ActualizarContacto(contacto, telefono, correo string) {
	p.contacto = contacto
	p.telefono = telefono
	p.correo = correo
}
