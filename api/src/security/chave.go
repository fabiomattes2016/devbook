package security

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GerarChave() string {
	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)

	return stringBase64
}
