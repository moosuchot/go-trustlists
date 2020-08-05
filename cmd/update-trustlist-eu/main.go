package main

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"encoding/xml"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/philhug/go-trustlists/internal"
)

type Certs struct {
	Certs []string
}

func main() {
	var certs Certs
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	si := internal.TrustServiceStatusList{}
	err = xml.NewDecoder(f).Decode(&si)
	if err != nil {
		panic(err)
	}
	for _, c := range si.SchemeInformation.PointersToOtherTSL.OtherTSLPointer.ServiceDigitalIdentities {
		s := c.ServiceDigitalIdentity.DigitalId.X509Certificate
		r := bytes.NewBuffer([]byte(s))
		cs := base64.NewDecoder(base64.StdEncoding, r)
		cb, err := ioutil.ReadAll(cs)
		if err != nil {
			panic(err)
		}
		_, err = x509.ParseCertificate(cb)
		if err != nil {
			panic(err)
		}
		certs.Certs = append(certs.Certs, s)
	}

	gen, err := os.Create(os.Args[3])
	defer gen.Close()
	if err != nil {
		panic(err)
	}
	certTemplate, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		panic(err)
	}
	t := template.Must(template.New("eucerts").Parse(string(certTemplate)))
	t.Execute(gen, certs)
}
