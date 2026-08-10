package services

import (
	"errors"
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
	if strings.TrimSpace(empresa) == "" {
		return nil, errors.New("la empresa del proveedor es obligatoria")
	}

	if strings.TrimSpace(pais) == "" {
		return nil, errors.New("el pais del proveedor es obligatorio")
	}

	if strings.TrimSpace(contacto) == "" {
		return nil, errors.New("el contacto del proveedor es obligatorio")
	}

	if strings.TrimSpace(telefono) == "" {
		return nil, errors.New("el telefono del proveedor es obligatorio")
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
		return nil, errors.New("el codigo del proveedor es obligatorio")
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
		return errors.New("el contacto es obligatorio")
	}

	if strings.TrimSpace(telefono) == "" {
		return errors.New("el telefono es obligatorio")
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
		return errors.New("no se puede eliminar un proveedor activo")
	}

	return s.repository.Eliminar(codigo)
}
