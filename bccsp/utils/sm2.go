package utils

import "github.com/tjfoc/gmsm/sm2"

// DERToSM2Certificate converts der to sm2
func DERToSM2Certificate(asn1Data []byte) (*sm2.Certificate, error) {
	return sm2.ParseCertificate(asn1Data)
}
