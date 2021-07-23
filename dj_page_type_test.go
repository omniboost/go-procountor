package bexio_test

import (
	"encoding/json"
	"log"
	"testing"

	bexio "github.com/omniboost/go-bexio"
)

func TestDjPageTypeTest(t *testing.T) {
	o := bexio.DjPage{
		DjCode:  "860",
		HdrDesc: "TEST",
	}
	lines := []bexio.DjLine{
		{
			AcctNr: "9999",
			Amount: 12.1,
			RecID:  "REC_ID",
		},
		{
			AcctNr: "9999",
			Amount: -12.1,
			RecID:  "REC_ID",
		},
	}
	req, err := o.ToBexioDataPostRequest(client, lines)
	if err != nil {
		t.Error(err)
		return
	}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
