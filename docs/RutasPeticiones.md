# Rutas y Peticiones en Gin

Este documento describe cómo trabajar con rutas y peticiones en el framework Gin para Go, incluyendo la creación de grupos de rutas, modularización y mejores prácticas.

## Grupos de Rutas en Gin

En Gin, puedes crear grupos de rutas para organizar endpoints relacionados bajo un prefijo común. Esto permite compartir middleware y mantener el código organizado.

### Sintaxis Básica

```go
group := r.Group("/prefijo")
group.GET("/subruta", handler)
```

Esto crea rutas como `/prefijo/subruta`.

### Ejemplo de Implementación

En [`src/routes/routes.go`](src/routes/routes.go), se define un grupo de rutas con prefijo `/prefijo`:

```go
package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    RoutesGroup := r.Group("/prefijo")
    RoutesGroup.GET("/", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{
            "SAS": "Down",
        })
    })
    RoutesGroup.GET("/spotify/:contex", func(ctx *gin.Context) {
        nombre := ctx.Param("contex")
        ctx.String(http.StatusOK, "Que es dios? %s", nombre)
    })
}
```

## Modularización de Rutas

Para mantener el código organizado, es recomendable separar las rutas en un paquete dedicado como `routes`. Esto facilita la reutilización y el mantenimiento.

### Llamada desde main.go

En [`main.go`](main.go), se crea el router y se pasan las rutas:

```go
package main

import (
    "Gin/src/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.New()
    routes.SetupRoutes(r)
    r.Run(":8000")
}
```

## Mejores Prácticas

- **Pasar el Router como Parámetro**: Crea el router en `main.go` y pásalo a las funciones de configuración. Esto permite añadir middleware global y facilita la escalabilidad.
- **Uso de Grupos**: Agrupa rutas relacionadas para compartir prefijos y middleware.
- **Parámetros de Ruta**: Asegúrate de que los nombres de parámetros coincidan entre la definición de la ruta y `ctx.Param()`.
- **Middleware**: Añade middleware al grupo o al router global según sea necesario.

### Ventajas de la Modularización

- Separación de responsabilidades.
- Facilita pruebas unitarias.
- Permite combinar rutas de múltiples módulos.

## Pruebas

Para probar las rutas:

- `curl http://localhost:8000/prefijo/` → Respuesta JSON: `{"SAS": "Down"}`
- `curl http://localhost:8000/prefijo/spotify/ejemplo` → Respuesta: `Que es dios? ejemplo`

## Referencias

- [Documentación oficial de Gin](https://gin-gonic.com/docs/)
