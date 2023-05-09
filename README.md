# EasyEncDec
Encrypt/Decrypt data easily with AES encryption

## Install:
`go get -u github.com/hektorinho/easyencdec`

## Generate Key File:
Generates a 15 character file that is only available to the user that created the key file. -
```golang
package main

import (
	"log"

	"github.com/hektorinho/easyencdec"
)

const output = "private.key"

func main() {
	if err := easyencdec.GenerateKeyFile(output); err != nil {
		log.Panicf("failed to write file to %s\n", output)
	}
}
```

## Encrypt File:
This will use my method of encrypting a file, it will only be available to the user that created the file and it will be encrypted by a randomly generated 15 character string plus the MAC address of your computer to ensure it only works if you have the MAC address of the of the computer that generated the file.
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
	if err := easyencdec.EncryptFile(samplefile, decryptedfile, keyfile); err != nil {
		log.Panicf("failed to encrypt file %s\n", samplefile)
	}
}
```

## Read Decrypted File:
The decryption uses the 15 character key file together with the mac address of your computer to build the key to decrypt the file to ensure a decent level of safety.
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
		log.Panicf("failed to decrypt datafile\n")
	}

	fmt.Println(string(byte))
}
```
