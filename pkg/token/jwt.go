package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func Sign(privateKey *rsa.PrivateKey, claims map[string]interface{}) (string, error) {
	key := jose.SigningKey{Algorithm: jose.RS256, Key: privateKey}

	signerOpts := jose.SignerOptions{}
	signerOpts.WithType("JWT")
	rsaSigner, err := jose.NewSigner(key, &signerOpts)
	if err != nil {
		return "", err
	}

	now := time.Now()
	exp := jwt.NewNumericDate(now.Add(time.Second * 10 * 60))

	stdClaims := jwt.Claims{
		Expiry:   exp,
		IssuedAt: jwt.NewNumericDate(now),
	}

	builder := jwt.Signed(rsaSigner).Claims(stdClaims).Claims(claims)
	rawJWT, err := builder.CompactSerialize()
	if err != nil {
		return "", err
	}
	return rawJWT, nil
}

func Verify(publicKey *rsa.PublicKey, rawJWT string) ([]byte, error) {
	parsedJWT, err := jwt.ParseSigned(rawJWT)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT:%+v", err)
	}
	claims := make(map[string]interface{})
	err = parsedJWT.Claims(publicKey, &claims)
	if err != nil {
		return nil, fmt.Errorf("failed to verify JWT:%w", err)
	}
	bytes, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func LoadPem(path string) (*rsa.PrivateKey, error) {
	pemBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return keyFromPem(pemBytes)
}

func keyFromPem(pemBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
