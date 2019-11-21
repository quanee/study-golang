package main

import (
	"fmt"
	"math/big"
)

func main() {
	mod := big.NewInt(49999)
	base := big.NewInt(49547)

	a := big.NewInt(49459)
	A := base.Exp(base, a, mod)

	b := big.NewInt(49429)
	B := base.Exp(base, b, mod)

	AS := B.Exp(B, a, mod)
	BS := A.Exp(A, b, mod)
	fmt.Println(AS.String(), BS.String())
}
