package models

// La entidad Transporte representa un medio de transporte utilizado para la importación de productos desde el proveedor hasta la empresa.
type TipoTransporte string

const (
	TransporteMaritimo  TipoTransporte = "Marítimo"
	TransporteAereo     TipoTransporte = "Aéreo"
	TransporteTerrestre TipoTransporte = "Terrestre"
)

type EstadoTransporte string

const (
	TransporteActivo   EstadoTransporte = "Activo"
	TransporteInactivo EstadoTransporte = "Inactivo"
)

type Transporte struct {
	codigo   string
	tipo     TipoTransporte
	empresa  string
	pais     string
	contacto string
	telefono string
	correo   string
	estado   EstadoTransporte
}

// NuevaTransporte crea una nueva instancia de Transporte con los datos proporcionados.
func NuevoTransporte(
	codigo string,
	tipo TipoTransporte,
	empresa string,
	pais string,
	contacto string,
	telefono string,
	correo string,
) *Transporte {
	return &Transporte{
		codigo:   codigo,
		tipo:     tipo,
		empresa:  empresa,
		pais:     pais,
		contacto: contacto,
		telefono: telefono,
		correo:   correo,
		estado:   TransporteActivo,
	}
}

func (t *Transporte) Codigo() string {
	return t.codigo
}

func (t *Transporte) Tipo() TipoTransporte {
	return t.tipo
}

func (t *Transporte) Empresa() string {
	return t.empresa
}

func (t *Transporte) Pais() string {
	return t.pais
}

func (t *Transporte) Contacto() string {
	return t.contacto
}

func (t *Transporte) Telefono() string {
	return t.telefono
}

func (t *Transporte) Correo() string {
	return t.correo
}

func (t *Transporte) Estado() EstadoTransporte {
	return t.estado
}

// Los métodos Activar y Desactivar permiten cambiar el estado del transporte entre activo e inactivo.
func (t *Transporte) Activar() {
	t.estado = TransporteActivo
}

func (t *Transporte) Desactivar() {
	t.estado = TransporteInactivo
}

func (t *Transporte) EstaActivo() bool {
	return t.estado == TransporteActivo
}

// ActualizarContacto permite modificar los datos de contacto del transporte.
func (t *Transporte) ActualizarContacto(contacto, telefono, correo string) {
	t.contacto = contacto
	t.telefono = telefono
	t.correo = correo
}
