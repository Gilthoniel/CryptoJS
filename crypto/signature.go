package crypto

import (
    "encoding/hex"
    "github.com/dedis/crypto/ed25519"
    "github.com/dedis/crypto/eddsa"
)

// Sign the given message using the given private key
func Sign(secret string, message string) string {
    secretBuffer, _ := hex.DecodeString(secret)
    e := eddsa.NewEdDSA(ConstantStream(secretBuffer))

    messageBuffer, _ := hex.DecodeString(message)
    signedBuffer, _ := e.Sign(messageBuffer)

    signed := hex.EncodeToString(signedBuffer)

    return signed
}

// Verify a given signature with the given public key
// return nil if the signature is correct, else an error
func Verify(pubkey string, msg, signature string) error {
    suite := ed25519.NewAES128SHA256Ed25519(false)

    buffer, _ := hex.DecodeString(pubkey)
    public := suite.Point()
    public.UnmarshalBinary(buffer)

    msgBuffer, _ := hex.DecodeString(msg)
    signatureBuffer, _ := hex.DecodeString(signature)
    return eddsa.Verify(public, msgBuffer, signatureBuffer)
}
