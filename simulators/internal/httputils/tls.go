package httputils

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/http"
	"os"
)

var errAppendCACert = errors.New("failed to append CA cert to pool")

const (
	mainPath       = "emu/"
	caFile         = "tls/root/ca.crt"
	serverCertFile = "tls/server/tls.crt"
	serverKeyFile  = "tls/server/tls.key"
	clientCertFile = "tls/client/tls.crt"
	clientKeyFile  = "tls/client/tls.key"
)

func loadCACertPool() (*x509.CertPool, error) {
	caPEM, err := os.ReadFile(mainPath + caFile)
	if err != nil {
		return nil, fmt.Errorf("reading CA cert: %w", err)
	}

	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(caPEM) {
		return nil, errAppendCACert
	}

	return pool, nil
}

// NewHTTPClient creates an HTTP client configured with TLS using the CA and client certificates.
func NewHTTPClient() (*http.Client, error) {
	caPool, err := loadCACertPool()
	if err != nil {
		return nil, fmt.Errorf("loading CA pool: %w", err)
	}

	cliCrt, err := tls.LoadX509KeyPair(mainPath+clientCertFile, mainPath+clientKeyFile)
	if err != nil {
		return nil, fmt.Errorf("loading client key pair: %w", err)
	}

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cliCrt},
				RootCAs:      caPool,
			},
		},
	}, nil
}

// NewHTTPServer creates an HTTP server configured with TLS using the CA and server certificates.
func NewHTTPServer(addr string, handler http.Handler) (*http.Server, error) {
	caPool, err := loadCACertPool()
	if err != nil {
		return nil, fmt.Errorf("loading CA pool: %w", err)
	}

	srvCrt, err := tls.LoadX509KeyPair(mainPath+serverCertFile, mainPath+serverKeyFile)
	if err != nil {
		return nil, fmt.Errorf("loading server key pair: %w", err)
	}

	return &http.Server{
		Addr:    addr,
		Handler: handler,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{srvCrt},
			ClientCAs:    caPool,
			ClientAuth:   tls.RequireAndVerifyClientCert,
		},
	}, nil
}
