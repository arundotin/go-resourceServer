package server

import (
	"encoding/json"
	"errors"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	X5t string `json:"x5t"`
	Alg string   `json:"alg"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func ADFSJWTTokenValidationMiddleware ()  jwtmiddleware.JWTMiddleware {

	audienceId := "urn:arunlabs:resource:api_arunlabsapi:qa"
	tokenIssuer := "https://masked-domain.com/adfs/services/trust"

	jwtValidationMiddleware:= jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			validAudience := token.Claims.(jwt.MapClaims).VerifyAudience(audienceId, false)

			if !validAudience {
				return token, errors.New("Invalid audience.")
			}

			validIss := token.Claims.(jwt.MapClaims).VerifyIssuer(tokenIssuer,false)

			if !validIss {
				panic("Invalid Issuer")
			}

			certificate, err := getPemCert(token)

			if err!=nil {
				return token, errors.New("Invalid issuer.")
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(certificate))
			return result, nil
		},

		SigningMethod: jwt.SigningMethodRS256,
	})

	return *jwtValidationMiddleware
}

func getPemCert (token *jwt.Token) (string, error) {
	certificate := ""
	adfsKeysUrl := "https://masked-domain.com/adfs/discovery/keys"
	adfsKeysResponse, err := http.Get(adfsKeysUrl)

	if (err!=nil) {
		return certificate,err
	}

	defer adfsKeysResponse.Body.Close()

	var jwks = Jwks{}

	err = json.NewDecoder(adfsKeysResponse.Body).Decode(&jwks)

	if (err!=nil) {
		return certificate,err
	}

	for keyId,_ := range jwks.Keys {
		x5t_ := token.Header["x5t"]
		if x5t_ == jwks.Keys[keyId].X5t {
			certificate = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[keyId].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if certificate == "" {
		err := errors.New("Unable to find appropriate key.")
		return certificate, err
	}

	return certificate, nil


}

