package token

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	key, err := LoadPem("testdata/private-key.pem")
	assert.NoError(t, err)
	jwt, err := Sign(key, map[string]interface{}{"foo": "bar"})
	assert.NoError(t, err)
	fmt.Println(jwt)
}

func TestVerify(t *testing.T) {
	key, err := LoadPem("testdata/private-key.pem")
	assert.NoError(t, err)
	jwt, err := Sign(key, map[string]interface{}{"foo": "bar"})
	assert.NoError(t, err)
	allclaims, err := Verify(&key.PublicKey, jwt)
	assert.NoError(t, err)
	json, err := json.Marshal(allclaims)
	assert.NoError(t, err)
	fmt.Println(string(json))
}
