package main

import (
    "github.com/gopherjs/gopherjs/js"
    "github.com/dedis/crypto-js/crypto"
)

/**
 * Encapsulate the library in the cryptoJS object that you can
 * find the global JS object
 */
func main() {
    js.Global.Set("cryptoJS", map[string]interface{}{
        "keyPair": crypto.KeyPairEdDSA,
        "keyPairPublic": crypto.KeyPairPublic,
        "sign": crypto.Sign,
        "verify": crypto.Verify,
        "aggregateKeys": crypto.AggregateKeys,
        "sha256": crypto.Sha256,
        "sha512": crypto.Sha512,
    })
}
