package trustlists

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"io/ioutil"

	"github.com/philhug/go-trustlists/internal/gen"
)

type Options struct {
}

func EUCertPool(opts Options) (*x509.CertPool, error) {
	roots := x509.NewCertPool()
	for _, s := range gen.EUCerts {
		r := bytes.NewBuffer([]byte(s))
		cs := base64.NewDecoder(base64.StdEncoding, r)
		cb, err := ioutil.ReadAll(cs)
		if err != nil {
			return nil, err
		}
		c, err := x509.ParseCertificate(cb)
		if err != nil {
			return nil, err
		}
		roots.AddCert(c)
	}
	return roots, nil
}
