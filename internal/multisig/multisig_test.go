package multisig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildP2SHAddressInvalidN(t *testing.T) {
	_, err := BuildP2SHAddress(100, 10, nil)
	assert.Error(t, err)
}

func TestBuildP2SHAddressInvalidM(t *testing.T) {
	_, err := BuildP2SHAddress(10, -1, nil)
	assert.Error(t, err)
}

func TestBuildP2SHAddressInvalidPubKeys(t *testing.T) {
	_, err := BuildP2SHAddress(3, 5, [][]byte{{}, {}, {}})
	assert.Error(t, err)
}

func TestBuildP2SHAddress(t *testing.T) {
	pubkeys := []string{"020f8796e0f870a9a3b269be3b1e78e380c9b569885f0de98a9ff061c4a66e79d2", "02dfa8990f3f015ff20e9b31b85ea36d47470220615fb2ac1597e20fc830727b25", "03fbfbdc5df9c60e4b747805552686199e85299a5e87804dbb66a14597ddabcf29"}
	b := [][]byte{}
	for _, key := range pubkeys {
		b = append(b, []byte(key))
	}

	address, err := BuildP2SHAddress(2, 3, b)

	assert.ErrorIs(t, err, nil)
	assert.Equal(t, address, "38GxnyunFjXkHQMUEU5m4gqhD1fmKcRuEd")
}
