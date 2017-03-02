package crypto

import (
    "github.com/dedis/crypto/ed25519"
    "github.com/dedis/crypto/eddsa"
)

// Sign the given message with the private key provided
// as a Marshal Binary representation
func Sign(secret, message []byte) []byte {
    stream := ConstantStream(secret)
    e := eddsa.NewEdDSA(stream)

    signed, _ := e.Sign(message)
    return signed
}

// Verify a given signature with the given public key
// return nil if the signature is correct, else an error
func Verify(pubkey, msg, signature []byte) error {
    suite := ed25519.NewAES128SHA256Ed25519(false)

    public := suite.Point()
    public.UnmarshalBinary(pubkey)

    return eddsa.Verify(public, msg, signature)
}
