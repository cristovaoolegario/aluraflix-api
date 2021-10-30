package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"net/http"
	"os"
)

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: ValidateToken,
	SigningMethod:       jwt.SigningMethodRS256,
})

func ValidateToken(token *jwt.Token) (interface{}, error){
		// Verify 'aud' claim
		audience := os.Getenv("AUD")
		checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(audience, false)
		if !checkAudience {
			return token, errors.New("Invalid audience.")
		}
		// Verify 'issuer' claim
		issuer := os.Getenv("ISS")
		checkIssuer := token.Claims.(jwt.MapClaims).VerifyIssuer(issuer, false)
		if !checkIssuer {
			return token, errors.New("Invalid issuer.")
		}

		cert, err := getPemCert(issuer, token)
		if err != nil {
			return nil, err
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
}

func getPemCert(issuer string, token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(fmt.Sprintf("%s.well-known/jwks.json", issuer))

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}
