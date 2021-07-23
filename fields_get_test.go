package bexio_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFieldsGet(t *testing.T) {
	req := client.NewFieldsGetRequest()
	req.QueryParams().BusinessObject = "FLD"
	req.QueryParams().PageSize = 1000
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
