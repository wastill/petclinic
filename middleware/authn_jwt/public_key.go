package authn_jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
)

func WithPublicKeyPath(p string) Option {
	return func(o *options) {
		o.publicKey = loadPubKey(p)
	}
}

func loadPubKey(path string) *rsa.PublicKey {
	f, e := ioutil.ReadFile(path)
	if e != nil {
		log.Fatalf("load pubkey %s failed", path)
		// return nil, fmt.Errorf("load pubkey %v failed", path)
	}

	pubPem, _ := pem.Decode(f)
	if pubPem.Type != "PUBLIC KEY" {
		log.Fatalf("RSA pubkey is of the wrong type: %s", pubPem.Type)
		// return nil, fmt.Errorf("RSA pubkey is of the wrong type: %v", pubPem.Type)
	}

	var parsedKey interface{}
	var err error
	if parsedKey, err = x509.ParsePKIXPublicKey(pubPem.Bytes); err != nil {
		log.Fatalf("unable to parse RSA public key %s", err.Error())
		// return nil, fmt.Errorf("unable to parse RSA public key %w", err)
	}

	var pubKey *rsa.PublicKey
	var ok bool
	pubKey, ok = parsedKey.(*rsa.PublicKey)
	if !ok {
		log.Fatalf("unable to cast to RSA public key %s", err.Error())
		// return nil, fmt.Errorf("unable to cast to RSA public key %v", err)
	}

	return pubKey
}
