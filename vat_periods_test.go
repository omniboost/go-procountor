package procountor_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestVATPeriodsGet(t *testing.T) {
	req := client.NewVATPeriodsGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
