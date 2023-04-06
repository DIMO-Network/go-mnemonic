package main

import (
	"encoding/hex"
	"fmt"

	"github.com/DIMO-Network/go-mnemonic"
)

func main() {
	s := "0xd744468B9192301650f8Cb5e390BdD824DFA6Dd9"

	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	words, _ := mnemonic.EntropyToMnemonic(data)
	fmt.Println(words)
}
