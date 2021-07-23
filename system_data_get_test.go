package bexio_test

import (
	"encoding/json"
	"fmt"
	"testing"

	bexio "github.com/omniboost/go-bexio"
)

func TestSystemDataGet(t *testing.T) {
	req := client.NewSystemDataGetRequest()
	req.QueryParams().BusinessObject = "fld"
	req.QueryParams().PageSize = 256
	req.QueryParams().FilterControlSource1 = "DBF_NAME"
	req.QueryParams().FilterOperator1 = "Contains"
	// req.QueryParams().FilterValue1 = "VAT     "
	// req.QueryParams().FilterValue1 = "LEDGER  "
	// req.QueryParams().FilterValue1 = "COMP_HDR"
	// req.QueryParams().FilterValue1 = "TRANSACT"
	// req.QueryParams().FilterValue1 = "DAILY   "
	// req.QueryParams().FilterValue1 = "ACCT_REC"
	// req.QueryParams().FilterValue1 = "COST"
	req.QueryParams().FilterValue1 = "DJ_LINE"
	// req.QueryParams().FilterValue1 = "ADM_LIST"
	req.QueryParams().FilterValueType1 = "C"
	// req.QueryParams().SortFields = "FLD_DESC"
	// req.QueryParams().SortOrder = "Ascending"
	req.QueryParams().Fields = bexio.Fields{
		"dct_fld.FLD_DESC",
		"dct_fld.FIELD_NAME",
		"dct_fld.DBF_NAME",
		"dct_fld.FIELD_LEN",
		"dct_fld.FIELD_TYPE",
		"dct_fld.INP_MAND",
		"dct_fld.FLD_FMT",
		"dct_fld.FLD_MASK",
		"dct_fld.COM_BLK",
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	// https://www.bexio/api/v3/accountviewsystemdata?businessobject=fld
	// &pagesize=256
	// &fields=dct_fld.FLD_DESC,dct_fld.FIELD_NAME,dct_fld.DBF_NAME,dct_fld.FIELD_LEN,dct_fld.FIELD_TYPE,dct_fld.INP_MAND,dct_fld.FLD_FMT,dct_fld.FLD_MASK,dct_fld.COM_BLK
	// &FilterControlSource1=DBF_NAME
	// &FilterOperator1=Contains
	// &FilterValue1=VAT
	// &FilterValueType1=C
	// &SortFields=FLD_DESC
	// &SortOrder=Ascending

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
