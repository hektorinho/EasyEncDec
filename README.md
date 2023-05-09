# EasyEncDec
Encrypt/Decrypt data easily with AES encryption

## Install:
`go get -u github.com/hektorinho/easyencdec`

## Generate Key File:
```golang
package main

import (
	"log"

	"github.com/hektorinho/easyencdec"
)

const output = "private.key"

func main() {
	if ok := easyencdec.GenerateKeyFile(output); !ok {
		log.Panicf("failed to write file to %s", output)
	}
}
```

## Encrypt File:
```golang
package main

import (
	"log"

	"github.com/hektorinho/easyencdec"
)

const (
	samplefile    = "test/postgresql.json"
	decryptedfile = "test/decrypt.crd"
	keyfile       = "private.key"
)

func main() {
	if ok := easyencdec.EncryptFile(samplefile, decryptedfile, keyfile); !ok {
		log.Panicf("failed to encrypt file %s", samplefile)
	}
}
```

## Read Decrypted File:
```golang
package main

import (
	"fmt"
	"log"

	"github.com/hektorinho/easyencdec"
)

const (
	samplefile    = "test/postgresql.json"
	decryptedfile = "test/decrypt.crd"
	keyfile       = "private.key"
)

func main() {
	byte, err := easyencdec.DecryptFile(decryptedfile, keyfile)
	if err != nil {
		log.Panicf("failed to decrypt datafile")
	}

	fmt.Println(string(byte))
}
```
