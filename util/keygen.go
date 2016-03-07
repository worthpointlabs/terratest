package util

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/ssh"
)

// Generate an RSA Keypair and return the public and private keys
func GenerateRSAKeyPair(keySize int) (string, string, error) {

	rsaKeyPair, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return "", "", fmt.Errorf("Failed to generate key: %s", err)
	}

	// Extract the private key
	keyPemBlock := &pem.Block{
		Type:	"RSA PRIVATE KEY",
		Bytes:	x509.MarshalPKCS1PrivateKey(rsaKeyPair),
	}

	keyPem := string(pem.EncodeToMemory(keyPemBlock))

	// Extract the public key
	sshPubKey, err := ssh.NewPublicKey(rsaKeyPair.Public())
	if err != nil {
		return "", "", fmt.Errorf("Unable to generate new OpenSSH public key.")
	}

	sshPubKeyBytes := ssh.MarshalAuthorizedKey(sshPubKey)
	sshPubKeyStr := string(sshPubKeyBytes)

	// Return
	return sshPubKeyStr, keyPem, nil
}
