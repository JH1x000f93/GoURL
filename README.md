# GoUrl

GoUrl es un paquete en Go diseñado para acortar URLs utilizando un servicio de acortamiento de URL externo. Este paquete incluye una función llamada `Short` que facilita la tarea de acortar URLs largas.

## Uso

Para utilizar este paquete, primero importa el paquete `GoUrl` en tu proyecto:

```shell
go get github.com/JH1x000f93/GoURL
```
```go
import "github.com/JH1x000f93/GoURL"
```
A continuación, puedes usar la función Short de la siguiente manera:

```go
package main

import (
    "fmt"

    "github.com/JH1x000f93/GoURL"
)
func main(){
    url := "https://example.com/long"
    apiKey := "678GHJghj78IG0osh7GJi7hmK"

    shortResponse, err := GoUrl.Short(url, apiKey)
    if err != nil {
        fmt.Fatal(err)
    }

    if shortResponse.Status == "OK" {
        fmt.Println(shortResponse.Short)
    } else {
        fmt.Println("Error: ", shortResponse.Info)
    }


```
La función Short toma dos argumentos: la URL larga que deseas acortar y una APIKEY para autenticarte con el servicio de acortamiento de URL externo. Devolverá una estructura ShortResponse que contiene información sobre la respuesta, incluido el enlace acortado si la solicitud se realizó correctamente.

Puedes optener una APIKEY [Aqui](https://app.shorten.cl/es/apikey)

## Estructura de Respuesta (ShortResponse)
La estructura ShortResponse se utiliza para representar la respuesta JSON del servicio de acortamiento de URL. Puede incluir los siguientes campos:

- Status: El estado de la solicitud.
- Permalink: El enlace permanente (si está disponible).
- Short: El enlace acortado.
- Info: Información adicional (si está disponible).

## Contribución
Si deseas contribuir a este proyecto, ¡no dudes en abrir una solicitud de extracción (pull request)!