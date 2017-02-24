package keys

import (
    "crypto/cipher"
)

type keyPair struct {
    public string
    private string
}

func (k *keyPair) Public() string {
    return k.public
}

func (k *keyPair) Private() string {
    return k.private
}

type constantStream struct {
    seed []byte
}

// ConstantStream is a cipher.Stream which always returns
// the same value.
func ConstantStream(buff []byte) cipher.Stream {
    return &constantStream{buff}
}

// XORKexStream implements the cipher.Stream interface
func (cs *constantStream) XORKeyStream(dst, src []byte) {
    copy(dst, cs.seed)
}
