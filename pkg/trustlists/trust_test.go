package trustlists

import (
	"fmt"
	"testing"
)

func TestEUCertPool(t *testing.T) {
	p, err := EUCertPool(Options{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(p)
}
