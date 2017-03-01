package crypto

import (
    "encoding/hex"
    "github.com/dedis/crypto/eddsa"
    "github.com/dedis/crypto/ed25519"
)

// Generate a random private key for an EdDSA signature
func GeneratePrivateKey() string {
    keys := eddsa.NewEdDSA(nil)

    dataSecret, _ := keys.Secret.MarshalBinary()
    return hex.EncodeToString(dataSecret)
}

// Generate the public key given the private key for an EdDSA signature
func GeneratePublicKey(privateKey string) string {
    buffer, _ := hex.DecodeString(privateKey)
    stream := ConstantStream(buffer)

    keys := eddsa.NewEdDSA(stream)
    dataPublic, _ := keys.Public.MarshalBinary()
    return hex.EncodeToString(dataPublic)
}

// Aggregate the given public keys in one single key
func AggregateKeys(keys []string) string {
    suite := ed25519.NewAES128SHA256Ed25519(false)
    aggKey := suite.Point().Null()

    for _, k := range keys {
        public := suite.Point()
        buffer, _ := hex.DecodeString(k)
        public.UnmarshalBinary(buffer)

        aggKey.Add(aggKey, public)
    }

    aggKeyString, _ := aggKey.MarshalBinary()
    return hex.EncodeToString(aggKeyString)
}
