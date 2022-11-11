package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"log"
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

	log.Println(brightcyan + "Testing key creation, verification and signing..." + nc)

	ttk := Ed25519Keys{}
	ttk.publicKey, ttk.privateKey, _ = ed25519.GenerateKey(rand.Reader)
	ttk.signedKey = signMyKey(ttk.publicKey, ttk.privateKey)
	log.Printf(brightgreen+"[✔]"+cyan+" Test pubk:"+nc+"\t%x"+cyan+"..."+nc, ttk.publicKey[:4])
	log.Printf(brightgreen+"[✔]"+cyan+" Test priv:"+nc+"\t%x"+cyan+"..."+nc, ttk.privateKey[:4])
	log.Printf(brightgreen+"[✔]"+cyan+" Test sign:"+nc+"\t%x"+cyan+"..."+nc, ttk.signedKey[:4])

	ttkVerify := verifySigBytes(ttk.publicKey, ttk.publicKey, ttk.signedKey)

	if !ttkVerify {
		log.Println(red + "Key signature " + brightred + "failed!\n" + red + "Hey but really though, something is wrong, we can't verify keys." + nc)
		panic("Key generation, signing, and verification is broken. Exiting.")
	}

	log.Println(cyan + "Key signature " + brightgreen + "verified" + nc)

	return true
}

func signMyKey(myPubKey ed25519.PublicKey, myprivkey ed25519.PrivateKey) []byte {
	return ed25519.Sign(myprivkey, myPubKey)
}

func verifySigBytes(pubkey, originalmsg, sig []byte) bool {
	return ed25519.Verify(pubkey, originalmsg, sig)
}

func keysToString() {

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
