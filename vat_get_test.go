package bexio_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestVATGet(t *testing.T) {
	req := client.NewVATGetRequest()
	req.QueryParams().PageSize = 20
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
