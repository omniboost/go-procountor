package procountor_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDimensionsGet(t *testing.T) {
	req := client.NewDimensionsGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
