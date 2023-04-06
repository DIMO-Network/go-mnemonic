# go-mnemonic

Go implementation of Bitcoin BIP39: Mnemonic code for generating deterministic keys

# Examples

```go
import (
	"fmt"

	"github.com/DIMO-Network/go-mnemonic"
)

func main() {

    words, _ := mnemonic.EntropyToMnemonic(token)
	fmt.Println(words)

	threeWords, _ := mnemonic.EntropyToMnemonicThreeWords(token)
	fmt.Println(threeWords)
}
```
