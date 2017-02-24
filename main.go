package main

import (
    "github.com/gopherjs/gopherjs/js"
    "github.com/dedis/crypto-js/keys"
)

func main() {
    js.Global.Set("cryptoJS", map[string]interface{}{
        "GenerateSecretAndPublic": keys.GenerateSecretAndPublic,
        "GeneratePublicFromSecret": keys.GeneratePublicFromSecret,
        "Sign": keys.Sign,
        "Verify": keys.Verify,
    })
}
