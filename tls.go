package handler

import (
	"crypto/tls"
	"crypto/x509"
	"os"
)

func createTLSConfig(caCertFile, certFile, keyFile string, skipVerification bool) (*tls.Config, error) {
	// Load the CA certificate to validate the server's certificate.
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// Create a certificate pool and add the CA certificate.
	certPool := x509.NewCertPool()
	caCert, err := os.ReadFile(caCertFile)
	if err != nil {
		return nil, err
	}
	certPool.AppendCertsFromPEM(caCert)

	// Create a TLS configuration with the certificate and key.
	return &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            certPool,
		InsecureSkipVerify: skipVerification,
	}, nil
}
