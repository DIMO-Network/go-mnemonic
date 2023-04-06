package main

import (
	"crypto/rand"
	"fmt"

	"github.com/DIMO-Network/go-mnemonic"
)

func main() {

	token := make([]byte, 20)
	rand.Read(token)

	words, _ := mnemonic.EntropyToMnemonic(token)
	fmt.Println(words)

	threeWords, _ := mnemonic.EntropyToMnemonicThreeWords(token)
	fmt.Println(threeWords)
}
