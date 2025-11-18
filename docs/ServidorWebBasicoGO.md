# Servidor Web Básico en Go con Gin

Este documento describe un servidor web básico implementado en Go utilizando el framework Gin.

## Descripción

El código crea un servidor web simple que responde a solicitudes GET en la ruta raíz ("/") con un mensaje JSON.

## Dependencias

- [Gin](https://github.com/gin-gonic/gin): Framework web para Go.

Para instalar la dependencia, ejecuta:

```bash
go get github.com/gin-gonic/gin
```

## Código

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.New()
    r.GET("/", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "mensaje": "hola",
        })
    })

    r.Run(":8000")
}
```

## Explicación del Código

1. **Importación**: Se importa el paquete `github.com/gin-gonic/gin` para utilizar el framework Gin.

2. **Creación del Router**: `r := gin.New()` crea una nueva instancia del router de Gin.

3. **Definición de Ruta**: `r.GET("/", ...)` define una ruta GET para la raíz ("/"). Cuando se accede a esta ruta, se ejecuta la función anónima que responde con un JSON.

4. **Respuesta JSON**: `ctx.JSON(200, gin.H{"mensaje": "hola"})` envía una respuesta HTTP con código de estado 200 (OK) y un objeto JSON que contiene el mensaje "hola".

5. **Ejecución del Servidor**: `r.Run(":8000")` inicia el servidor en el puerto 8000.

## Cómo Ejecutar

1. Asegúrate de tener Go instalado.
2. Instala las dependencias: `go mod tidy` (si usas módulos de Go).
3. Ejecuta el programa: `go run main.go`

El servidor se iniciará en `http://localhost:8000`.

## Prueba

Puedes probar el servidor abriendo un navegador y visitando `http://localhost:8000`, o usando curl:

```bash
curl http://localhost:8000
```

Deberías recibir una respuesta JSON como:

```json
{
  "mensaje": "hola"
}
```
