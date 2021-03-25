package json

import (
	"fmt"
	"testing"
)

var (
	token = ""
)

type persona struct {
	Apellidos  string `json:"apellidos"`
	Nombres    string `json:"nombres"`
	Edad       int
	Domicilios []domicilio `json:"domicilios"`
}

type domicilio struct {
	Calle     string `json:"calle"`
	Numero    int    `json:"nro"`
	Provincia string `json:"provincia"`
}

func TestTextoAObjeto(t *testing.T) {
	var s = `{"apellidos": "Pallares", "nombres": "Fabian", "edad": 50, "domicilios": [
		{"calle":"Calle uno", "nro": 211, "provincia": "Buenos Aires"},	
		{"calle":"Calle dos", "nro": 300, "provincia": "CABA"}	
	]}`

	var p persona
	if err := TextoAObjeto(s, &p, true); err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", p)
}
func TestObjetoATexto(t *testing.T) {
	var p = persona{
		Apellidos: "Pallares",
		Nombres:   "Fabian",
		Edad:      50,
		Domicilios: []domicilio{
			{"Calle uno", 211, "Buenos Aires"},
			{"Calle dos", 300, "CABA"},
		},
	}

	s, err := ObjetoATexto(p, true)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(s)
}
