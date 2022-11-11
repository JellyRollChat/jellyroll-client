package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"log"
)

func genkeys() {
	pubpub, privpriv, errrr := ed25519.GenerateKey(rand.Reader)
	log.Printf("pub: %x", pubpub)
	log.Printf("priv: %x", privpriv)
	if errrr != nil {
		log.Println(errrr)
	}

	log.Printf("Signing pubkey: %x", pubpub)
	pubsig := ed25519.Sign(privpriv, pubpub)
	log.Printf("pub sig: %x", pubsig)

	sigverify := ed25519.Verify(pubpub, pubpub, pubsig)

	log.Printf("verified? %v", sigverify)
}
