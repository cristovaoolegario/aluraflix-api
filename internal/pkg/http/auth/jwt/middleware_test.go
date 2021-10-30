package jwt

import (
	"errors"
	"github.com/form3tech-oss/jwt-go"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestValidateToken_ShouldReturnTokenAndInvalidAudienceError_WhenTokenAudienceIsNotValid(t *testing.T) {
	os.Setenv("AUD", "https://unit-test-audience")
	claims := jwt.MapClaims{
		"aud": []string{"https://not-unit-test-audience"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := ValidateToken(token)

	assert.Equal(t, token, result)
	assert.Equal(t, "Invalid audience.", err.Error())
}

func TestValidateToken_ShouldReturnTokenAndInvalidIssuerError_WhenTokenIssuerIsNotValid(t *testing.T) {
	os.Setenv("AUD", "https://unit-test-audience")
	os.Setenv("ISS", "https://unit-test-issuer.us.auth0.com")
	claims := jwt.MapClaims{
		"iss": "https://not-unit-test-issuer.us.auth0.com/",
		"aud": []string{"https://unit-test-audience"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := ValidateToken(token)

	assert.Equal(t, token, result)
	assert.Equal(t, "Invalid issuer.", err.Error())
}

func TestValidateToken_ShouldReturnError_WhenTheresAProblemGettingTheCertificate(t *testing.T) {

	os.Setenv("AUD", "https://unit-test-audience/")
	os.Setenv("ISS", "https://unit-test-issuer.us.auth0.com/")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"https://unit-test-issuer.us.auth0.com/.well-known/jwks.json",
		func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("There's an error.")
		},
	)

	claims := jwt.MapClaims{
		"exp": 15000,
		"iss": "https://unit-test-issuer.us.auth0.com/",
		"aud": []string{"https://unit-test-audience/"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := ValidateToken(token)

	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestValidateToken_ShouldReturnError_WhenTheresAProblemDecodingTheCertificate(t *testing.T) {

	os.Setenv("AUD", "https://unit-test-audience/")
	os.Setenv("ISS", "https://unit-test-issuer.us.auth0.com/")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"https://unit-test-issuer.us.auth0.com/.well-known/jwks.json",
		httpmock.NewStringResponder(200, "resp string"),
	)

	claims := jwt.MapClaims{
		"exp": 15000,
		"iss": "https://unit-test-issuer.us.auth0.com/",
		"aud": []string{"https://unit-test-audience/"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := ValidateToken(token)

	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestValidateToken_ShouldReturnError_WhenTheresAProblemWithTheCertificateValidation(t *testing.T) {

	os.Setenv("AUD", "https://unit-test-audience/")
	os.Setenv("ISS", "https://unit-test-issuer.us.auth0.com/")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"https://unit-test-issuer.us.auth0.com/.well-known/jwks.json",
		httpmock.NewStringResponder(200, "resp string"),
	)

	claims := jwt.MapClaims{
		"exp": 15000,
		"iss": "https://unit-test-issuer.us.auth0.com/",
		"aud": []string{"https://unit-test-audience/"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := ValidateToken(token)

	assert.Nil(t, result)
	assert.NotNil(t, err)
}
