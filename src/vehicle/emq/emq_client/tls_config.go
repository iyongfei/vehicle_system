package emq_client

import (
	"crypto/tls"
	"crypto/x509"
)

var (
	caBlock = []byte(
		`
-----BEGIN CERTIFICATE-----
MIIDpzCCAo+gAwIBAgIJAOOmrT4yd1QbMA0GCSqGSIb3DQEBDQUAMGoxFzAVBgNV
BAMMDkFuIE1RVFQgYnJva2VyMRYwFAYDVQQKDA1Pd25UcmFja3Mub3JnMRQwEgYD
VQQLDAtnZW5lcmF0ZS1DQTEhMB8GCSqGSIb3DQEJARYSbm9ib2R5QGV4YW1wbGUu
bmV0MB4XDTE5MDQxMDAyMTUzMVoXDTI5MDQwNzAyMTUzMVowajEXMBUGA1UEAwwO
QW4gTVFUVCBicm9rZXIxFjAUBgNVBAoMDU93blRyYWNrcy5vcmcxFDASBgNVBAsM
C2dlbmVyYXRlLUNBMSEwHwYJKoZIhvcNAQkBFhJub2JvZHlAZXhhbXBsZS5uZXQw
ggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDVrdBDzMvZ55GtOGhgj91Q
bOLcZkeFagUncWzw0n9EDqGAb5jyulaQysC8HlS4+cDWta0W72roeKImvK2hEU2T
5DuUW+DM2WXOv34KNtDQRLSggtcj3bQjsNBCgf3bM/Afy5kyvu/Ncf+LGjHrn4S0
crIAsSWo2geBsSrjXZHCmO1TfsYPsPiI7P1B8FO+6gEhEPXohHGIjNl5oIAN1fHK
Yh/E5GRCGxrLsxPE+4mFp9+xOAtseF5SWSxE1B/hHZt+ARgk7dh6rv0mBL5D7wKO
lGskLrNA7pxXlzAGA77hqFnvCnrFQei9rCHkcx1LvYjvPL4v1BhDWxcVuLP/8Ofl
AgMBAAGjUDBOMB0GA1UdDgQWBBQOmUIxoCW+b2c9PkKTGW9BmGvDKjAfBgNVHSME
GDAWgBQOmUIxoCW+b2c9PkKTGW9BmGvDKjAMBgNVHRMEBTADAQH/MA0GCSqGSIb3
DQEBDQUAA4IBAQDCRR4sKdFiC124qA0q8VqpphKEEmbWE6zWVthdtTP5h2xJzwbu
TCRXtAsDanUo4+4gXz5hLpZahhTKf0TOtZnF1+SVZoBu2zahBzb9R2XGoN7+6mo7
wkwoUdzUrEBW2phplY0nDcdagucWnkHXMs1F8fHMadG3ef2t/OeI0h5q0Kzx3f4a
YP6qORYpdzci5eAYA3mHH9PaD49TkA6WEDfyTEEskYtjvxmHcpD7XqJsYxycfCla
xmzfn9PNpf5eq7+2FUahOXwDvWZszUs8zjKyye/++6h5AtmL2LKKPTt8G8DF6Arn
2sRe0m0lT6+NOOQtJEOs7j4oamHubfG6tBIX
-----END CERTIFICATE-----
		`)
)


func NewTLSConfig() *tls.Config  {
	certpool := x509.NewCertPool()
	certpool.AppendCertsFromPEM(caBlock)
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: false,
		// Certificates = list of certs client sends to server.
		// Certificates: []tls.Certificate{cert},
	}
}