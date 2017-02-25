package keys

import (
    "encoding/hex"
    "github.com/dedis/crypto/eddsa"
    "github.com/gopherjs/gopherjs/js"
    "crypto/cipher"
    "github.com/dedis/crypto/ed25519"
)

func GeneratePublicFromSecret(secret string) *js.Object {
    var stream cipher.Stream
    if secret != "" {
        buffer, _ := hex.DecodeString(secret)
        stream = ConstantStream(buffer)
    }

    keys := eddsa.NewEdDSA(stream)
    dataPublic, _ := keys.Public.MarshalBinary()
    public := hex.EncodeToString(dataPublic)

    dataSecret, _ := keys.Secret.MarshalBinary()
    private := hex.EncodeToString(dataSecret)

    return js.MakeWrapper(&keyPair{public, private})
}

func GenerateSecretAndPublic() *js.Object {
    return GeneratePublicFromSecret("")
}

func Sign(secret string, message string) string {
    secretBuffer, _ := hex.DecodeString(secret)
    e := eddsa.NewEdDSA(ConstantStream(secretBuffer))

    messageBuffer, _ := hex.DecodeString(message)
    signedBuffer, _ := e.Sign(messageBuffer)

    signed := hex.EncodeToString(signedBuffer)

    return signed
}

func Verify(pubkey string, msg, signature string) error {
    suite := ed25519.NewAES128SHA256Ed25519(false)

    buffer, _ := hex.DecodeString(pubkey)
    public := suite.Point()
    public.UnmarshalBinary(buffer)

    msgBuffer, _ := hex.DecodeString(msg)
    signatureBuffer, _ := hex.DecodeString(signature)
    return eddsa.Verify(public, msgBuffer, signatureBuffer)
}

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
