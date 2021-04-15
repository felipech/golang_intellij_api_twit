package routers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/intellij_test_golang/bd"
	"github.com/intellij_test_golang/models"
	"strings"
)

// Email variable usada en los endpoint
var Email string

// IDUsuario id usado en los endpoint
var IDUsuario string

// ProcesoToken se procesa el token para extraer los valores
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {

	miClave := []byte("Desarrollo_APP_Golang")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("token invalido")
	}
	return claims, false, "", err

}
