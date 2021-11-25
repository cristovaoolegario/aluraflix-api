package jwt

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/form3tech-oss/jwt-go"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestValidateToken(t *testing.T) {
	t.Run("Should return token and invalid audience error when token audience is not valid", func(t *testing.T) {
		os.Setenv("AUD", "https://unit-test-audience")
		claims := jwt.MapClaims{
			"aud": []string{"https://not-unit-test-audience"},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
		result, err := ValidateToken(token)

		assert.Equal(t, token, result)
		assert.Equal(t, "invalid audience", err.Error())
	})

	t.Run("Should return token and invalid issuer error when token issuer is not valid", func(t *testing.T) {
		os.Setenv("AUD", "https://unit-test-audience")
		os.Setenv("ISS", "https://unit-test-issuer.us.auth0.com")
		claims := jwt.MapClaims{
			"iss": "https://not-unit-test-issuer.us.auth0.com/",
			"aud": []string{"https://unit-test-audience"},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
		result, err := ValidateToken(token)

		assert.Equal(t, token, result)
		assert.Equal(t, "invalid issuer", err.Error())
	})

	t.Run("Should return error when theres a problem getting the certificate", func(t *testing.T) {
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
	})

	t.Run("Should return error when theres a problem decoding the certificate", func(t *testing.T) {
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
	})

	t.Run("Should return error when theres a problem with the certificate validation", func(t *testing.T) {
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
	})

	t.Run("Should return unable to find appropriate key when theres no equivalent kid key on the certificate", func(t *testing.T) {
		os.Setenv("AUD", "https://unit-test-audience/")
		os.Setenv("ISS", "https://unit-test-issuer.us.auth0.com/")

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		json := `{"keys":[{"alg":"RS256","kty":"RSA","use":"sig","n":"x-N4R5lgHyXWfjf-izlxrrr2LAn7bUq1cL069yB0G4sy9FCM1RBeet1tHeQ3szbCxYxZIZ1ODRu9BuK34TyEkyBNtAOITU5WjUVuMrWd9iK-noIVJwhykLooGwHVSUCMPjeRNd7sxf2WW3uwR1R3PglZeu25pBR0e9PxI8tUU8QWsMOdrCRw5tMyoqC5SQsa1J4HIzuTaYfuOClF4kpv933_c79VquvdrWEJ1MzDHG2Lfrb_wxaFuOMXzPSTnOsINWwG2-0rb0UXXm_emsa8NrDu2Wi-nlw0UYAwVUQEtwK_5KegZWI39pp3aaDR62jdiEiL85BulrEjefxGZVHDUw","e":"AQAB","kid":"M0Xo-5mQq2nlEDgkbZeEm","x5t":"4Qi2aVFszWUy5QBdl52a-BLZjZU","x5c":["MIIDETCCAfmgAwIBAgIJWh9HI7fDZC1/MA0GCSqGSIb3DQEBCwUAMCYxJDAiBgNVBAMTG2FsdXJhLWZsaXgtYXBpLnVzLmF1dGgwLmNvbTAeFw0yMTA4MTQwNDQ2NDlaFw0zNTA0MjMwNDQ2NDlaMCYxJDAiBgNVBAMTG2FsdXJhLWZsaXgtYXBpLnVzLmF1dGgwLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMfjeEeZYB8l1n43/os5ca669iwJ+21KtXC9OvcgdBuLMvRQjNUQXnrdbR3kN7M2wsWMWSGdTg0bvQbit+E8hJMgTbQDiE1OVo1FbjK1nfYivp6CFScIcpC6KBsB1UlAjD43kTXe7MX9llt7sEdUdz4JWXrtuaQUdHvT8SPLVFPEFrDDnawkcObTMqKguUkLGtSeByM7k2mH7jgpReJKb/d9/3O/Varr3a1hCdTMwxxti362/8MWhbjjF8z0k5zrCDVsBtvtK29FF15v3prGvDaw7tlovp5cNFGAMFVEBLcCv+SnoGViN/aad2mg0eto3YhIi/OQbpaxI3n8RmVRw1MCAwEAAaNCMEAwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQU/FadD6LzgRq42C86qGuHfwmzlB0wDgYDVR0PAQH/BAQDAgKEMA0GCSqGSIb3DQEBCwUAA4IBAQBNDGgCFs5Wy71637Zon7VDEP8LdnsaeAACedMEJKMxh80AifEQviqSufo9LWgck4vsSfTeAWREDPxJ7rFhh4siHemQpm+8fExPmZc1NSH0+2xaPGJfeBX+GrUAVlHmObzbgChfKXvOI07+41JmxCKqTYAbu5/AHCAwOyF65JS3XEiatmmisuECOoM71+QSMxNhJOFMUK9Rysjb5XidpFB3mC2OLFy7SEvHbZuGUyS+sE4k9xSYl5zxO+DO8e2dCGdDs3MKX8XNIEvnTdR65i6gm0+a1/aastr4GNNvbPxiI7ELBFcn6iWI/0zL54Rvv5rc0WWJ6t772hDG3+JCJqs9"]},{"alg":"RS256","kty":"RSA","use":"sig","n":"malZ2q_aHX7VD_ykryOYQIOHmyKT1Q94rdUKZGLxp0Rw0s_livESCmOgrKqLxVjEQmUUokqThMhAiDi7OPcrzy150iYk5J7wmj-D3eDvFFiABnBDlvt2lSLPmUY4R-NTRQ1wNfbLKmQycOrWTAGT9P4VXp45IARuRdFtjU9lsXmifWpCEcLlv61WPMzL0b9ld_GBAWvvE-sbINOpzm_xBrPwcIsImQNAsN9mFmZSSaiVQ7bQOpExergecF39yaTxXA0PfSorcsVW6XEvi3UQgS9HCdVjX2VXuCdu_HvnC-rRuqXrXPqSMq3QmPvqLwWK53DEhCxroHGKKoG2CKgbTw","e":"AQAB","kid":"GKhfLaIlbtpIESk_Aedrc","x5t":"gvRF9c9nmTlsv0-O0_Oik4lOBjc","x5c":["MIIDETCCAfmgAwIBAgIJb279+r/8IMU9MA0GCSqGSIb3DQEBCwUAMCYxJDAiBgNVBAMTG2FsdXJhLWZsaXgtYXBpLnVzLmF1dGgwLmNvbTAeFw0yMTA4MTQwNDQ2NDlaFw0zNTA0MjMwNDQ2NDlaMCYxJDAiBgNVBAMTG2FsdXJhLWZsaXgtYXBpLnVzLmF1dGgwLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAJmpWdqv2h1+1Q/8pK8jmECDh5sik9UPeK3VCmRi8adEcNLP5YrxEgpjoKyqi8VYxEJlFKJKk4TIQIg4uzj3K88tedImJOSe8Jo/g93g7xRYgAZwQ5b7dpUiz5lGOEfjU0UNcDX2yypkMnDq1kwBk/T+FV6eOSAEbkXRbY1PZbF5on1qQhHC5b+tVjzMy9G/ZXfxgQFr7xPrGyDTqc5v8Qaz8HCLCJkDQLDfZhZmUkmolUO20DqRMXq4HnBd/cmk8VwND30qK3LFVulxL4t1EIEvRwnVY19lV7gnbvx75wvq0bql61z6kjKt0Jj76i8FiudwxIQsa6BxiiqBtgioG08CAwEAAaNCMEAwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUdsVi3xAtWNTvD8hYUbjerqtCTbkwDgYDVR0PAQH/BAQDAgKEMA0GCSqGSIb3DQEBCwUAA4IBAQBumE6HlpDk8Gw8KSkkay75qfWzx3meilu3RqpcoKEXausq70Xr5HfVnXl493trW5aBwgZCn5OzPfWWTIi4XpmSMeAwZRM9zJ3WfMQzO/M0ObF7K5s3wYLcc0t+djha/dZggdiOTWaw6i/KpyrJ1DRF3pybhae46I13pGQqGL4c7eJqlGo3l2t75h69H/NjwG+4lFDzoZUK+ca2nuglaHxbIeGoNO/Pm+cSMhl7kqvWZiL4/WKFpDAJVnA1QJ9pnq99/X9kbNMsxbNuOSKSO3pbHzVQetCEGAeYmj7KaCvGSXSbHwcoiFOkHFWfbPrmsHjDwltBziJRjADz1brQ6J/D"]}]}`

		httpmock.RegisterResponder(
			"GET",
			"https://unit-test-issuer.us.auth0.com/.well-known/jwks.json",
			httpmock.NewStringResponder(200, json),
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
		assert.Equal(t, "unable to find appropriate key", err.Error())
	})

	t.Run("Should return validated token when everything is Ok with the token", func(t *testing.T) {
		os.Setenv("AUD", "https://unit-test-audience/")
		os.Setenv("ISS", "https://unit-test-issuer.us.auth0.com/")

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		json := `{"keys":[{"alg":"RS256","kty":"RSA","use":"sig","n":"x-N4R5lgHyXWfjf-izlxrrr2LAn7bUq1cL069yB0G4sy9FCM1RBeet1tHeQ3szbCxYxZIZ1ODRu9BuK34TyEkyBNtAOITU5WjUVuMrWd9iK-noIVJwhykLooGwHVSUCMPjeRNd7sxf2WW3uwR1R3PglZeu25pBR0e9PxI8tUU8QWsMOdrCRw5tMyoqC5SQsa1J4HIzuTaYfuOClF4kpv933_c79VquvdrWEJ1MzDHG2Lfrb_wxaFuOMXzPSTnOsINWwG2-0rb0UXXm_emsa8NrDu2Wi-nlw0UYAwVUQEtwK_5KegZWI39pp3aaDR62jdiEiL85BulrEjefxGZVHDUw","e":"AQAB","kid":"M0Xo-5mQq2nlEDgkbZeEm","x5t":"4Qi2aVFszWUy5QBdl52a-BLZjZU","x5c":["MIIDETCCAfmgAwIBAgIJWh9HI7fDZC1/MA0GCSqGSIb3DQEBCwUAMCYxJDAiBgNVBAMTG2FsdXJhLWZsaXgtYXBpLnVzLmF1dGgwLmNvbTAeFw0yMTA4MTQwNDQ2NDlaFw0zNTA0MjMwNDQ2NDlaMCYxJDAiBgNVBAMTG2FsdXJhLWZsaXgtYXBpLnVzLmF1dGgwLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMfjeEeZYB8l1n43/os5ca669iwJ+21KtXC9OvcgdBuLMvRQjNUQXnrdbR3kN7M2wsWMWSGdTg0bvQbit+E8hJMgTbQDiE1OVo1FbjK1nfYivp6CFScIcpC6KBsB1UlAjD43kTXe7MX9llt7sEdUdz4JWXrtuaQUdHvT8SPLVFPEFrDDnawkcObTMqKguUkLGtSeByM7k2mH7jgpReJKb/d9/3O/Varr3a1hCdTMwxxti362/8MWhbjjF8z0k5zrCDVsBtvtK29FF15v3prGvDaw7tlovp5cNFGAMFVEBLcCv+SnoGViN/aad2mg0eto3YhIi/OQbpaxI3n8RmVRw1MCAwEAAaNCMEAwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQU/FadD6LzgRq42C86qGuHfwmzlB0wDgYDVR0PAQH/BAQDAgKEMA0GCSqGSIb3DQEBCwUAA4IBAQBNDGgCFs5Wy71637Zon7VDEP8LdnsaeAACedMEJKMxh80AifEQviqSufo9LWgck4vsSfTeAWREDPxJ7rFhh4siHemQpm+8fExPmZc1NSH0+2xaPGJfeBX+GrUAVlHmObzbgChfKXvOI07+41JmxCKqTYAbu5/AHCAwOyF65JS3XEiatmmisuECOoM71+QSMxNhJOFMUK9Rysjb5XidpFB3mC2OLFy7SEvHbZuGUyS+sE4k9xSYl5zxO+DO8e2dCGdDs3MKX8XNIEvnTdR65i6gm0+a1/aastr4GNNvbPxiI7ELBFcn6iWI/0zL54Rvv5rc0WWJ6t772hDG3+JCJqs9"]},{"alg":"RS256","kty":"RSA","use":"sig","n":"malZ2q_aHX7VD_ykryOYQIOHmyKT1Q94rdUKZGLxp0Rw0s_livESCmOgrKqLxVjEQmUUokqThMhAiDi7OPcrzy150iYk5J7wmj-D3eDvFFiABnBDlvt2lSLPmUY4R-NTRQ1wNfbLKmQycOrWTAGT9P4VXp45IARuRdFtjU9lsXmifWpCEcLlv61WPMzL0b9ld_GBAWvvE-sbINOpzm_xBrPwcIsImQNAsN9mFmZSSaiVQ7bQOpExergecF39yaTxXA0PfSorcsVW6XEvi3UQgS9HCdVjX2VXuCdu_HvnC-rRuqXrXPqSMq3QmPvqLwWK53DEhCxroHGKKoG2CKgbTw","e":"AQAB","kid":"GKhfLaIlbtpIESk_Aedrc","x5t":"gvRF9c9nmTlsv0-O0_Oik4lOBjc","x5c":["MIIDETCCAfmgAwIBAgIJb279+r/8IMU9MA0GCSqGSIb3DQEBCwUAMCYxJDAiBgNVBAMTG2FsdXJhLWZsaXgtYXBpLnVzLmF1dGgwLmNvbTAeFw0yMTA4MTQwNDQ2NDlaFw0zNTA0MjMwNDQ2NDlaMCYxJDAiBgNVBAMTG2FsdXJhLWZsaXgtYXBpLnVzLmF1dGgwLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAJmpWdqv2h1+1Q/8pK8jmECDh5sik9UPeK3VCmRi8adEcNLP5YrxEgpjoKyqi8VYxEJlFKJKk4TIQIg4uzj3K88tedImJOSe8Jo/g93g7xRYgAZwQ5b7dpUiz5lGOEfjU0UNcDX2yypkMnDq1kwBk/T+FV6eOSAEbkXRbY1PZbF5on1qQhHC5b+tVjzMy9G/ZXfxgQFr7xPrGyDTqc5v8Qaz8HCLCJkDQLDfZhZmUkmolUO20DqRMXq4HnBd/cmk8VwND30qK3LFVulxL4t1EIEvRwnVY19lV7gnbvx75wvq0bql61z6kjKt0Jj76i8FiudwxIQsa6BxiiqBtgioG08CAwEAAaNCMEAwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUdsVi3xAtWNTvD8hYUbjerqtCTbkwDgYDVR0PAQH/BAQDAgKEMA0GCSqGSIb3DQEBCwUAA4IBAQBumE6HlpDk8Gw8KSkkay75qfWzx3meilu3RqpcoKEXausq70Xr5HfVnXl493trW5aBwgZCn5OzPfWWTIi4XpmSMeAwZRM9zJ3WfMQzO/M0ObF7K5s3wYLcc0t+djha/dZggdiOTWaw6i/KpyrJ1DRF3pybhae46I13pGQqGL4c7eJqlGo3l2t75h69H/NjwG+4lFDzoZUK+ca2nuglaHxbIeGoNO/Pm+cSMhl7kqvWZiL4/WKFpDAJVnA1QJ9pnq99/X9kbNMsxbNuOSKSO3pbHzVQetCEGAeYmj7KaCvGSXSbHwcoiFOkHFWfbPrmsHjDwltBziJRjADz1brQ6J/D"]}]}`

		httpmock.RegisterResponder(
			"GET",
			"https://unit-test-issuer.us.auth0.com/.well-known/jwks.json",
			httpmock.NewStringResponder(200, json),
		)

		claims := jwt.MapClaims{
			"exp": 15000,
			"iss": "https://unit-test-issuer.us.auth0.com/",
			"aud": []string{"https://unit-test-audience/"},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
		token.Header = map[string]interface{}{"kid": "M0Xo-5mQq2nlEDgkbZeEm"}
		result, err := ValidateToken(token)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})
}
