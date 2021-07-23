package bexio_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDjPageGet(t *testing.T) {
	req := client.NewDjPageGetRequest()
	req.QueryParams().PageSize = 20
	req.QueryParams().FilterControlSource1 = "DJ_CODE"
	req.QueryParams().FilterOperator1 = "Equal"
	req.QueryParams().FilterValueType1 = "C"
	req.QueryParams().FilterValue1 = "860"
	req.QueryParams().FilterValue1 = "600"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
