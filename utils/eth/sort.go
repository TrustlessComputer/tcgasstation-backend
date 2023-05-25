package eth

import (
	"crypto/ecdsa"
	"math/big"
)

type ValidatorKeyInfo struct {
	Address string
	PrivKey *ecdsa.PrivateKey
	Big     *big.Int
}

// Implement the sort.Interface for the ValidatorKeyInfo type
type ValidatorKeyInfoSlice []ValidatorKeyInfo

func (s ValidatorKeyInfoSlice) Len() int {
	return len(s)
}

func (s ValidatorKeyInfoSlice) Less(i, j int) bool {
	return s[i].Big.Cmp(s[j].Big) == -1
}

func (s ValidatorKeyInfoSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
