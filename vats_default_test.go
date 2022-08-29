package procountor_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestVatsDefaultGet(t *testing.T) {
	req := client.NewVatsDefaultGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
