package crypto

import (
    "crypto/sha256"
    "crypto/sha512"
    "encoding/hex"
)

// Return the hex encoded sha256 hash of the provided string
func Sha256(text string) string {
    hash := sha256.New()

    bytes := []byte(text)
    hash.Write(bytes)

    return hex.EncodeToString(hash.Sum(nil))
}

// Return the hex encoded sha512 hash of the provided string
func Sha512(text string) string {
    hash := sha512.New()

    bytes := []byte(text)
    hash.Write(bytes)

    return hex.EncodeToString(hash.Sum(nil))
}
