package main

import (
	"crypto/ed25519"
	"crypto/rand"
)

func initKeys() *Ed25519Keys {
	if !fileExists(privKeyFilePath) {
		generateKeys()
	}
	keys := Ed25519Keys{}
	keyspublicKey := readFileBytes(pubKeyFilePath)
	keysprivateKey := readFileBytes(privKeyFilePath)
	keyssignedKey := readFileBytes(signedKeyFilePath)
	keys.publicKey = keyspublicKey
	keys.privateKey = keysprivateKey
	keys.signedKey = keyssignedKey
	return &keys
}

func keyTest() bool {

	ttk := Ed25519Keys{}
	ttk.publicKey, ttk.privateKey, _ = ed25519.GenerateKey(rand.Reader)
	ttk.signedKey = signMyKey(ttk.publicKey, ttk.privateKey)

	ttkVerify := verifySigBytes(ttk.publicKey, ttk.publicKey, ttk.signedKey)

	if !ttkVerify {
		panic("Key generation, signing, and verification is broken. Exiting.")
	}

	return true
}

func signMyKey(myPubKey ed25519.PublicKey, myprivkey ed25519.PrivateKey) []byte {
	return ed25519.Sign(myprivkey, myPubKey)
}

func verifySigBytes(pubkey, originalmsg, sig []byte) bool {
	return ed25519.Verify(pubkey, originalmsg, sig)
}

func generateKeys() *Ed25519Keys {
	keys := Ed25519Keys{}
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	handle("error: ", err)
	keys.privateKey = privKey
	keys.publicKey = pubKey
	signedKey := ed25519.Sign(privKey, pubKey)
	keys.signedKey = signedKey
	createFile(pubKeyFilePath)
	createFile(privKeyFilePath)
	createFile(signedKeyFilePath)
	writeFileBytes(pubKeyFilePath, keys.publicKey)
	writeFileBytes(privKeyFilePath, keys.privateKey)
	writeFileBytes(signedKeyFilePath, keys.signedKey)
	return &keys
}
