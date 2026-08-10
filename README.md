# 🚢 SIGI — Sistema Integral de Gestión de Importaciones

## Información general

SIGI es una aplicación de consola desarrollada en Go para gestionar de manera integral diferentes operaciones relacionadas con la importación de mercancías.

El sistema permite administrar proveedores, órdenes de compra, productos, transportes, importaciones, tracking, inventario y reportes.

La información se maneja en memoria durante la ejecución y el proyecto utiliza la biblioteca estándar de Go.

---

## Objetivo

Representar de manera sencilla y modular el flujo de una operación de importación, aplicando conceptos de programación orientada a objetos y separación de responsabilidades.

### Flujo principal

Proveedor
    ↓
Orden de compra
    ↓
Productos
    ↓
Confirmación
    ↓
Transporte
    ↓
Importación
    ↓
Tracking
    ↓
Llegada a bodega
    ↓
Inventario
    ↓
Reporte general

---

---

## ✨ Funcionalidades principales

### 👤 Proveedores

* Registrar proveedores.
* Listar proveedores.
* Activar y desactivar proveedores.
* Actualizar información de contacto.
* Eliminar proveedores inactivos.
* Generación automática de códigos.

Ejemplo:

```text
PRV-0001
PRV-0002
PRV-0003
```

---

### 🛒 Órdenes de compra

* Crear órdenes asociadas a proveedores.
* Agregar múltiples productos.
* Registrar cantidades.
* Registrar precios unitarios.
* Calcular subtotales.
* Calcular el total de la orden.
* Confirmar órdenes.
* Cancelar órdenes cuando corresponde.
* Generación automática de códigos.

Ejemplo:

```text
ORD-0001
ORD-0002
```

---

### 🚚 Transportes

Permite administrar empresas de transporte de diferentes tipos:

* 🚢 Marítimo
* ✈️ Aéreo
* 🚛 Terrestre

También permite:

* Registrar transportes.
* Listarlos.
* Activarlos y desactivarlos.
* Actualizar información de contacto.
* Eliminar transportes inactivos.

Códigos automáticos:

```text
TRN-0001
TRN-0002
```

---

### 🌎 Importaciones

Una importación relaciona:

* Una orden de compra confirmada.
* Un transporte activo.
* Ciudad de origen.
* Ciudad de destino.

Los códigos se generan automáticamente:

```text
IMP-0001
IMP-0002
```

El usuario no necesita escribir manualmente el código.

---

## 📍 Tracking

Las importaciones utilizan un flujo controlado de estados:

```text
En preparación
      ↓
En tránsito
      ↓
En aduana
      ↓
Llegada a bodega
```

El sistema valida las transiciones para evitar saltos de estados no permitidos.

### 📦 Llegada a bodega

Cuando una importación llega a bodega:

```text
Importación
     ↓
Llegada a bodega
     ↓
Procesamiento automático
     ↓
Actualización del inventario
```

De esta manera, el inventario no necesita movimientos manuales desde el menú.

---

## 📦 Inventario

El módulo de inventario permite consultar:

* Producto.
* Cantidad disponible.
* Proveedor.
* Importación.
* Ubicación.

La información del inventario se actualiza automáticamente como consecuencia del proceso de importación.

---

## 📊 Reporte general

SIGI utiliza un **reporte general resumido** para evitar mostrar grandes cantidades de registros individualmente.

El reporte muestra información como:

```text
PROVEEDORES
Total
Activos
Inactivos

ORDENES DE COMPRA
Total
Confirmadas
Canceladas
Total de compras

TRANSPORTES
Total
Activos
Inactivos

IMPORTACIONES
Total
En tránsito
En aduana
En bodega

INVENTARIO
Registros
Unidades disponibles
```

También muestra el **total de dinero correspondiente a las órdenes de compra**.

---

# 🧱 Arquitectura

El proyecto utiliza una organización modular por paquetes:

```text
sigi/
│
├── main.go
│
├── models/
│   ├── proveedor.go
│   ├── producto.go
│   ├── orden_compra.go
│   ├── transporte.go
│   ├── importacion.go
│   └── inventario.go
│
├── interfaces/
│   └── repositories.go
│
├── repository/
│   ├── proveedor_repository.go
│   ├── orden_repository.go
│   ├── transporte_repository.go
│   ├── importacion_repository.go
│   └── inventario_repository.go
│
├── services/
│   ├── proveedor_service.go
│   ├── orden_service.go
│   ├── transporte_service.go
│   ├── importacion_service.go
│   ├── inventario_service.go
│   └── reporte_service.go
│
├── menu/
│   ├── menu.go
│   ├── proveedores_menu.go
│   ├── ordenes_menu.go
│   ├── transporte_menu.go
│   ├── importaciones_menu.go
│   ├── inventario_menu.go
│   └── reportes_menu.go
│
└── utils/
    ├── codigo.go
    ├── errors.go
    ├── input.go
    └── validation.go
```

---

## 🧩 Responsabilidad de cada paquete

| Paquete      | Responsabilidad                              |
| ------------ | -------------------------------------------- |
| `models`     | Entidades y comportamiento del dominio       |
| `interfaces` | Contratos de los repositorios                |
| `repository` | Gestión de datos en memoria                  |
| `services`   | Reglas de negocio                            |
| `menu`       | Interacción con el usuario                   |
| `utils`      | Funciones auxiliares y generación de códigos |

---

## 🔢 Generación automática de códigos

El sistema cuenta con un generador centralizado que mantiene contadores independientes por tipo de entidad.

```text
PRV-0001  → Proveedor
PRV-0002  → Proveedor

ORD-0001  → Orden
ORD-0002  → Orden

TRN-0001  → Transporte
TRN-0002  → Transporte

IMP-0001  → Importación
IMP-0002  → Importación
```

Cada prefijo posee su propio contador.

---

# 🛠️ Tecnologías utilizadas

| Tecnología    | Uso                                 |
| ------------- | ----------------------------------- |
| 🐹 Go         | Lenguaje principal                  |
| 💻 CLI        | Interfaz del sistema                |
| 🧩 Structs    | Modelado de entidades               |
| 🔌 Interfaces | Abstracción de repositorios         |
| 📦 Packages   | Organización modular                |
| 💾 Memoria    | Almacenamiento durante la ejecución |

> No se utiliza una base de datos ni librerías externas.

---

# 🚀 Instalación

## Requisitos

Se necesita tener instalado:

* **Go**
* **Git** (para clonar el repositorio)

---

## 1️⃣ Clonar el repositorio

```bash
git clone https://github.com/pj1pj/sigi.git
```

## 2️⃣ Entrar al proyecto

```bash
cd sigi/sigi
```

## 3️⃣ Ejecutar

```bash
go run .
```

---

# 🧪 Verificación

Para comprobar que todos los paquetes compilan correctamente:

```bash
go test ./...
```

La salida puede mostrar:

```text
[no test files]
```

Esto significa que el paquete no contiene archivos de pruebas automatizadas y **no representa un error**.

---

# 🧹 Formatear código

Go incluye `gofmt` para mantener un formato estándar.

Ejemplo:

```bash
gofmt -w main.go
```

En PowerShell también se pueden formatear los archivos de un directorio:

```powershell
Get-ChildItem menu -Filter *.go | ForEach-Object { gofmt -w $_.FullName }
```

---

# 🔄 Ejemplo de uso completo

Una prueba completa del sistema puede realizarse siguiendo este flujo:

```text
1. Registrar proveedor
        ↓
2. Registrar transporte
        ↓
3. Crear orden
        ↓
4. Agregar varios productos
        ↓
5. Confirmar orden
        ↓
6. Registrar importación
        ↓
7. Cambiar a "En tránsito"
        ↓
8. Cambiar a "En aduana"
        ↓
9. Cambiar a "Llegada a bodega"
        ↓
10. Consultar inventario
        ↓
11. Consultar reporte general
```

Ejemplo de códigos generados:

```text
PRV-0001
TRN-0001
ORD-0001
IMP-0001
```

---

# 📚 Conceptos de Go aplicados

El proyecto utiliza diferentes conceptos del lenguaje:

* `struct`
* Métodos
* Interfaces
* Punteros
* Slices
* Mapas
* Paquetes
* Manejo de errores
* Funciones
* Encapsulamiento mediante métodos
* Separación de responsabilidades

---

# 🎓 Proyecto académico

**SIGI — Sistema Integral de Gestión de Importaciones**

Proyecto desarrollado para la asignatura de **Programación en Go**.

---

## 🔗 Repositorio

[![GitHub](https://img.shields.io/badge/GitHub-pj1pj%2Fsigi-181717?style=for-the-badge\&logo=github)](https://github.com/pj1pj/sigi)

---

<p align="center">
  🚢 <b>SIGI</b> — Gestión integral de importaciones en Go
  <br>
  <sub>Proyecto académico</sub>
</p>
```
