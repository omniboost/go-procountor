package procountor_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestInvoicesGet(t *testing.T) {
	req := client.NewInvoicesGetRequest()
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

