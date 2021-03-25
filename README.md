# json: Paquete que convierte texto json a objeto y objeto a texto json.

[![Go Report Card](https://goreportcard.com/badge/github.com/fabianpallares/json)](https://goreportcard.com/report/github.com/fabianpallares/json) [![GoDoc](https://godoc.org/github.com/fabianpallares/json?status.svg)](https://godoc.org/github.com/fabianpallares/json)

## Instalación:
Para instalar el paquete utilice la siguiente sentencia:
```
go get -u github.com/fabianpallares/json
```

## Convertir texto a objeto:
Para convertir un texto con formato json a un objeto (puntero de estructura), utilizar la siguiente función:

```GO
package main

import (
    "fmt"
    fpjson "github.com/fabianpallares/json"
)

func main() {
	var s = `{"apellidos": "Pallares", "nombres": "Fabian", "edad": 50, "domicilios": [
		{"calle":"Calle uno", "nro": 211, "provincia": "Buenos Aires"},	
		{"calle":"Calle dos", "nro": 300, "provincia": "CABA"}	
	]}`

	var p persona
	if err := fpjson.TextoAObjeto(s, &p, true); err != nil {
		t.Error(err)
	}

    fmt.Printf("%#v\n", p)
}
```

## Convertir objeto a texto:
Para convertir un objeto (puntero de estructura) a un texto con con formato json, utilizar la siguiente función:

```GO
package main

import (
    "fmt"
    fpjson "github.com/fabianpallares/json"
)

func main() {
	var p = persona{
		Apellidos: "Pallares",
		Nombres:   "Fabian",
		Edad:      50,
		Domicilios: []domicilio{
			{"Calle uno", 211, "Buenos Aires"},
			{"Calle dos", 300, "CABA"},
		},
	}

	s, err := fpjson.ObjetoATexto(p, true)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(s)
}
```

#### Documentación:
[Documentación en godoc](https://godoc.org/github.com/fabianpallares/json)