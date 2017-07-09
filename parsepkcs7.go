package main

import (
	"encoding/pem"
	"fmt"
	"log"

	"go.mozilla.org/pkcs7"
)

func main() {
	derBlock, _ := pem.Decode([]byte(`-----BEGIN PKCS7-----
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
-----END PKCS7-----`))
	sig, err := pkcs7.Parse(derBlock.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	sig.Content = []byte("cariboumaurice\n")
	err = sig.Verify()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("verified")
}