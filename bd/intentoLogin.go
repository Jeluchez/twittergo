package bd

import (
	"github.com/jeluchez/twittergo/models"
	"golang.org/x/crypto/bcrypt"
)

// IntentoLogin verifica el chequelo de login en la DB
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	// verifica si el imail existe
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}
	// verificar que las password con la cual se intenta logear es la misma que la guardad en la DB
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
