package signer

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/drand/kyber"
	"github.com/drand/kyber/pairing/bn256"
)

const (
	KeyVersion = 'T'
)

func PrivateKeyFromHex(s string) (kyber.Scalar, error) {
	seed, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	suite := bn256.NewSuiteG2()
	scalar := suite.Scalar().SetBytes(seed)
	return scalar, nil
}

func PubKeyFromBase58(s string) (kyber.Point, error) {
	data, ver, err := base58.CheckDecode(s)
	if err != nil {
		return nil, err
	}
	if ver != KeyVersion {
		return nil, fmt.Errorf("invalid version %d", ver)
	}
	suite := bn256.NewSuiteG2()
	point := suite.G2().Point()
	err = point.UnmarshalBinary(data)
	return point, err
}