package rmcserver

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"github.com/howeyc/gopass"
)

// LoadPEMCertificateFromFile loads a x509 PEM certificate from a file
func (server *RMCServer) LoadPEMCertificateFromFile(pemLocation string) (err error) {
	var certBytes []byte
	if certBytes, err = ioutil.ReadFile(pemLocation); err != nil {
		return
	}

	// turn this into a pem block so we can check if it's encrypted
	var block *pem.Block
	if block, _ = pem.Decode(certBytes); err != nil {
		return
	}

	// if it's encrypted unlock it, take input from command line
	if x509.IsEncryptedPEMBlock(block) {

		// Get password from command line input
		fmt.Printf("Certificate password: ")
		var password []byte
		if password, err = gopass.GetPasswd(); err != nil {
			return
		}

		// Decrypt the PEM block into certBytes (which will then be DER encoded)
		if certBytes, err = x509.DecryptPEMBlock(block, password); err != nil {
			return
		}

	}

	// Finally get the certificate
	if server.certificate, err = x509.ParseCertificate(certBytes); err != nil {
		return
	}

	return
}
