package procountor_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestChartOfAccountsGet(t *testing.T) {
	req := client.NewChartOfAccountsGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
