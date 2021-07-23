package bexio_test

import (
	"encoding/json"
	"log"
	"testing"

	bexio "github.com/omniboost/go-bexio"
)

func TestBexioDataPost(t *testing.T) {
	req := client.NewBexioDataPostRequest()
	req.SetRequestBody(bexio.BexioDataPostRequestBody{
		BookDate:       "2021-07-02T10:39:05.276Z",
		BusinessObject: "DJ2",
		Table: bexio.Table{
			Definition: bexio.TableDefinition{
				Name: "DJ_PAGE",
				Fields: bexio.TableDefinitionFields{
					{
						Name:      "RowId",
						FieldType: "C",
					},
					{
						Name:      "PAGE_NR",
						FieldType: "C",
					},
					{
						Name:      "DJ_CODE",
						FieldType: "C",
					},
					{
						Name:      "HDR_DESC",
						FieldType: "C",
					},
				},
			},
			DetailDefinitions: bexio.TableDetailDefinitions{
				{
					Name: "DJ_LINE",
					Fields: bexio.TableDefinitionFields{
						{
							Name:      "RowId",
							FieldType: "C",
						},
						{
							Name:      "HeaderId",
							FieldType: "C",
						},
						{
							Name:      "PAGE_NR",
							FieldType: "C",
						},
						{
							Name:      "DJ_CODE",
							FieldType: "C",
						},
						{
							Name:      "SRC_ID",
							FieldType: "C",
						},
						{
							Name:      "TRN_DATE",
							FieldType: "T",
						},
						{
							Name:      "ACCT_NR",
							FieldType: "C",
						},
						{
							Name:      "REC_ID",
							FieldType: "C",
						},
					},
				},
			},
		},
		TableData: bexio.TableData{
			Data: bexio.TableDataData{
				Rows: bexio.Rows{
					{
						Values: []interface{}{
							"1",
							"TEST",
							"860",
							"Test",
						},
					},
				},
			},
			DetailData: bexio.DetailData{
				{
					Rows: bexio.Rows{
						{
							Values: []interface{}{
								"1",
								"1",
								"TEST",
								"TEST",
								"TEST",
								"2021-11-30T13:37:37.450Z",
								"3811",
								"TEST",
							},
						},
					},
				},
			},
		},
	})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
