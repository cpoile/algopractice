package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

func PrivateKey(p *big.Int) *big.Int {
	n, _ := rand.Int(rand.Reader, p)
	if n.Cmp(big.NewInt(2)) < 0 {
		return PrivateKey(p)
	} else {
		return n
	}
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	A := new(big.Int)
	return A.Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	public := new(big.Int)
	return public.Exp(public2, private1, p)
}
