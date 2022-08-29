package procountor_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCompanyGet(t *testing.T) {
	req := client.NewCompanyGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
