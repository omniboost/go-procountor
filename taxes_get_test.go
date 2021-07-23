package bexio_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTaxesGet(t *testing.T) {
	req := client.NewTaxesGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
