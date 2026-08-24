package services

import (
	"fmt"
	"strings"

	"sigi/interfaces"
	"sigi/models"
	"sigi/utils"
)

type ProveedorService struct {
	repository interfaces.ProveedorRepository
	generador  *utils.GeneradorCodigo
}

func NuevoProveedorService(
	repository interfaces.ProveedorRepository,
	generador *utils.GeneradorCodigo,
) *ProveedorService {
	return &ProveedorService{
		repository: repository,
		generador:  generador,
	}
}

func (s *ProveedorService) Registrar(
	empresa string,
	pais string,
	contacto string,
	telefono string,
	correo string,
) (*models.Proveedor, error) {
	if !utils.TextoValido(empresa) {
		return nil, fmt.Errorf("%w: la empresa del proveedor es obligatoria", utils.ErrDatoObligatorio)
	}

	if !utils.TextoValido(pais) {
		return nil, fmt.Errorf("%w: el pais del proveedor es obligatorio", utils.ErrDatoObligatorio)
	}

	if !utils.TextoValido(contacto) {
		return nil, fmt.Errorf("%w: el contacto del proveedor es obligatorio", utils.ErrDatoObligatorio)
	}

	if !utils.TelefonoValido(telefono) {
		return nil, fmt.Errorf("%w: el telefono del proveedor no es válido", utils.ErrDatoInvalido)
	}

	if strings.TrimSpace(correo) != "" && !utils.CorreoValido(correo) {
		return nil, fmt.Errorf("%w: el correo del proveedor no es válido", utils.ErrDatoInvalido)
	}

	codigo := s.generador.Generar("PRV")

	proveedor := models.NuevoProveedor(
		codigo,
		empresa,
		pais,
		contacto,
		telefono,
		correo,
	)

	if err := s.repository.Agregar(proveedor); err != nil {
		return nil, err
	}

	return proveedor, nil
}

func (s *ProveedorService) Buscar(codigo string) (*models.Proveedor, error) {
	if strings.TrimSpace(codigo) == "" {
		return nil, fmt.Errorf("%w: el codigo del proveedor es obligatorio", utils.ErrDatoObligatorio)
	}

	return s.repository.BuscarPorCodigo(codigo)
}

func (s *ProveedorService) ObtenerTodos() []*models.Proveedor {
	return s.repository.ObtenerTodos()
}

func (s *ProveedorService) Activar(codigo string) error {
	proveedor, err := s.Buscar(codigo)
	if err != nil {
		return err
	}

	proveedor.Activar()

	return s.repository.Actualizar(proveedor)
}

func (s *ProveedorService) Desactivar(codigo string) error {
	proveedor, err := s.Buscar(codigo)
	if err != nil {
		return err
	}

	proveedor.Desactivar()

	return s.repository.Actualizar(proveedor)
}

func (s *ProveedorService) ActualizarContacto(
	codigo string,
	contacto string,
	telefono string,
	correo string,
) error {
	if strings.TrimSpace(contacto) == "" {
		return fmt.Errorf("%w: el contacto es obligatorio", utils.ErrDatoObligatorio)
	}

	if strings.TrimSpace(telefono) == "" {
		return fmt.Errorf("%w: el telefono no es válido", utils.ErrDatoInvalido)
	}

	proveedor, err := s.Buscar(codigo)
	if err != nil {
		return err
	}

	proveedor.ActualizarContacto(contacto, telefono, correo)

	return s.repository.Actualizar(proveedor)
}

func (s *ProveedorService) Eliminar(codigo string) error {
	proveedor, err := s.Buscar(codigo)
	if err != nil {
		return err
	}

	if proveedor.EstaActivo() {
		return fmt.Errorf("%w: no se puede eliminar un proveedor activo", utils.ErrOperacionNoPermitida)
	}

	return s.repository.Eliminar(codigo)
}
