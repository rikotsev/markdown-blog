package server

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rikotsev/markdown-blog/be/gen"
	"github.com/rikotsev/markdown-blog/be/internal/config"
	"math/big"
	"net/http"
	"strings"
	"time"
)

type (
	Identity struct {
		Email string
	}
	AuthenticationProvider interface {
		Handle(r *http.Request) (Identity, error)
	}
	Key struct {
		Kty string   `json:"kty,omitempty"`
		Use string   `json:"use,omitempty"`
		N   string   `json:"n,omitempty"`
		E   string   `json:"e,omitempty"`
		Kid string   `json:"kid,omitempty"`
		X5t string   `json:"x5t,omitempty"`
		X5c []string `json:"x5c,omitempty"`
		Alg string   `json:"alg,omitempty"`
	}
	Jwks struct {
		Keys []Key `json:"keys,omitempty"`
	}
	oktaAuthenticationProvider struct {
		JwksUrl string
		Jwks    Jwks
	}
	ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error)
)

var _ AuthenticationProvider = (*oktaAuthenticationProvider)(nil)
var ErrUnauthorized = errors.New("the client lacks proper credentials")

func Okta(cfg *config.Config) AuthenticationProvider {
	return &oktaAuthenticationProvider{
		JwksUrl: cfg.Auth.JwksUrl,
		Jwks:    Jwks{},
	}
}

func (o *oktaAuthenticationProvider) fetchJwks() error {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(o.JwksUrl)
	if err != nil {
		return fmt.Errorf("failed to get jwks: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get jwks: %w", err)
	}

	err = json.NewDecoder(resp.Body).Decode(&o.Jwks)
	if err != nil {
		return fmt.Errorf("failed to read jwks: %w", err)
	}

	return nil
}

func (o *oktaAuthenticationProvider) findKey(kid string) *Key {
	for _, k := range o.Jwks.Keys {
		if k.Kid == kid {
			return &k
		}
	}

	return nil
}

func (o *oktaAuthenticationProvider) keyFunc(token *jwt.Token) (interface{}, error) {
	kid, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("kid not found in token header")
	}

	//TODO cache based on expiry header or pre-defined interval
	if len(o.Jwks.Keys) == 0 {
		err := o.fetchJwks()
		if err != nil {
			return nil, fmt.Errorf("could not fetch jwks: %w", err)
		}
	}

	key := o.findKey(kid)
	if key == nil {
		return nil, errors.New("could not find matching key")
	}

	nBytes, err := base64.RawURLEncoding.DecodeString(key.N)
	if err != nil {
		return nil, errors.New("could not decode N")
	}

	eBytes, err := base64.RawURLEncoding.DecodeString(key.E)
	if err != nil {
		return nil, errors.New("could not decode E")
	}

	pubKey := rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: int(new(big.Int).SetBytes(eBytes).Int64()),
	}

	return &pubKey, nil
}

func (o *oktaAuthenticationProvider) Handle(r *http.Request) (Identity, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return Identity{}, ErrUnauthorized
	}

	token := strings.Replace(authHeader, "Bearer ", "", 1)

	parseToken, err := jwt.Parse(token, o.keyFunc, jwt.WithAudience("urn:markdown-blog:api"))
	if err != nil {
		return Identity{}, err
	}

	if !parseToken.Valid {
		return Identity{}, err
	}

	subject, err := parseToken.Claims.GetSubject()
	if err != nil {
		return Identity{}, err
	}

	return Identity{
		Email: subject,
	}, nil
}

func AuthAsMiddleware(provider AuthenticationProvider, errorHandler ErrorHandler) gen.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "GET" {
				next.ServeHTTP(w, r)
				return
			}

			_, err := provider.Handle(r)
			if err != nil {
				errorHandler(w, r, err)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
