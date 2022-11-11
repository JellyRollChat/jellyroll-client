package main

var (
	GlobalBuddyList   = Buddylist{}
	GlobalRequestList = Buddylist{}
)

type Buddylist struct {
	Buddys []Buddy
}

type Buddy struct {
	Username string
	Server   string
	Pubkey   string
	Friend   bool
}

var (
	configPath        = "config"
	buddyListPath     = configPath + "/buddy.list"
	requestListPath   = configPath + "/request.list"
	keysPath          = configPath + "/keys"
	pubKeyFilePath    = keysPath + "/pub.key"
	privKeyFilePath   = keysPath + "/priv.key"
	signedKeyFilePath = keysPath + "/signed.key"
	// selfCertFilePath  = keysPath + "/self.cert"
)

type Ed25519Keys struct {
	publicKey  []byte
	privateKey []byte
	signedKey  []byte
	// selfCert   []byte
}
