package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

// ObjetoABytes convierte el objeto recibido (no es necesario que sea un puntero de una estructura)
// a []byte.
func ObjetoABytes(objeto any, conIndentacion bool) ([]byte, error) {
	return objetoABytes(objeto, conIndentacion)
}

// ObjetoATexto convierte el objeto recibido (no es necesario que sea un puntero de una estructura)
// a un texto con formato JSON.
func ObjetoATexto(objeto any, conIndentacion bool) (string, error) {
	b, err := objetoABytes(objeto, conIndentacion)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// BytesAObjeto utiliza el slice de bytes recibido (debe contener un formato json válido)
// para completar/rellenar el objeto (debe ser un puntero de una estructura).
func BytesAObjeto(b []byte, objeto any, validarCamposDesconocidos bool) error {
	var descifrador = json.NewDecoder(bytes.NewReader(b))
	return aObjeto(descifrador, objeto, validarCamposDesconocidos)
}

// TextoAObjeto utiliza el texto recibido (debe contener un formato json válido)
// para completar/rellenar el objeto (debe ser un puntero de una estructura).
func TextoAObjeto(s string, objeto any, validarCamposDesconocidos bool) error {
	var descifrador = json.NewDecoder(strings.NewReader(s))
	return aObjeto(descifrador, objeto, validarCamposDesconocidos)
}

func objetoABytes(objeto any, conIndentacion bool) ([]byte, error) {
	var b []byte
	var err error
	if conIndentacion {
		b, err = json.MarshalIndent(objeto, "", "    ")
	} else {
		b, err = json.Marshal(objeto)
	}
	if err != nil {
		return nil, fmt.Errorf("Error al convertir el objeto recibido a formato JSON: %w", err)
	}

	return b, nil
}

func aObjeto(descifrador *json.Decoder, objeto any, validarCamposDesconocidos bool) error {
	if validarCamposDesconocidos {
		// no permite recibir campos desconocidos.
		descifrador.DisallowUnknownFields()
	}

	var err = descifrador.Decode(objeto)
	if err == nil {
		return nil
	}

	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	switch {
	case errors.Is(err, io.EOF):
		// verificar que se haya recibido información en el cuerpo del mensaje.
		return fmt.Errorf("El formato JSON recibido es incorrecto. Contenido vacío")

	case errors.Is(err, io.ErrUnexpectedEOF):
		// verificar la lectura del cuerpo del mensaje.
		return fmt.Errorf("El formato JSON recibido es incorrecto. Se ha llegado al final de la lectura de manera inesperada")

	case errors.As(err, &syntaxError):
		// verificar si el formato es correcto, si faltan dobles comillas,
		// comillas, comas, llaves, corchetes; etc.
		return fmt.Errorf("El formato JSON recibido es incorrecto. Error en la posición %v", syntaxError.Offset)

	case errors.As(err, &unmarshalTypeError):
		// verificar si hay un error de tipo de campo (valores erroneos).
		var valor string
		switch unmarshalTypeError.Value {
		case "number":
			valor = "numérico"
		case "string":
			valor = "texto"
		case "bool":
			valor = "lógico"
		default:
			valor = unmarshalTypeError.Value
		}
		return fmt.Errorf("El formato JSON recibido es incorrecto. Error en el campo \"%v\", tipo de valor recibido %v, posición %v", unmarshalTypeError.Field, valor, unmarshalTypeError.Offset)

	case strings.HasPrefix(err.Error(), "json: unknown field "):
		// verificar si se recibieron campos adicionales que no están en la
		// estructura recibida.
		var campo = strings.TrimPrefix(err.Error(), "json: unknown field ")
		return fmt.Errorf("El formato JSON recibido es incorrecto. Se ha recibido un nombre de campo inexistente: %v", campo)

	case strings.HasPrefix(err.Error(), "json: Unmarshal(non-pointer "):
		// verificar si el objeto recibido es un puntero de una estructura.
		return fmt.Errorf("El objeto recibido no es un puntero de una estructura")

	case err.Error() == "http: request body too large":
		// verificar contenido muy largo.
		return fmt.Errorf("El formato JSON recibido es incorrecto. El texto recibido es demasiado grande")
	}

	// cualquier otro tipo de error.
	return fmt.Errorf("El formato JSON recibido es incorrecto: %w", err)
}
