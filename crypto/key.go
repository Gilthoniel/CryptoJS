package crypto

import (
    "github.com/dedis/crypto/eddsa"
    "github.com/dedis/crypto/ed25519"
)

// Return a byte array representation of the key pair
func KeyPairEdDSA() []byte {
    e := eddsa.NewEdDSA(nil)

    result, _ := e.Secret.MarshalBinary()
    return result
}

func KeyPairPublic(secret []byte) []byte {
    stream := ConstantStream(secret)
    e := eddsa.NewEdDSA(stream)

    result, _ := e.Public.MarshalBinary()
    return result
}

// Aggregate the given public keys in one single key
func AggregateKeys(keys [][]byte) []byte {
    suite := ed25519.NewAES128SHA256Ed25519(false)
    aggKey := suite.Point().Null()

    for _, k := range keys {
        public := suite.Point()
        public.UnmarshalBinary(k)

        aggKey.Add(aggKey, public)
    }

    result, _ := aggKey.MarshalBinary()
    return result
}
