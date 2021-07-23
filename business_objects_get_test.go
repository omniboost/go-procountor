package bexio_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestBusinessObjectsGet(t *testing.T) {
	req := client.NewBusinessObjectsGetRequest()
	// req.QueryParams().BusinessObject = "FLD"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
