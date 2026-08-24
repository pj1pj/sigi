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
* Actualizar información de contacto mediante el Service (la consola no expone esta opción).
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
* Actualizar información de contacto mediante el Service (la consola no expone esta opción).
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
En preparación
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

├── api/
│   ├── dto.go
│   ├── response.go
│   ├── server.go
│   └── server_test.go

└── app/
    └── sistema.go
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
| 🌐 `net/http` | Servidor y endpoints HTTP           |
| 🧾 `encoding/json` | Serialización y deserialización JSON |
| 🧪 `httptest` | Pruebas de la API sin red externa    |

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

El modo predeterminado es la aplicación de consola. También puede indicarse
explícitamente:

```bash
go run . -modo consola
```

## 🌐 API HTTP y JSON

SIGI incluye una API HTTP basada únicamente en `net/http` y `encoding/json` de
la biblioteca estándar. Para iniciarla en `localhost:8080`:

```bash
go run . -modo api
```

Los datos se almacenan en memoria, por lo que se reinician cuando se detiene el
proceso. Cada ejecución usa un sistema independiente: el modo consola y el modo
API no comparten datos entre procesos.

### Endpoints

| Método | Ruta | Funcionalidad |
| --- | --- | --- |
| `POST` | `/api/v1/proveedores` | Registrar proveedor |
| `GET` | `/api/v1/proveedores` | Listar proveedores |
| `GET` | `/api/v1/proveedores/{codigo}` | Consultar proveedor |
| `PATCH` | `/api/v1/proveedores/{codigo}/estado` | Activar o desactivar proveedor |
| `POST` | `/api/v1/transportes` | Registrar transporte |
| `GET` | `/api/v1/transportes` | Listar transportes |
| `POST` | `/api/v1/ordenes` | Crear orden de compra |
| `POST` | `/api/v1/ordenes/{codigo}/productos` | Agregar producto a una orden |
| `PATCH` | `/api/v1/ordenes/{codigo}/confirmacion` | Confirmar una orden |
| `GET` | `/api/v1/ordenes` | Listar órdenes |
| `POST` | `/api/v1/importaciones` | Registrar importación |
| `PATCH` | `/api/v1/importaciones/{codigo}/tracking` | Avanzar tracking de importación |
| `GET` | `/api/v1/inventario` | Consultar inventario |
| `GET` | `/api/v1/reportes/general` | Consultar reporte general |

### Códigos HTTP principales

- `200 OK`: consulta o actualización exitosa.
- `201 Created`: creación exitosa de proveedor, transporte, orden o importación.
- `400 Bad Request`: JSON inválido o datos que no cumplen las validaciones.
- `404 Not Found`: recurso o ruta inexistente.
- `405 Method Not Allowed`: método no permitido para una ruta existente.
- `409 Conflict`: operación incompatible con el estado actual del dominio.

### Ejemplos JSON

Registrar un proveedor:

```json
{
  "empresa": "ACME Importaciones",
  "pais": "China",
  "contacto": "Li Wei",
  "telefono": "+86 1234567",
  "correo": "li@acme.example"
}
```

Crear una orden:

```json
{ "codigo_proveedor": "PRV-0001" }
```

Agregar un producto:

```json
{
  "nombre": "Computadora portátil",
  "cantidad": 3,
  "precio_unitario": 850.50
}
```

Registrar una importación:

```json
{
  "codigo_orden": "ORD-0001",
  "codigo_transporte": "TRN-0001",
  "ciudad_origen": "Shenzhen",
  "ciudad_destino": "Guayaquil"
}
```

Actualizar el tracking:

```json
{ "estado": "En tránsito" }
```

Para la llegada a bodega se debe indicar la ubicación, lo que genera el
inventario automáticamente:

```json
{ "estado": "Llegó a bodega", "ubicacion": "Bodega A-01" }
```

Ejemplo de respuesta de error uniforme:

```json
{ "error": "no se puede confirmar una orden sin productos" }
```

En PowerShell se puede realizar una solicitud así:

```powershell
$proveedor = @{
  empresa = "ACME Importaciones"
  pais = "China"
  contacto = "Li Wei"
  telefono = "+86 1234567"
  correo = "li@acme.example"
} | ConvertTo-Json

Invoke-RestMethod -Method Post `
  -Uri http://localhost:8080/api/v1/proveedores `
  -ContentType "application/json" `
  -Body $proveedor
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

El paquete `api` sí contiene pruebas automatizadas con `httptest`. La suite
comprueba el flujo completo, respuestas JSON, validaciones, recursos
inexistentes, métodos HTTP no permitidos, transiciones de tracking y generación
automática de inventario.

También se puede verificar estáticamente el proyecto con:

```bash
go vet ./...
```

Cuando el entorno tiene CGo y un compilador C disponible, también puede
ejecutarse `go test -race ./...` para detectar carreras de datos. La API y sus
componentes en memoria están protegidos con mutexes para atender solicitudes
concurrentes de forma segura.

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

## 🎓 Relación con las cuatro unidades

### Unidad 1

El proyecto utiliza sintaxis Go, variables, constantes, condicionales,
`switch`, ciclos `for`, funciones, métodos, paquetes e importaciones.

### Unidad 2

Se utilizan slices, mapas, structs, punteros, constructores y métodos para
modelar proveedores, órdenes, productos, transportes, importaciones e inventario.

### Unidad 3

Los modelos mantienen campos privados y exponen métodos de acceso. Las
interfaces de repositorio desacoplan los Services de los repositorios en
memoria y permiten polimorfismo mediante inyección de dependencias. Los
errores identificables se comparten mediante `utils/errors.go`.

### Unidad 4

La API utiliza servicios web HTTP con `net/http` y serialización JSON con
`encoding/json`. Las pruebas HTTP utilizan `httptest`. El servidor recibe
solicitudes concurrentes y protege los repositorios, el generador de códigos y
el acceso a los modelos compartidos mediante mutexes. No se agregaron
goroutines ni canales artificiales porque la lógica de negocio no requiere una
tarea asíncrona independiente.

## Limitaciones y mejoras futuras

Actualmente los datos se almacenan solo en memoria, no existe autenticación y
la API no expone todas las operaciones disponibles en los Services, como
actualización de contactos o cancelación de órdenes. Tampoco hay persistencia,
paginación ni configuración externa del puerto. Los cuerpos JSON tienen un
límite de 1 MiB para evitar solicitudes excesivamente grandes. Como mejoras futuras se
podrían agregar una base de datos, autenticación, apagado graceful, métricas,
paginación y endpoints adicionales sin trasladar reglas de negocio fuera de
los Services.

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
