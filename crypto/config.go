package crypto

import (
	"crypto/sha256"
	"encoding/binary"
	"sort"
	"github.com/gopherjs/gopherjs/js"
	"github.com/dedis/crypto/eddsa"
	"github.com/dedis/crypto/ed25519"
	"github.com/dedis/crypto/random"
)

// Given a Config object, return the hash
func
HashConfig(config *js.Object) []byte {

	hash := sha256.New()

	// Hash threshold
	binary.Write(hash, binary.LittleEndian, int32(config.Get("threshold").Int()));

	// Extract and sort owners
	var owners []string
	for s := range config.Get("device").Interface().(map[string]interface{}) {
		owners = append(owners, s)
	}
	sort.Strings(owners)

	// Iterate over sorted owners
	for _, s := range owners {
		// Hash owner
		hash.Write([]byte(s))

		// Hash data
		if config.Get("data").Get(s) != js.Undefined {

			// TODO: why Get(s) ?!? ssh:tel1:server != tel1 (owner)
			data := (config.Get("data").Get(s).String())
			hash.Write([]byte(data))
		}

		// Hash point
		hash.Write([]byte(config.Get("device").Get(s).Get("point").Interface().([]byte)))
	}

	return hash.Sum(nil)
}

func SchnorrSignature(keyPair, msg []byte) []byte {

	key := eddsa.NewEdDSA(nil)
	key.UnmarshalBinary(keyPair)

	suite := ed25519.NewAES128SHA256Ed25519(false)

	k := suite.Scalar().Pick(random.Stream)
	r := suite.Point().Mul(nil, k)

	// [call to hash] ...compute challenge 'e'
	rBuf, _ := r.MarshalBinary()
	cipher := suite.Cipher(rBuf)
	cipher.Message(nil, nil, msg)
	e := suite.Scalar().Pick(cipher)

	// Compute response 's'
	xe := suite.Scalar().Mul(key.Secret, e)
	s := suite.Scalar().Sub(k, xe)

	return append(e.Bytes(), s.Bytes()...)
}

// TODO: Extract challenge (e) and response (s) from byte array
/*
func SchnorrVerify(keyPair, msg, signature []byte) bool {

	suite := ed25519.NewAES128SHA256Ed25519(false)

	gs := suite.Point().Mul(nil, s)
	ye := suite.Point().Mul(key.Public, e)
	rv := suite.Point().Add(gs, ye)

	// [call to hash]'
	rBuf2, _ := rv.MarshalBinary()
	cipher2 := suite.Cipher(rBuf2)
	cipher2.Message(nil, nil, msg)
	er := suite.Scalar().Pick(cipher2)

	return er.Equal(e)
}
*/