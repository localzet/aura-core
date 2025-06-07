// Package crypto provides common crypto libraries for Aura.
package crypto // import "github.com/localzet/aura/common/crypto"

import (
	"crypto/rand"
	"math/big"
)

func RandBetween(from int64, to int64) int64 {
	if from == to {
		return from
	}
	bigInt, _ := rand.Int(rand.Reader, big.NewInt(to-from))
	return from + bigInt.Int64()
}
